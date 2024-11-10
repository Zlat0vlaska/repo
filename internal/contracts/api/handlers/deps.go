package handlers

import (
	"context"
	"reliab-test/internal/domain"
)

type vacancyServiceInf interface {
	GetVacancyByID(ctx context.Context, vacancyId string) (*domain.Vacancy, error)
	GetVacancies(ctx context.Context, filter domain.VacancyFilter) ([]domain.Vacancy, error)
}

type resumeService interface {
	GetResumeByID(ctx context.Context, resumeId int64) (*domain.Resume, error)
	GetAllResumes(ctx context.Context, filter domain.ResumeFilter) ([]domain.Resume, int, int, error)
}

type ResumeVacancyService interface {
	CreateLink(ctx context.Context, link domain.ResumeVacancy) (string, error)
	UpdateLinkStatus(ctx context.Context, resumeID, vacancyID string, request domain.ResumeVacancy) error
	GetResumesByVacancy(ctx context.Context, vacancyID string) ([]domain.ResumeVacancy, error)
	GetVacanciesByResume(ctx context.Context, resumeID string) ([]domain.ResumeVacancy, error)
}
