package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) Registration(user models.User, roleName string) (err error) {
	_, err = s.Repo.GetUserByEmail(user)

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

	err = s.Repo.CreateUser(user)

	return
}
