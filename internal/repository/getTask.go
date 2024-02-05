package repository

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetTask(task models.Task) (taskFromBD models.Task, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
			select * from tasks 
			where title = $1`, task.Title)

	err = row.Scan(&taskFromBD.Id, &taskFromBD.Title, &taskFromBD.Description, &taskFromBD.Status.Id, &taskFromBD.Controller.Id, &taskFromBD.Executor.Id, &taskFromBD.Project.Id, &taskFromBD.StartDate, &taskFromBD.DeadLine)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
