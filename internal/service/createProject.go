package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) CreateProject(project models.Project, status models.Status, managerEmail string) (err error) {
	_, err = s.Repo.CheckProjectByName(project)

	if err != errors.ErrDataNotFound {
		if err == nil {
			return errors.ErrProjectAlreadyExists
		}
		return
	}

	statusFromDB, err := s.Repo.GetStatusByName(status)
	if err != nil {
		return
	}
	project.Status.Id = statusFromDB.Id

	manager, err := s.Repo.GetUserIdByEmail(managerEmail)
	if err != nil {
		return
	} else if manager.Id != 2 {
		return errors.ErrAccessDenied
	}
	project.Manager.Id = manager.Id

	err = s.Repo.CreateProject(project)

	return
}
