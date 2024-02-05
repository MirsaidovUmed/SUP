package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) GetProject(project models.Project) (projectFRomDB models.Project, err error) {
	projectFRomDB, err = s.Repo.GetProject(project)
	if err != errors.ErrDataNotFound {
		return
	}
	return projectFRomDB, nil
}
