package service

import (
	"SUP/internal/models"
)

func (s *Service) CreateTask(task models.Task, status models.Status) (err error) {
	err = s.Repo.CheckTaskByTitle(task)

	if err != nil {
		return err
	}

	statusFromDB, err := s.Repo.GetStatusByName(status)
	if err != nil {
		return err
	}

	task.Status.Id = statusFromDB.Id

	controller, err := s.Repo.GetUserIdByEmail(task.Controller.Email)
	if err != nil {
		return err
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
