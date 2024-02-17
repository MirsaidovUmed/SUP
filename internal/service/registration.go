package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Registration(user models.User, roleName string) (err error) {
	_, err = s.Repo.GetUserIdByEmail(user.Email)

	if err != errors.ErrDataNotFound {
		if err == nil {
			return errors.ErrAlreadyHasUser
		}

		return
	}

	roleId, err := s.Repo.GetRoleByName(roleName)
	if err != nil {
		return
	}

	user.Role = roleId

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = s.Repo.CreateUser(user)

	return
}
