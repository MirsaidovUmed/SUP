package repository

import (
	"SUP/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetStatusByName(status models.Status) (statusFromDB models.Status, err error) {
	err = repo.Conn.QueryRow(context.Background(), `
		SELECT id, name
		FROM statuses
		WHERE name = $1
	`, status.Name).Scan(&statusFromDB.Id, &statusFromDB.Name)

	fmt.Println("Status Name:", status.Name)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"status_name": status,
			"err":         err,
		}).Error("error in repo, GetStatusByName")
		return status, errors.New("status not found")
	}
	return
}
