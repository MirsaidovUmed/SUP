package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) UpdateProject(project models.Project) (err error) {
	existingProject, err := s.Repo.GetProject(project)
	if err != nil {
		return err
	}

	project.Id = existingProject.Id
	if project.Name != "" {
		existingProject.Name = project.Name
	}

	if project.Description != "" {
		existingProject.Description = project.Description
	}

	statusId, err := s.Repo.GetStatusByName(project.Status)
	if err != nil {
		return err
	}

	existingProject.Status.Id = statusId.Id

	managerId, err := s.Repo.GetUserIdByEmail(project.Manager.Email)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
	} else if managerId.Id != 2 {
		return errors.ErrAccessDenied
	}

	existingProject.Manager.Id = managerId.Id

	err = s.Repo.UpdateProject(existingProject)

	return
}
