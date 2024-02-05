package repository

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) CheckTaskByTitle(task models.Task) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
			select title from tasks 
			where title = $1`, task.Title)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
