package dto

import (
	"fmt"
	"github.com/google/uuid"
	"reliab-test/internal/domain"
	"time"
)

type MeetingDTO struct {
	ID              string    `json:"id,omitempty"`
	Name            string    `json:"name"`
	Place           string    `json:"place"`
	Comment         *string   `json:"comment,omitempty"`
	RecipientEmails *[]string `json:"recipient_emails,omitempty"`
	ApplicantEmail  string    `json:"applicant_email"`
	StartDate       string    `json:"start_date"` // dd.mm.yyyy
	StartTime       string    `json:"start_time"` // hh:mm
	EndDate         string    `json:"end_date"`   // dd.mm.yyyy
	EndTime         string    `json:"end_time"`   // hh:mm
	IsFullDay       bool      `json:"is_full_day"`
	IsOnline        bool      `json:"is_online"`
	AuthorEmail     string    `json:"author_email"`
}

// requests
type CreateMeetingRequest struct {
	MeetingDTO
}

type EditMeetingRequest struct {
	MeetingDTO
}

type GetMeetingResponse struct {
	MeetingDTO
}

type CreateMeetingResponse struct {
	ID string `json:"uuid"`
}

// builders
func BuildDtoToMeetingMultiple(dto []MeetingDTO) ([]domain.Meeting, error) {
	meetings := make([]domain.Meeting, 0, len(dto))

	for _, data := range dto {
		meeting, err := BuildDtoToMeeting(data)
		if err != nil {
			return []domain.Meeting{}, err
		}

		meetings = append(meetings, meeting)
	}

	return meetings, nil
}

func BuildMeetingToDtoMultiple(meetings []domain.Meeting) []MeetingDTO {
	dto := make([]MeetingDTO, 0, len(meetings))

	for _, meeting := range meetings {
		data := BuildMeetingToDto(meeting)
		dto = append(dto, data)
	}
	return dto
}

func BuildDtoToMeeting(dto MeetingDTO) (domain.Meeting, error) {
	const (
		dateLayout = "02.01.2006"
		timeLayout = "15:04"
	)

	startDate, err := time.Parse(dateLayout, dto.StartDate)
	if err != nil {
		return domain.Meeting{}, fmt.Errorf("invalid start date: %w", err)
	}

	startTime, err := time.Parse(timeLayout, dto.StartTime)
	if err != nil {
		return domain.Meeting{}, fmt.Errorf("invalid start time: %w", err)
	}

	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(),
		startTime.Hour(), startTime.Minute(), 0, 0, startDate.Location())

	endDate, err := time.Parse(dateLayout, dto.EndDate)
	if err != nil {
		return domain.Meeting{}, fmt.Errorf("invalid end date: %w", err)
	}

	endTime, err := time.Parse(timeLayout, dto.EndTime)
	if err != nil {
		return domain.Meeting{}, fmt.Errorf("invalid end time: %w", err)
	}

	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(),
		endTime.Hour(), endTime.Minute(), 0, 0, endDate.Location())

	return domain.Meeting{
		ID:              uuid.New().String(),
		Name:            dto.Name,
		Place:           dto.Place,
		Comment:         dto.Comment,
		RecipientEmails: *dto.RecipientEmails,
		ApplicantEmail:  dto.ApplicantEmail,
		StartDate:       start,
		EndDate:         end,
		IsFullDay:       dto.IsFullDay,
		IsOnline:        dto.IsOnline,
		AuthorEmail:     dto.AuthorEmail,
	}, nil
}

func BuildMeetingToDto(meeting domain.Meeting) MeetingDTO {
	startDate := meeting.StartDate.Format("02.01.2006")
	startTime := meeting.StartDate.Format("15:04")
	endDate := meeting.EndDate.Format("02.01.2006")
	endTime := meeting.EndDate.Format("15:04")

	return MeetingDTO{
		ID:              meeting.ID,
		Name:            meeting.Name,
		Place:           meeting.Place,
		Comment:         meeting.Comment,
		RecipientEmails: &meeting.RecipientEmails,
		ApplicantEmail:  meeting.ApplicantEmail,
		StartDate:       startDate,
		StartTime:       startTime,
		EndDate:         endDate,
		EndTime:         endTime,
		IsFullDay:       meeting.IsFullDay,
		IsOnline:        meeting.IsOnline,
		AuthorEmail:     meeting.AuthorEmail,
	}
}
