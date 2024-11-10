package dto

import (
	"reliab-test/internal/domain"
	"strconv"
	"strings"
)

type ResumeDTO struct {
	ID                     string   `json:"id"`
	CandidateName          string   `json:"candidate_name"`
	Profession             string   `json:"profession"`
	Grade                  string   `json:"grade"`
	Country                string   `json:"country"`
	City                   string   `json:"city"`
	Schedule               string   `json:"schedule"`
	Citizenship            string   `json:"citizenship"`
	BusinessTripsReadiness string   `json:"business_trips_readiness"`
	Permission             string   `json:"permission"`
	Salary                 float64  `json:"salary"`
	RelocationReadiness    string   `json:"relocation_readiness"`
	ExperienceYears        int      `json:"experience_years"`
	Skills                 []string `json:"skills"`
	Languages              []string `json:"languages"`
	Status                 string   `json:"status"`
	Education              string   `json:"education"`
}

type GetAllResumes struct {
	Keyword                string `form:"query"`
	Education              string `form:"education"`
	Skills                 string `form:"skills"`
	Languages              string `form:"languages"`
	Schedule               string `form:"schedule"`
	RelocationReadiness    string `form:"relocation_readiness"`
	Citizenship            string `form:"citizenship"`
	BusinessTripsReadiness string `form:"business_trips_readiness"`
	Profession             string `form:"profession"`
	SalaryMin              string `form:"salary_min"`
	SalaryMax              string `form:"salary_max"`
	Limit                  string `form:"limit"`
	Page                   string `form:"page"`
	Offset                 int    `form:"offset"`
	SortBy                 string `form:"sort"`
	SortOrder              string `form:"order"`
}

func GetResumesDtoToFilter(dto GetAllResumes) domain.ResumeFilter {

	page, err := strconv.Atoi(dto.Page)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(dto.Limit)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	sortBy := dto.SortBy

	sortOrder := "date_desc"
	if dto.SortOrder != "" {
		sortOrder = "date_asc"
	}

	salaryMin, _ := strconv.Atoi(dto.SalaryMin)
	salaryMax, _ := strconv.Atoi(dto.SalaryMax)

	return domain.ResumeFilter{
		Keyword:                strings.ToLower(dto.Keyword),
		Profession:             dto.Profession,
		Schedule:               dto.Schedule,
		Citizenship:            dto.Citizenship,
		BusinessTripsReadiness: dto.BusinessTripsReadiness,
		SalaryMin:              salaryMin,
		SalaryMax:              salaryMax,
		RelocationReadiness:    dto.RelocationReadiness,
		Skills:                 dto.Skills,
		Languages:              dto.Languages,
		Education:              dto.Education,
		Limit:                  limit,
		Offset:                 offset,
		SortBy:                 sortBy,
		SortOrder:              sortOrder,
	}
}
func BuildResumeToGetResponseMultiple(resumes []domain.Resume) []ResumeDTO {
	dto := make([]ResumeDTO, 0, len(resumes))

	for _, resume := range resumes {
		dto = append(dto, BuildResumeToGetResponse(resume))
	}

	return dto
}

func BuildResumeToGetResponse(resume domain.Resume) ResumeDTO {
	return ResumeDTO{
		ID:                     resume.ID,
		CandidateName:          resume.CandidateName,
		Profession:             resume.Profession,
		Grade:                  resume.Grade,
		Country:                resume.Country,
		City:                   resume.City,
		Schedule:               resume.Schedule,
		Citizenship:            resume.Citizenship,
		BusinessTripsReadiness: resume.BusinessTripsReadiness,
		Permission:             resume.Permission,
		Salary:                 resume.Salary,
		RelocationReadiness:    resume.RelocationReadiness,
		ExperienceYears:        resume.ExperienceYears,
		Skills:                 resume.Skills,
		Languages:              resume.Languages,
		Status:                 resume.Status,
		Education:              resume.Education,
	}
}
