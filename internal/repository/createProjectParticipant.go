package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateProjectParticipant(projectParticipant models.ProjectParticipant) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO projects_participant(
			participant_id, project_id
		) VALUES(
			$1, $2
		)
	`, projectParticipant.Participant.Id, projectParticipant.Project.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"projectParticipant": projectParticipant,
			"err":                err,
		}).Error("error in repo, CreateProjectParticipant")
	}
	return
}
