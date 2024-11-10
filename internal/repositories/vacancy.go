package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"reliab-test/internal/domain"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type VacancyRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func BuildVacancyRepository(db *sqlx.DB, log *slog.Logger) *VacancyRepository {
	return &VacancyRepository{db: db, log: log}
}

func (vr *VacancyRepository) GetAllVacancies(ctx context.Context, filter domain.VacancyFilter) ([]domain.Vacancy, error) {
	queryBuilder := sq.Select(
		"v.id",
		"v.name AS name",
		"v.company",
		"v.state",
		"v.job_tittle",
		"v.salary",
		"v.description",
		"string_agg(DISTINCT g.name, ',') AS grades",
		"string_agg(DISTINCT s.name, ',') AS skills",
		"c.name AS country",
		"string_agg(DISTINCT r.name, ',') AS regions",
		"string_agg(DISTINCT ci.name, ',') AS cities",
		"v.date_create",
		"v.is_favorite"). 
		From("vacancies v").
		Join("grades_vacancies gv ON v.id = gv.vacancy_id").
		Join("grades g ON g.id = gv.grade_id").
		Join("vacancies_skills vs ON v.id = vs.vacancy_id").
		Join("skills s ON s.id = vs.skill_id").
		Join("vacancies_cities vc ON v.id = vc.vacancy_id").
		Join("cities ci ON vc.city_id = ci.id").
		Join("regions r ON ci.region_id = r.id").
		Join("countries c ON r.country_id = c.id").
		GroupBy("v.id", "c.name").
		PlaceholderFormat(sq.Dollar)

	if filter.Keyword != "" {
		likePattern := fmt.Sprintf("%%%s%%", filter.Keyword)
		queryBuilder = queryBuilder.Where(sq.Or{
			sq.Like{"lower(v.name)": likePattern},
			sq.Like{"s.name": likePattern},
		})
	}

	if filter.Country != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"c.name": filter.Country})
	}

	if filter.Region != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"r.name": filter.Region})
	}

	if filter.City != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"ci.name": filter.City})
	}

	if filter.SalaryMin > 0 {
		queryBuilder = queryBuilder.Where(sq.GtOrEq{"v.salary": filter.SalaryMin})
	}
	if filter.SalaryMax > 0 {
		queryBuilder = queryBuilder.Where(sq.LtOrEq{"v.salary": filter.SalaryMax})
	}

	if filter.Grade != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"g.name": filter.Grade})
	}

	if filter.HrID != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"v.is_favorite": true, "v.hr_id": filter.HrID})
	}

	if filter.SortBy == "date_asc" {
		queryBuilder = queryBuilder.OrderBy("v.date_create ASC")
	} else {
		queryBuilder = queryBuilder.OrderBy("v.date_create DESC")
	}

	queryBuilder = queryBuilder.Limit(uint64(filter.Limit)).Offset(uint64(filter.Offset))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	vr.log.Info(fmt.Sprintf("query string: %v, params: %v", query, args))

	rows, err := vr.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []domain.Vacancy

	for rows.Next() {
		var (
			vacancy                         domain.Vacancy
			grades, skills, regions, cities sql.NullString
		)

		err := rows.Scan(
			&vacancy.ID,
			&vacancy.Name,
			&vacancy.Company,
			&vacancy.State,
			&vacancy.JobTitle,
			&vacancy.Salary,
			&vacancy.Description,
			&grades,
			&skills,
			&vacancy.Country,
			&regions,
			&cities,
			&vacancy.DateCreate,
			&vacancy.IsFavorite,
		)
		if err != nil {
			return nil, err
		}

		if grades.Valid {
			vacancy.Grades = strings.Split(grades.String, ",")
		} else {
			vacancy.Grades = []string{}
		}

		if skills.Valid {
			vacancy.Skills = strings.Split(skills.String, ",")
		} else {
			vacancy.Skills = []string{}
		}

		if regions.Valid {
			vacancy.Regions = strings.Split(regions.String, ",")
		} else {
			vacancy.Regions = []string{}
		}

		if cities.Valid {
			vacancy.Cities = strings.Split(cities.String, ",")
		} else {
			vacancy.Cities = []string{}
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (vr *VacancyRepository) GetVacancyByID(ctx context.Context, id string) (*domain.Vacancy, error) {
	query, args, err := sq.Select(
		"v.id",
		"v.name",
		"v.company",
		"v.state",
		"v.job_tittle",
		"v.salary",
		"v.description",
		"string_agg(DISTINCT g.name, ',') AS grades",
		"string_agg(DISTINCT s.name, ',') AS skills",
		"c.name AS country",
		"string_agg(DISTINCT r.name, ',') AS regions",
		"string_agg(DISTINCT ci.name, ',') AS cities",
		"v.date_create",
		"v.is_favorite").
		From("vacancies v").
		LeftJoin("grades_vacancies gv ON v.id = gv.vacancy_id").
		LeftJoin("grades g ON g.id = gv.grade_id").
		LeftJoin("vacancies_skills vs ON v.id = vs.vacancy_id").
		LeftJoin("skills s ON s.id = vs.skill_id").
		LeftJoin("vacancies_cities vc ON v.id = vc.vacancy_id").
		LeftJoin("cities ci ON vc.city_id = ci.id").
		LeftJoin("regions r ON ci.region_id = r.id").
		LeftJoin("countries c ON r.country_id = c.id").
		Where(sq.Eq{"v.id": id}).
		GroupBy("v.id", "c.name").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var (
		vacancy                         domain.Vacancy
		grades, skills, regions, cities sql.NullString
	)

	vr.log.Info(fmt.Sprintf("query string: %v, params: %v", query, args))

	err = vr.db.QueryRowxContext(ctx, query, args...).Scan(
		&vacancy.ID,
		&vacancy.Name,
		&vacancy.Company,
		&vacancy.State,
		&vacancy.JobTitle,
		&vacancy.Salary,
		&vacancy.Description,
		&grades,
		&skills,
		&vacancy.Country,
		&regions,
		&cities,
		&vacancy.DateCreate,
		&vacancy.IsFavorite,
	)
	if err != nil {
		return nil, err
	}

	if grades.Valid {
		vacancy.Grades = strings.Split(grades.String, ",")
	} else {
		vacancy.Grades = []string{}
	}

	if skills.Valid {
		vacancy.Skills = strings.Split(skills.String, ",")
	} else {
		vacancy.Skills = []string{}
	}

	if regions.Valid {
		vacancy.Regions = strings.Split(regions.String, ",")
	} else {
		vacancy.Regions = []string{}
	}

	if cities.Valid {
		vacancy.Cities = strings.Split(cities.String, ",")
	} else {
		vacancy.Cities = []string{}
	}

	return &vacancy, nil
}
