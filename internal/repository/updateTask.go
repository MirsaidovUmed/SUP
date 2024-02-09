package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) UpdateTask(task models.Task) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
	UPDATE tasks
	SET
    	title = $2,
    	description = $3,
    	status_id = $4,
    	controller_id = $5,
    	executor_id = $6,
    	project_id = $7,
    	start_date = $8,
    	dead_line = $9
	WHERE id = $1`, task.Id, task.Title, task.Description, task.Status.Id, task.Controller.Id, task.Executor.Id, task.Project.Id, task.StartDate, task.DeadLine)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"task": task,
			"err":  err,
		}).Error("error in repo, UpdateTask")
	}

	return
}
