package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"reliab-test/internal/domain"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ResumeRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func BuildResumeRepository(db *sqlx.DB, log *slog.Logger) *ResumeRepository {
	return &ResumeRepository{db: db, log: log}
}

func (rr *ResumeRepository) GetAllResumes(ctx context.Context, filter domain.ResumeFilter) ([]domain.Resume, int, error) {
	validSortFields := map[string]string{
		"name":       "c.name",
		"profession": "r.profession",
		//"experience": "r.experience_years",
	}

	sortColumn, ok := validSortFields[filter.SortBy]
	if !ok {
		sortColumn = "c.name"
	}

	if filter.SortOrder != "asc" && filter.SortOrder != "desc" {
		filter.SortOrder = "asc"
	}

	queryBuilder := sq.Select(
		"r.id",
		"c.name AS candidate_name",
		"s.name AS status",
		"g.name AS grade",
		"cn.name AS country",
		"ci.name AS city",
		"r.profession",
		"r.schedule",
		"r.citizenship",
		"r.business_trips_readiness",
		"r.permission",
		"r.salary",
		"r.relocation_readiness",
		//"r.experience_years",
		"string_agg(sk.name, ',') AS skills",
		"string_agg(l.name, ',') AS languages",
		"string_agg(ed.description, ',') AS educations",
		//"ed.description AS education",
	).
		From("resumes r").
		LeftJoin("candidates c ON r.candidate_id = c.id").
		LeftJoin("statuses s ON r.status_id = s.id").
		LeftJoin("grades g ON r.grade_id = g.id").
		LeftJoin("countries cn ON r.country_id = cn.id").
		LeftJoin("resumes_cities rc ON r.id = rc.resume_id").
		LeftJoin("cities ci ON ci.id = rc.city_id").
		LeftJoin("resumes_educations red ON r.id = red.resume_id").
		LeftJoin("educations ed ON ed.id = red.education_id").
		LeftJoin("resumes_skills rs ON r.id = rs.resume_id").
		LeftJoin("skills sk ON sk.id = rs.skill_id").
		LeftJoin("resumes_languages rl ON r.id = rl.resume_id").
		LeftJoin("languages l ON l.id = rl.languages_id").
		GroupBy("r.id, c.name, s.name, g.name, cn.name, ci.name").
		PlaceholderFormat(sq.Dollar).
		Offset(uint64(filter.Offset)).
		OrderBy(fmt.Sprintf("%s %s", sortColumn, filter.SortOrder))

	if filter.Keyword != "" {
		likePattern := fmt.Sprintf("%%%s%%", filter.Keyword)
		queryBuilder = queryBuilder.Where(sq.Or{
			sq.Like{"r.profession": likePattern},
			sq.Like{"sk.name": likePattern},
			sq.Like{"r.permission": likePattern},
		})
	}

	if filter.Education != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"ed.educations": filter.Education})
	}

	if filter.Skills != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"s.skills": filter.Skills})
	}

	if filter.Languages != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.languages": filter.Languages})
	}

	if filter.Schedule != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.schedule": filter.Schedule})
	}

	if filter.RelocationReadiness != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.relocation_readiness": filter.RelocationReadiness})
	}

	if filter.Citizenship != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.citizenship": filter.Citizenship})
	}

	if filter.BusinessTripsReadiness != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.business_trips_readiness": filter.BusinessTripsReadiness})
	}
	if filter.Profession != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.profession": filter.Profession})
	}

	if filter.Languages != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.languages": filter.Languages})
	}

	if filter.SalaryMin > 0 {
		queryBuilder = queryBuilder.Where(sq.GtOrEq{"r.salary": filter.SalaryMin})
	}
	if filter.SalaryMax > 0 {
		queryBuilder = queryBuilder.Where(sq.LtOrEq{"r.salary": filter.SalaryMax})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, 0, err
	}

	rr.log.Info(fmt.Sprintf("query string: %v, params: %v", query, args))

	rows, err := rr.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var resumes []domain.Resume

	for rows.Next() {
		var (
			resume                                domain.Resume
			cities, skills, languages, educations sql.NullString
		)

		err := rows.Scan(
			&resume.ID,
			&resume.CandidateName,
			&resume.Status,
			&resume.Grade,
			&resume.Country,
			&cities,
			&resume.Profession,
			&resume.Schedule,
			&resume.Citizenship,
			&resume.BusinessTripsReadiness,
			&resume.Permission,
			&resume.Salary,
			&resume.RelocationReadiness,
			//&resume.ExperienceYears,
			&skills,
			&languages,
			&educations,
		)
		if err != nil {
			return nil, 0, err
		}

		if skills.Valid {
			resume.Skills = strings.Split(skills.String, ",")
		} else {
			resume.Skills = []string{}
		}

		if languages.Valid {
			resume.Languages = strings.Split(languages.String, ",")
		} else {
			resume.Languages = []string{}
		}

		resumes = append(resumes, resume)
	}
	totalResumes, err := rr.GetTotalResumeCount(ctx, filter.Keyword)
	if err != nil {
		return nil, 0, err
	}
	return resumes, totalResumes, nil
}
func (rr *ResumeRepository) GetTotalResumeCount(ctx context.Context, keyword string) (int, error) {
	queryBuilder := sq.Select("COUNT(*)").
		From("resumes r").
		LeftJoin("resumes_skills rs ON r.id = rs.resume_id").
		LeftJoin("skills sk ON sk.id = rs.skill_id").
		LeftJoin("resumes_languages rl ON r.id = rl.resume_id").
		LeftJoin("languages l ON l.id = rl.languages_id").
		PlaceholderFormat(sq.Dollar)

	if keyword != "" {
		likePattern := fmt.Sprintf("%%%s%%", keyword)
		queryBuilder = queryBuilder.Where(sq.Or{
			sq.Like{"r.profession": likePattern},
			sq.Like{"sk.name": likePattern},
			sq.Like{"r.permission": likePattern},
		})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var totalResumes int
	err = rr.db.GetContext(ctx, &totalResumes, query, args...)
	if err != nil {
		return 0, err
	}

	return totalResumes, nil
}

func (rr *ResumeRepository) GetResumeByID(ctx context.Context, id int64) (*domain.Resume, error) {
	query, args, err := sq.Select(
		"r.id",
		"c.name AS candidate_name",
		"s.name AS status",
		"g.name AS grade",
		"cn.name AS country",
		"ci.name AS city",
		"r.profession",
		"r.schedule",
		"r.citizenship",
		"r.business_trips_readiness",
		"r.permission",
		"r.salary",
		"r.relocation_readiness",
		//"r.experience_years",
		"string_agg(sk.name, ',') AS skills",
		"string_agg(l.name, ',') AS languages",
		"string_agg(ed.description, ',') AS educations",
		//"ed.description AS education",
	).
		From("resumes r").
		Join("candidates c ON r.candidate_id = c.id").
		Join("statuses s ON r.status_id = s.id").
		Join("grades g ON r.grade_id = g.id").
		Join("countries cn ON r.country_id = cn.id").
		Join("resumes_cities rc ON r.id = rc.resume_id").
		Join("cities ci ON ci.id = rc.city_id").
		Join("resumes_educations red ON r.id = red.resume_id").
		Join("educations ed ON ed.id = red.education_id").
		Join("resumes_skills rs ON r.id = rs.resume_id").
		Join("skills sk ON sk.id = rs.skill_id").
		Join("resumes_languages rl ON r.id = rl.resume_id").
		Join("languages l ON l.id = rl.languages_id").
		Where(sq.Eq{"r.id": id}).
		GroupBy("r.id, c.name, s.name, g.name, cn.name, ci.name").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	var (
		resume                                domain.Resume
		cities, skills, languages, educations sql.NullString
	)

	rr.log.Info(fmt.Sprintf("query string: %v, params: %v", query, args))

	err = rr.db.QueryRowxContext(ctx, query, args...).Scan(
		&resume.ID,
		&resume.CandidateName,
		&resume.Status,
		&resume.Grade,
		&resume.Country,
		&cities,
		&resume.Profession,
		&resume.Schedule,
		&resume.Citizenship,
		&resume.BusinessTripsReadiness,
		&resume.Permission,
		&resume.Salary,
		&resume.RelocationReadiness,
		//&resume.ExperienceYears,
		&skills,
		&languages,
		&educations,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if skills.Valid {
		resume.Skills = strings.Split(skills.String, ",")
	} else {
		resume.Skills = []string{}
	}

	if languages.Valid {
		resume.Languages = strings.Split(languages.String, ",")
	} else {
		resume.Languages = []string{}
	}

	return &resume, nil
}
