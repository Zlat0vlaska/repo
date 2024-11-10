package domain

const (
	UserDirectoryType      = "USER"
	ApplicantDirectoryType = "APPLICANT"
)

type User struct {
	ID            int64   `db:"id"`
	Email         string  `db:"email"`
	FirstName     string  `db:"first_name"`
	LastName      string  `db:"last_name"`
	Surname       *string `db:"surname"`
	DirectoryType string  `db:"directory_type"`
}
