package domain

import "time"

type ResumeVacancy struct {
	ResumeID      string `db:"resume_id"`
	VacancyID     string `db:"vacancy_id"`
	Status        string `db:"status"`
	ResumeStatus  string `db:"resume_status"`
	VacancyStatus string `db:"vacancy_status"`
	Notes         string `db:"notes"`
}

type ResumeVacancyChangeHistory struct {
	ID                    string    `db:"id"`
	CommunicationStatusID string    `db:"communication_status_id"`
	Date                  time.Time `db:"date"`
}
