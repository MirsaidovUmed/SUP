package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateTask(task models.Task) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO tasks(title, description, status_id, controller_id, executor_id, project_id, start_date, dead_line
		)VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	`, task.Title, task.Description, task.Status.Id, task.Controller.Id, task.Executor.Id, task.Project.Id, task.StartDate, task.DeadLine)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"tasks": task,
			"err":   err,
		}).Error("error in repo, CreateTask")
	}

	return
}
