package services

import (
	"context"
	"log/slog"
	"reliab-test/internal/domain"
)

type ResumeVacancyService struct {
	resumeVacancyRepository ResumeVacancyRepository
	log                     *slog.Logger
}

func BuildResumeVacancyService(
	log *slog.Logger,
	resumeVacancyRepository ResumeVacancyRepository) *ResumeVacancyService {
	return &ResumeVacancyService{resumeVacancyRepository: resumeVacancyRepository, log: log}
}

// CreateLink создает новую связь между резюме и вакансией
func (s *ResumeVacancyService) CreateLink(ctx context.Context, request domain.ResumeVacancy) (string, error) {
	link := domain.ResumeVacancy{
		ResumeID:      request.ResumeID,
		VacancyID:     request.VacancyID,
		Status:        request.Status,
		ResumeStatus:  request.ResumeStatus,
		VacancyStatus: request.VacancyStatus,
		Notes:         request.Notes,
	}
	return s.resumeVacancyRepository.CreateLink(ctx, link)
}

// UpdateLinkStatus обновляет статус существующей связи между резюме и вакансией
func (s *ResumeVacancyService) UpdateLinkStatus(ctx context.Context, resumeID, vacancyID string, request domain.ResumeVacancy) error {
	return s.resumeVacancyRepository.UpdateLinkStatus(ctx, resumeID, vacancyID, request.Status, request.ResumeStatus, request.VacancyStatus, request.Notes)
}

// GetResumesByVacancy возвращает резюме, связанные с вакансией
func (s *ResumeVacancyService) GetResumesByVacancy(ctx context.Context, vacancyID string) ([]domain.ResumeVacancy, error) {
	links, err := s.resumeVacancyRepository.GetResumesByVacancy(ctx, vacancyID)
	if err != nil {
		return nil, err
	}

	var responses []domain.ResumeVacancy
	for _, link := range links {
		responses = append(responses, domain.ResumeVacancy{
			ResumeID:      link.ResumeID,
			VacancyID:     link.VacancyID,
			Status:        link.Status,
			ResumeStatus:  link.ResumeStatus,
			VacancyStatus: link.VacancyStatus,
			Notes:         link.Notes,
		})
	}
	return responses, nil
}

// GetVacanciesByResume возвращает вакансии, связанные с резюме
func (s *ResumeVacancyService) GetVacanciesByResume(ctx context.Context, resumeID string) ([]domain.ResumeVacancy, error) {
	links, err := s.resumeVacancyRepository.GetVacanciesByResume(ctx, resumeID)
	if err != nil {
		return nil, err
	}

	var responses []domain.ResumeVacancy
	for _, link := range links {
		responses = append(responses, domain.ResumeVacancy{
			VacancyID:     link.VacancyID,
			ResumeID:      link.ResumeID,
			Status:        link.Status,
			ResumeStatus:  link.ResumeStatus,
			VacancyStatus: link.VacancyStatus,
			Notes:         link.Notes,
		})
	}
	return responses, nil
}
