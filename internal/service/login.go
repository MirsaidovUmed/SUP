package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/utils"
	"fmt"
)

func (s *Service) Login(user models.User) (accessToken string, err error) {
	userFromDB, err := s.Repo.GetUserByEmail(user)
	if err != nil {
		return
	}

	fmt.Println(userFromDB.Password)

	if userFromDB.Password != user.Password {
		err = errors.ErrWrongPassword
		return
	}

	accessToken, err = utils.CreateToken(s.Config.JwtSecretKey, userFromDB.Id)

	return
}
