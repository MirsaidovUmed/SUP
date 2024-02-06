package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"fmt"
)

func (s *Service) CreateProjectParticipant(projectParticipant models.ProjectParticipant) (err error) {
	projectFromDB, err := s.Repo.CheckProjectByName(projectParticipant.Project)
	fmt.Println(projectFromDB)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrProjectNotFound
		}
		return err
	}

	participant, err := s.Repo.GetUserIdByEmail(projectParticipant.Participant.Email)
	if err != nil {
		fmt.Println(1)
		return
	}
	projectParticipant.Participant.Id = participant.Id

	projectParticipant.Project.Id = projectFromDB.Id

	err = s.Repo.CreateProjectParticipant(projectParticipant)
	return
}
