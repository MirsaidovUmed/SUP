package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) GetTask(task models.Task) (taskFromDB models.Task, err error) {
	taskFromDB, err = s.Repo.GetTask(task)
	if err != errors.ErrDataNotFound {
		return
	}
	return taskFromDB, nil
}
