package services

import (
	"context"
	"fmt"
	"log/slog"
	"reliab-test/internal/domain"
)

type ResumeService struct {
	resumeRepository resumeRepository
	log              *slog.Logger
}

func BuildResumeService(log *slog.Logger, resumeRepository resumeRepository) *ResumeService {
	return &ResumeService{resumeRepository: resumeRepository, log: log}
}

func (rs *ResumeService) GetResumeByID(ctx context.Context, resumeId int64) (*domain.Resume, error) {
	resume, err := rs.resumeRepository.GetResumeByID(ctx, resumeId)
	if err != nil {
		return nil, fmt.Errorf("Faild to get resume: %w", err)
	}

	return resume, nil
}

func (rs *ResumeService) GetAllResumes(ctx context.Context, filter domain.ResumeFilter) ([]domain.Resume, int, int, error) {
	resumes, totalResumes, err := rs.resumeRepository.GetAllResumes(ctx, filter)

	if err != nil {
		return nil, 0, 0, fmt.Errorf("Faild to get all resumes: %w", err)
	}

	totalPages := (totalResumes + filter.Limit - 1) / filter.Limit
	return resumes, totalResumes, totalPages, nil
}
