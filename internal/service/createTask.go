package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) CreateTask(task models.Task, status models.Status) (err error) {
	_, err = s.Repo.CheckTaskByTitle(task)

	if err != errors.ErrDataNotFound {
		if err == nil {

			return errors.ErrTaskAlreadyExists
		}
		return
	}

	if task.Title == "" || task.Description == "" || task.Controller.Email == "" || task.Project.Name == "" || task.Executor.Email == "" || status.Name == "" {
		return errors.ErrDataNotFound
	}

	statusFromDB, err := s.Repo.GetStatusByName(status)
	if err != nil {
		return err
	}

	task.Status.Id = statusFromDB.Id

	controller, err := s.Repo.GetUserIdByEmail(task.Controller.Email)
	if err != nil {
		return err
	} else if controller.Role.Id != 2 {
		return errors.ErrAccessDenied
	}

	task.Controller.Id = controller.Id

	executor, err := s.Repo.GetUserIdByEmail(task.Executor.Email)
	if err != nil {
		return err
	}
	task.Executor.Id = executor.Id

	project, err := s.Repo.CheckProjectByName(task.Project)
	if err != nil {
		return err
	}
	task.Project.Id = project.Id
	err = s.Repo.CreateTask(task)

	return
}
