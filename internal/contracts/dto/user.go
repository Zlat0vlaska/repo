package dto

import (
	_ "github.com/go-playground/validator/v10"
	"reliab-test/internal/domain"
	"strings"
)

type UserDTO struct {
	Type  string `json:"type" validate:"oneof=USER APPLICANT"`
	Email string `json:"email"`
	FIO   string `json:"fio"`
}

// Builders
func MapUsersToDto(users []domain.User) []UserDTO {
	usersDto := make([]UserDTO, 0, len(users))

	for _, user := range users {
		var userDto UserDTO

		userDto.Email = user.Email
		userDto.Type = user.DirectoryType
		userDto.FIO = getFio(user.FirstName, user.LastName, user.Surname)

		usersDto = append(usersDto, userDto)
	}

	return usersDto
}

func getFio(name string, lastName string, surname *string) string {
	var fioBuilder strings.Builder

	fioBuilder.WriteString(lastName)
	fioBuilder.WriteString(" ")
	fioBuilder.WriteString(name)

	if surname != nil {
		fioBuilder.WriteString(" ")
		fioBuilder.WriteString(*surname)
	}

	return fioBuilder.String()
}
