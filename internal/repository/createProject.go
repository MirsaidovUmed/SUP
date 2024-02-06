package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateProject(project models.Project) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO projects(
			name, description, status_id, manager_id, start_date, dead_line
		) VALUES(
			$1, $2, $3, $4, $5, $6
		)
	`, project.Name, project.Description, project.Status.Id, project.Manager.Id, project.StartDate, project.DeadLine)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"project": project,
			"err":     err,
		}).Error("error in repo, CreateProject")
	}
	return
}
