package services

import (
	"context"
	"fmt"
	"log/slog"
	"reliab-test/internal/contracts/domain_errors"
	"reliab-test/internal/domain"
	"reliab-test/internal/repositories"
)

type MeetingService struct {
	meetingRepository *repositories.MeetingRepository
	usersRepository   *repositories.UserRepository
	log               *slog.Logger
}

func BuildMeetingService(
	log *slog.Logger,
	meetingRepository *repositories.MeetingRepository,
	userRepository *repositories.UserRepository) *MeetingService {
	return &MeetingService{meetingRepository: meetingRepository, usersRepository: userRepository, log: log}
}

func (ms *MeetingService) CreateMeeting(ctx context.Context, meeting domain.Meeting) (string, error) {
	recipientEmails := meeting.RecipientEmails
	for _, email := range recipientEmails {
		recipient, err := ms.usersRepository.GetUserByEmail(ctx, email)

		if err != nil {
			return "", fmt.Errorf("%v: %v :%w", domain_errors.UserNotExist, email, err)
		}

		if recipient.DirectoryType != domain.UserDirectoryType {
			return "", fmt.Errorf("%v: %v :%w", domain_errors.NotHomie, recipient, err)
		}
	}

	applicantEmail := meeting.ApplicantEmail
	applicant, err := ms.usersRepository.GetUserByEmail(ctx, applicantEmail)
	if err != nil {
		return "", fmt.Errorf("%v: %v :%w", domain_errors.UserNotExist, applicantEmail, err)
	}
	if applicant.DirectoryType != domain.ApplicantDirectoryType {
		return "", fmt.Errorf("%v: %v :%w", domain_errors.NotApplicant, applicant, err)
	}

	authorEmail := meeting.AuthorEmail
	author, err := ms.usersRepository.GetUserByEmail(ctx, authorEmail)
	if err != nil {
		return "", fmt.Errorf("%v: %v :%w", domain_errors.UserNotExist, authorEmail, err)
	}

	if author.DirectoryType != domain.UserDirectoryType {
		return "", fmt.Errorf("%v: %v :%w", domain_errors.NotHomie, author, err)
	}

	id, err := ms.meetingRepository.CreateMeeting(ctx, meeting)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (ms *MeetingService) GetMeetings(ctx context.Context) ([]domain.Meeting, error) {
	meetings, err := ms.meetingRepository.GetMeetings(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get meetings: %w", err)
	}

	return meetings, nil
}

func (ms *MeetingService) GetMeetingByID(ctx context.Context, id string) (*domain.Meeting, error) {
	meeting, err := ms.meetingRepository.GetMeetingByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get meeting by ID: %w", err)
	}

	return meeting, nil
}

func (ms *MeetingService) DeleteMeetingByID(ctx context.Context, id string) error {
	err := ms.meetingRepository.DeleteMeetingByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete meeting by ID: %w", err)
	}

	return nil
}
