package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) UpdateUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		UPDATE users
		SET
			first_name = $2,
			second_name = $3,
			email = $4,
			password = $5,
			role_id = $6
		WHERE id = $1
	`, user.Id, user.FirstName, user.SecondName, user.Email, user.Password, user.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user": user,
			"err":  err,
		}).Error("error in repo, UpdateUser")
	}

	return
}
