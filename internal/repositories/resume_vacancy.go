package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"reliab-test/internal/domain"

	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ResumeVacancyRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func BuildResumeVacancyRepository(db *sqlx.DB, log *slog.Logger) *ResumeVacancyRepository {
	return &ResumeVacancyRepository{db: db, log: log}
}

// CreateLink создает новую связь между резюме и вакансией
func (r *ResumeVacancyRepository) CreateLink(ctx context.Context, link domain.ResumeVacancy) (string, error) {
	query, args, err := sq.Insert("communication").
		Columns("resume_id", "vacancy_id", "communication_status", "resume_status", "vacancy_status", "name").
		Values(link.ResumeID, link.VacancyID, link.Status, link.ResumeStatus, link.VacancyStatus, link.Notes).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", err
	}

	var insertedID string
	err = r.db.QueryRowxContext(ctx, query, args...).Scan(&insertedID)
	if err != nil {
		return "", err
	}
	r.log.Info("Executing CreateLink query", "query", query, "args", args)
	return insertedID, nil
}

// UpdateLinkStatus обновляет статус существующей связи между резюме и вакансией
func (r *ResumeVacancyRepository) UpdateLinkStatus(ctx context.Context, resumeID, vacancyID, status, resumeStatus, vacancyStatus, notes string) error {
	updateQuery, updateArgs, err := sq.Update("communication").
		Set("status", status).
		Set("resume_status", resumeStatus).
		Set("vacancy_status", vacancyStatus).
		Where(sq.Eq{"resume_id": resumeID, "vacancy_id": vacancyID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	r.log.Info("Executing UpdateLinkStatus query", "query", updateQuery, "args", updateArgs)

	_, err = r.db.ExecContext(ctx, updateQuery, updateArgs...)
	if err != nil {
		return err
	}

	historyQuery, historyArgs, err := sq.Insert("change_history").
		Columns("resume_id", "vacancy_id", "status", "notes").
		Values(resumeID, vacancyID, status, notes).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	// Логирование запроса на добавление в историю изменений
	r.log.Info("Executing history insert query", "query", historyQuery, "args", historyArgs)

	_, err = r.db.ExecContext(ctx, historyQuery, historyArgs...)
	return err
}

// GetChangeHistory возвращает историю изменений статуса связи между резюме и вакансиями
func (r *ResumeVacancyRepository) GetChangeHistory(ctx context.Context, resumeID, vacancyID string) ([]domain.ResumeVacancyChangeHistory, error) {
	var history []domain.ResumeVacancyChangeHistory
	query, args, err := sq.Select("*").
		From("change_history").
		Where(sq.Eq{"resume_id": resumeID, "vacancy_id": vacancyID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err := r.db.SelectContext(ctx, &history, query, args...); err != nil {
		return nil, err
	}
	return history, nil
}
func isEmptyUUID(id string) bool {
	return id == ""
}
func (r *ResumeVacancyRepository) GetResumesByVacancy(ctx context.Context, vacancyID string) ([]domain.ResumeVacancy, error) {
	uuivacancyID, err := uuid.Parse(vacancyID)
	if err != nil {
		return nil, fmt.Errorf("invalid resumeID format: %v", err)
	}
	if isEmptyUUID(vacancyID) {
		return nil, errors.New("resumeID is empty")
	}

	r.log.Info("Checking vacancyID for GetResumesByVacancy", "vacancyID", vacancyID)
	if isEmptyUUID(vacancyID) {
		return nil, errors.New("vacancyID is empty")
	}
	query, args, err := sq.Select(
		"c.resume_id",
		"c.vacancy_id",
		"cs.name AS communication_status",
		"s.name AS statuses",
		"v.state AS vacancies",
		"c.name").
		From("communication c").
		LeftJoin("communication_status cs ON c.id = cs.communication_id").
		LeftJoin("resumes r ON c.resume_id = r.id").
		LeftJoin("statuses s ON s.id = r.status_id").
		LeftJoin("vacancies v ON c.vacancy_id = v.id").
		Where(sq.Eq{"c.vacancy_id": uuivacancyID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	r.log.Info("Executing GetResumesByVacancy query", "query", query, "args", args)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.ResumeVacancy
	for rows.Next() {
		var (
			link                                      domain.ResumeVacancy
			communication_status, statuses, vacancies sql.NullString
		)

		err := rows.Scan(
			&link.ResumeID,
			&link.VacancyID,
			&communication_status,
			&statuses,
			&vacancies,
			&link.Notes,
		)

		if err != nil {
			return nil, err
		}
		results = append(results, link)
	}

	return results, nil
}

func (r *ResumeVacancyRepository) GetVacanciesByResume(ctx context.Context, resumeID string) ([]domain.ResumeVacancy, error) {

	// Построение запроса с подзапросами и соединениями
	query, args, err := sq.Select(
		"r.profession AS resumes",
		"v.name AS vacancies",
		"cs.name AS communication_status",
		"s.name AS statuses",
		"v.state AS vacancies",
		"c.name AS notes").
		From("communication c").
		LeftJoin("resumes r ON r.id = c.resume_id").
		LeftJoin("vacancies v ON v.id = c.vacancy_id").
		LeftJoin("communication_status cs ON cs.communication_id = c.id").
		LeftJoin("statuses s ON s.id = cs.name").
		LeftJoin("vacancies v ON c.vacancy_id = v.id").
		Where(sq.Eq{"c.resume": resumeID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	r.log.Info("Executing GetVacanciesByResume query", "query", query, "args", args)

	// Выполнение запроса
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		r.log.Info("errrrrrr", err)

		return nil, err
	}
	defer rows.Close()

	// Инициализация результата
	var results []domain.ResumeVacancy
	for rows.Next() {
		var (
			link                                      domain.ResumeVacancy
			communication_status, statuses, vacancies sql.NullString
		)

		// Извлечение данных с обработкой NULL значений
		err = r.db.QueryRowxContext(ctx, query, args...).Scan(
			&link.ResumeID,
			&link.VacancyID,
			&communication_status,
			&statuses,
			&vacancies,
			&link.Notes,
		)
		if err != nil {
			return nil, err
		}

		// Присваиваем извлеченные данные в структуру `link`
		link.Status = communication_status.String
		link.ResumeStatus = statuses.String
		link.VacancyStatus = vacancies.String

		// Добавление в результат
		results = append(results, link)
	}

	// Проверка на ошибки после завершения сканирования
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
