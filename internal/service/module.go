package service

import (
	"SUP/internal/models"
	"SUP/internal/repository"
	"SUP/pkg/config"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repo   repository.RepositoryInterface
	Config *config.Config
	Logger *logrus.Logger
}

type ServiceInterface interface {
	Registration(user models.User, roleName string) (err error)
	Login(user models.User) (accessToken string, err error)
	CheckUserById(userID int) (err error)
	CreateProject(project models.Project, status models.Status, managerEmail string) (err error)
	RoleExists(roleName string) (err error)
	StatusExists(status models.Status) (err error)
	CreateTask(task models.Task, status models.Status) (err error)
	GetProject(project models.Project) (projectFromDB models.Project, err error)
	GetTask(task models.Task) (taskFromDB models.Task, err error)
}

func NewService(repo repository.RepositoryInterface, config *config.Config, logger *logrus.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Logger: logger,
	}
}
