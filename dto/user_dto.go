package dto

import "bjzdgt.com/ants/model"

//UserDto is dto user
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// UserToDto func for user to UserDto
func UserToDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
