package repository

import (
	"SUP/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) DeleteProjectParticipant(participant models.ProjectParticipant) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		DELETE FROM projects_participant
		WHERE participant_id = $1 AND project_id = $2
	`, participant.Participant.Id, participant.Project.Id)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"projectParticipant": participant,
			"err":                err,
		}).Error("error in repo, DeleteProjectParticipant")
	}
	return
}
