package repositories

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"reliab-test/internal/domain"
	"strings"
)

type MeetingRepository struct {
	db *sqlx.DB
}

func BuildMeetingRepository(db *sqlx.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

func (mr *MeetingRepository) CreateMeeting(ctx context.Context, meeting domain.Meeting) (string, error) {
	query, args, err := sq.Insert("meetings").
		Columns(
			"name",
			"place",
			"comment",
			"recipient_emails",
			"applicant_email",
			"start_date",
			"end_date",
			"is_full_day",
			"is_online",
			"author_email").
		Values(
			meeting.Name,
			meeting.Place,
			meeting.Comment,
			meeting.RecipientEmails,
			meeting.ApplicantEmail,
			meeting.StartDate,
			meeting.EndDate,
			meeting.IsFullDay,
			meeting.IsOnline,
			meeting.AuthorEmail).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var insertedID string
	err = mr.db.QueryRowxContext(ctx, query, args...).Scan(&insertedID)
	if err != nil {
		return "", err
	}

	return insertedID, nil
}

func (mr *MeetingRepository) GetMeetings(ctx context.Context) ([]domain.Meeting, error) {
	rows, err := mr.db.QueryxContext(ctx, `
        SELECT
            id,
            name,
            place,
            comment,
            recipient_emails,
            applicant_email,
            start_date,
            end_date,
            is_full_day,
            is_online,
            author_email
        FROM meetings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []domain.Meeting

	for rows.Next() {
		var meeting domain.Meeting
		var recipientEmails string

		err := rows.Scan(
			&meeting.ID,
			&meeting.Name,
			&meeting.Place,
			&meeting.Comment,
			&recipientEmails,
			&meeting.ApplicantEmail,
			&meeting.StartDate,
			&meeting.EndDate,
			&meeting.IsFullDay,
			&meeting.IsOnline,
			&meeting.AuthorEmail,
		)
		if err != nil {
			return nil, err
		}

		meeting.RecipientEmails = parseRecipientEmails(recipientEmails)

		meetings = append(meetings, meeting)
	}

	return meetings, nil
}

func (mr *MeetingRepository) GetMeetingByID(ctx context.Context, id string) (*domain.Meeting, error) {
	row := mr.db.QueryRowxContext(ctx, `
        SELECT
            id,
            name,
            place,
            comment,
            recipient_emails,
            applicant_email,
            start_date,
            end_date,
            is_full_day,
            is_online,
            author_email
        FROM meetings
        WHERE id = $1`, id)

	var meeting domain.Meeting
	var recipientEmails string

	// Выполняем сканирование строки
	err := row.Scan(
		&meeting.ID,
		&meeting.Name,
		&meeting.Place,
		&meeting.Comment,
		&recipientEmails,
		&meeting.ApplicantEmail,
		&meeting.StartDate,
		&meeting.EndDate,
		&meeting.IsFullDay,
		&meeting.IsOnline,
		&meeting.AuthorEmail,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Возвращаем nil, если митинг с таким ID не найден
		}
		return nil, err
	}

	meeting.RecipientEmails = parseRecipientEmails(recipientEmails)

	return &meeting, nil
}

func (mr *MeetingRepository) DeleteMeetingByID(ctx context.Context, id string) error {
	builder := sq.Delete("meetings").
		Where(sq.Eq{"id": id})

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = mr.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

// remove postgres curly braces from array
func parseRecipientEmails(recipientEmails string) []string {
	recipientEmails = strings.Trim(recipientEmails, "{}")
	if recipientEmails == "" {
		return []string{}
	}
	return strings.Split(recipientEmails, ",")
}
