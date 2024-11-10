package services

import (
	"context"
	"fmt"
	"log/slog"
	"reliab-test/internal/domain"
)

type VacancyService struct {
	vacancyRepository vacancyRepositoryInf
	log               *slog.Logger
}

func BuildVacancyService(
	log *slog.Logger,
	vacancyRepository vacancyRepositoryInf) *VacancyService {
	return &VacancyService{vacancyRepository: vacancyRepository, log: log}
}

func (vs *VacancyService) GetVacancyByID(ctx context.Context, vacancyId string) (*domain.Vacancy, error) {
	vacancy, err := vs.vacancyRepository.GetVacancyByID(ctx, vacancyId)

	if err != nil {
		return nil, fmt.Errorf("failed to get vacancy: %w", err)
	}

	return vacancy, nil
}

func (vs *VacancyService) GetVacancies(ctx context.Context, filter domain.VacancyFilter) ([]domain.Vacancy, error) {

	vacancies, err := vs.vacancyRepository.GetAllVacancies(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("failed to get all vacancies: %w", err)
	}

	return vacancies, nil
}
