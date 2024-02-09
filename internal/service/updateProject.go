package service

import (
	"SUP/internal/models"
	"fmt"
)

func (s *Service) UpdateProject(project models.Project) (err error) {
	existingProject, err := s.Repo.GetProject(project)
	if err != nil {
		fmt.Println(2)
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
		fmt.Println(1)
		return err
	}

	existingProject.Status.Id = statusId.Id

	managerId, err := s.Repo.GetUserIdByEmail(project.Manager.Email)
	if err != nil {
		fmt.Println(1)
		return err
	}

	existingProject.Manager.Id = managerId.Id

	err = s.Repo.UpdateProject(existingProject)

	return err
}
