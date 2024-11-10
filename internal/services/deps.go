package services

import (
	"context"
	"reliab-test/internal/domain"
)

type vacancyRepositoryInf interface {
	GetAllVacancies(ctx context.Context, filter domain.VacancyFilter) ([]domain.Vacancy, error)
	GetVacancyByID(ctx context.Context, id string) (*domain.Vacancy, error)
}

type resumeRepository interface {
	GetAllResumes(ctx context.Context, filter domain.ResumeFilter) ([]domain.Resume, int, error)
	GetResumeByID(ctx context.Context, id int64) (*domain.Resume, error)
}

type ResumeVacancyRepository interface {
	CreateLink(ctx context.Context, link domain.ResumeVacancy) (string, error)
	UpdateLinkStatus(ctx context.Context, resumeID, vacancyID string, status, resumeStatus, vacancyStatus, notes string) error
	GetChangeHistory(ctx context.Context, resumeID, vacancyID string) ([]domain.ResumeVacancyChangeHistory, error)

	GetResumesByVacancy(ctx context.Context, vacancyID string) ([]domain.ResumeVacancy, error)
	GetVacanciesByResume(ctx context.Context, resumeID string) ([]domain.ResumeVacancy, error)
}
