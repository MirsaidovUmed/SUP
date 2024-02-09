package service

import (
	"SUP/internal/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UpdateUser(user models.User) (err error) {
	existingUser, err := s.Repo.GetUserIdByEmail(user.Email)
	if err != nil {
		fmt.Println(2)
		return err
	}

	user.Id = existingUser.Id
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}

	if user.SecondName != "" {
		existingUser.SecondName = user.SecondName
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}

	roleId, err := s.Repo.GetRoleByName(user.Role.Name)
	if err != nil {
		fmt.Println(1)
		return err
	}

	existingUser.Role.Id = roleId.Id

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUser.Password = string(hashedPassword)
	}

	err = s.Repo.UpdateUser(existingUser)

	return err
}
