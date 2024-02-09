package service

import (
	"SUP/internal/models"
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
		return err
	}

	existingTask.Controller.Id = controllerId.Id

	executorId, err := s.Repo.GetUserIdByEmail(task.Executor.Email)
	if err != nil {
		return err
	}

	existingTask.Executor.Id = executorId.Id

	projectId, err := s.Repo.CheckProjectByName(task.Project)
	if err != nil {
		return err
	}

	existingTask.Project.Id = projectId.Id

	err = s.Repo.UpdateTask(existingTask)

	return err
}
