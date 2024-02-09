package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) UpdateProject(project models.Project) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
	UPDATE projects
	SET
    	name = $2,
    	description = $3,
    	status_id = $4,
    	manager_id = $5,
    	start_date = $6,
    	dead_line = $7
	WHERE id = $1`, project.Id, project.Name, project.Description, project.Status.Id, project.Manager.Id, project.StartDate, project.DeadLine)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"project": project,
			"err":     err,
		}).Error("error in repo, UpdateProject")
	}

	return
}
