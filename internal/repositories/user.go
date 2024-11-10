package repositories

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"reliab-test/internal/domain"
)

type UserRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func BuildUserRepository(db *sqlx.DB, log *slog.Logger) *UserRepository {
	return &UserRepository{db: db, log: log}
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	builder := sq.Select("email, first_name, last_name, surname, directory_type").From("users")

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	err = ur.db.SelectContext(ctx, &users, sql, args...)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserSuggestionsByType(ctx context.Context, searchParams []string, userType string) ([]domain.User, error) {
	var users []domain.User

	builder := sq.Select("email, first_name, last_name, surname, directory_type").From("users").Where("directory_type = $1", userType)

	switch len(searchParams) {
	case 2:
		builder = builder.Where("(first_name LIKE '%' || $2 || '%' AND last_name LIKE '%' || $3 || '%')", searchParams[0], searchParams[1])
	case 1:
		builder = builder.Where("(first_name LIKE '%' || $2 || '%' OR last_name LIKE '%' || $2 || '%')", searchParams[0])
	}

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()

	ur.log.Info(fmt.Sprintf("query string: %v, params: %v", sql, args))
	err = ur.db.SelectContext(ctx, &users, sql, args...)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	builder := sq.Select("email, first_name, last_name, surname, directory_type").From("users").Where("email = $1", email)

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()

	ur.log.Info(fmt.Sprintf("query string: %v, params: %v", sql, args))
	err = ur.db.GetContext(ctx, &user, sql, args...)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
