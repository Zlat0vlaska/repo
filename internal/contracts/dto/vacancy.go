package dto

import (
	"reliab-test/internal/domain"
	"strconv"
	"strings"
)

// requests
type GetVacanciesRequest struct {
	Keyword   string `form:"keyword"`
	Country   string `form:"country"`
	Region    string `form:"region"`
	City      string `form:"city"`
	Grade     string `form:"grade"`
	SortBy    string `form:"sort"`
	HrID      string `form:"hr_id"`
	SalaryMin string `form:"salary_min"`
	SalaryMax string `form:"salary_max"`
	Page      string `form:"page"`
	Limit     string `form:"limit"`
}

// responses
type GetVacancyResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Company     string   `json:"company"`
	State       string   `json:"state"`
	JobTitle    string   `json:"job_title"`
	Salary      string   `json:"salary"`
	Description string   `json:"description"`
	Grades      []string `json:"grades"`
	DateCreate  string   `json:"date_create"`
	IsFavorite  bool     `json:"is_favorite"`
	Skills      []string `json:"skills"`
	Country     string   `json:"country"`
	Regions     []string `json:"regions"`
	Cities      []string `json:"cities"`
}

// builders
func BuildGetVacanciesDtoToFilter(dto GetVacanciesRequest) domain.VacancyFilter {
	salaryMin, _ := strconv.Atoi(dto.SalaryMin)
	salaryMax, _ := strconv.Atoi(dto.SalaryMax)

	page, err := strconv.Atoi(dto.Page)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(dto.Limit)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	sortBy := "date_desc"
	if dto.SortBy != "" {
		sortBy = "date_asc"
	}

	return domain.VacancyFilter{
		Keyword:   strings.ToLower(dto.Keyword),
		Country:   dto.Country,
		Region:    dto.Region,
		City:      dto.City,
		Grade:     dto.Grade,
		SortBy:    sortBy,
		HrID:      dto.HrID,
		SalaryMin: salaryMin,
		SalaryMax: salaryMax,
		Limit:     limit,
		Offset:    offset,
	}
}

func BuildVacancyToGetVacancyResponseMultiple(vacancies []domain.Vacancy) []GetVacancyResponse {
	dto := make([]GetVacancyResponse, 0, len(vacancies))

	for _, vacancy := range vacancies {
		dto = append(dto, BuildVacancyToGetVacancyResponse(vacancy))
	}

	return dto
}

func BuildVacancyToGetVacancyResponse(vacancy domain.Vacancy) GetVacancyResponse {
	return GetVacancyResponse{
		ID:          vacancy.ID,
		Name:        vacancy.Name,
		Company:     vacancy.Company,
		State:       vacancy.State,
		JobTitle:    vacancy.JobTitle,
		Salary:      vacancy.Salary,
		Description: vacancy.Description,
		Grades:      vacancy.Grades,
		DateCreate:  vacancy.DateCreate,
		IsFavorite:  vacancy.IsFavorite,
		Skills:      vacancy.Skills,
		Country:     vacancy.Country,
		Regions:     vacancy.Regions,
		Cities:      vacancy.Cities,
	}
}
