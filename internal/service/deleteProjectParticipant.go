package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) DeleteProjectParticipant(projectParticipant models.ProjectParticipant) (err error) {
	projectFromDB, err := s.Repo.CheckProjectByName(projectParticipant.Project)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrProjectNotFound
		}
		return err
	}

	participant, err := s.Repo.GetUserIdByEmail(projectParticipant.Participant.Email)
	if err != nil {
		return
	}
	projectParticipant.Participant.Id = participant.Id

	projectParticipant.Project.Id = projectFromDB.Id
	err = s.Repo.DeleteProjectParticipant(projectParticipant)
	return
}
