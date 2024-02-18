package repository

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) CheckTaskByTitle(task models.Task) (taskFromDb models.Task, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
			select id from tasks 
			where title = $1`, task.Title)
	err = row.Scan(&taskFromDb.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
