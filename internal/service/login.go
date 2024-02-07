package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(user models.User) (accessToken string, err error) {
	userFromDB, err := s.Repo.GetUserIdByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.ErrWrongPassword
		return "", err
	} else if err != nil {
		return "", err
	}

	accessToken, err = utils.CreateToken(s.Config.JwtSecretKey, userFromDB.Id)
	return accessToken, err
}
