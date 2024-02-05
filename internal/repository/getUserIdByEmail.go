package repository

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserIdByEmail(userEmail string) (userFromDB models.User, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
			SELECT id, first_name, second_name, email, password, role_id
			FROM users 
			WHERE email = $1`, userEmail)

	err = row.Scan(&userFromDB.Id, &userFromDB.FirstName, &userFromDB.SecondName, &userFromDB.Email, &userFromDB.Password, &userFromDB.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user": userFromDB,
			"err":  err,
		}).Error("error in repo, GetUserIdByEmail")
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}
	return
}
