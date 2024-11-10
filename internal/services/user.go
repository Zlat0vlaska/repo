package services

import (
	"context"
	"log/slog"
	"reliab-test/internal/domain"
	"reliab-test/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
	log            *slog.Logger
}

func BuildUserService(userRepository *repositories.UserRepository, log *slog.Logger) *UserService {
	return &UserService{userRepository: userRepository, log: log}
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := us.userRepository.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetUsers(ctx context.Context, searchParams []string) ([]domain.User, error) {
	users, err := us.userRepository.GetUserSuggestionsByType(ctx, searchParams, domain.UserDirectoryType)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetApplicants(ctx context.Context, searchStr []string) ([]domain.User, error) {
	users, err := us.userRepository.GetUserSuggestionsByType(ctx, searchStr, domain.ApplicantDirectoryType)

	if err != nil {
		return nil, err
	}

	return users, nil
}
