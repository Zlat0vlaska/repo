package dto

import (
	"reliab-test/internal/domain"
)

type LinkResumeVacancyRequest struct {
	Status        string `json:"status" binding:"required"`
	ResumeStatus  string `json:"resume_status"`
	VacancyStatus string `json:"vacancy_status"`
	Notes         string `json:"notes"`
}

type LinkResumeVacancyResponse struct {
	Message       string `json:"message"`
	Link          string `json:"link"`
	ResumeStatus  string `json:"resume_status"`
	VacancyStatus string `json:"vacancy_status"`
}

type CreateLinkRequest struct {
	ResumeID      string `json:"resume_id" binding:"required"`
	VacancyID     string `json:"vacancy_id" binding:"required"`
	Status        string `json:"status" binding:"required"`
	ResumeStatus  string `json:"resume_status"`
	VacancyStatus string `json:"vacancy_status"`
	Notes         string `json:"notes"`
}
type UpdateLinkStatusRequest struct {
	Status        string `json:"status" binding:"required"`
	ResumeStatus  string `json:"resume_status"`
	VacancyStatus string `json:"vacancy_status"`
	Notes         string `json:"notes"`
}

func BuildDtoToLink(dto CreateLinkRequest) domain.ResumeVacancy {

	return domain.ResumeVacancy{
		ResumeID:      dto.ResumeID,
		VacancyID:     dto.VacancyID,
		Status:        dto.Status,
		ResumeStatus:  dto.ResumeStatus,
		VacancyStatus: dto.VacancyStatus,
		Notes:         dto.Notes,
	}
}
