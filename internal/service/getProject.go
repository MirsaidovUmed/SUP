package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) GetProject(project models.Project) (models.Project, error) {
	projectFromDB, err := s.Repo.GetProject(project)
	if err == errors.ErrDataNotFound {
		return models.Project{}, errors.ErrDataNotFound
	} else if err != nil {
		return models.Project{}, err
	}

	return projectFromDB, nil
}
