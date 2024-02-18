package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UpdateUser(user models.User) (err error) {
	if user.Role.Name != "" {
		return errors.ErrChangeRole
	}
	existingUser, err := s.Repo.GetUserIdByEmail(user.Email)
	if err != nil {
		return err
	}

	user.Id = existingUser.Id
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}

	if user.SecondName != "" {
		existingUser.SecondName = user.SecondName
	}

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
