package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) CreateProjectParticipant(projectParticipant models.ProjectParticipant) (err error) {

	if projectParticipant.Participant.Email == "" || projectParticipant.Project.Name == "" {
		return errors.ErrDataNotFound
	}

	projectFromDB, err := s.Repo.CheckProjectByName(projectParticipant.Project)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrProjectNotFound
		}
		return err
	}
	projectParticipant.Project.Id = projectFromDB.Id

	err = s.Repo.CheckProjectParticipant(projectParticipant)
	if err != nil {
		return errors.ErrAlreadyHasUser
	}

	participant, err := s.Repo.GetUserIdByEmail(projectParticipant.Participant.Email)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}
	projectParticipant.Participant.Id = participant.Id

	err = s.Repo.CreateProjectParticipant(projectParticipant)
	return
}
