package repository

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) CheckProjectParticipant(projectParticipant models.ProjectParticipant) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		SELECT project_id, user_id 
        FROM project_participant
        WHERE project_id = $1 AND user_id = $2`, projectParticipant.Project.Id, projectParticipant.Participant.Id)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
