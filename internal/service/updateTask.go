package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) UpdateTask(task models.Task) (err error) {
	existingTask, err := s.Repo.GetTask(task)
	if err != nil {
		return err
	}

	task.Id = existingTask.Id
	if task.Title != "" {
		existingTask.Title = task.Title
	}

	if task.Description != "" {
		existingTask.Description = task.Description
	}

	statusId, err := s.Repo.GetStatusByName(task.Status)
	if err != nil {
		return err
	}

	existingTask.Status.Id = statusId.Id

	controllerId, err := s.Repo.GetUserIdByEmail(task.Controller.Email)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return
	} else if controllerId.Id != 2 {
		return errors.ErrAccessDenied
	}

	existingTask.Controller.Id = controllerId.Id

	executorId, err := s.Repo.GetUserIdByEmail(task.Executor.Email)
	if err == errors.ErrDataNotFound {
		return errors.ErrUserNotFound
	}

	existingTask.Executor.Id = executorId.Id

	projectId, err := s.Repo.CheckProjectByName(task.Project)
	if err == errors.ErrDataNotFound {
		return errors.ErrProjectNotFound
	}

	existingTask.Project.Id = projectId.Id

	err = s.Repo.UpdateTask(existingTask)

	return
}
