package repository

import (
	"SUP/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Conn   *pgx.Conn
	Logger *logrus.Logger
}

type RepositoryInterface interface {
	CreateUser(user models.User) (err error)
	GetUserByEmail(user models.User) (userFromDB models.User, err error)
	CheckUserById(userID int) (err error)
	CreateProject(project models.Project) (err error)
	GetRoleByName(roleName string) (role models.Role, err error)
	GetStatusByName(status models.Status) (statusFromDB models.Status, err error)
	CheckProjectByName(project models.Project) (projectFromDB models.Project, err error)
	RoleExists(roleName string) (err error)
	CreateTask(task models.Task) (err error)
	CheckTaskByTitle(task models.Task) (err error)
	GetUserIdByEmail(userEmail string) (userFromDB models.User, err error)
	GetProject(project models.Project) (projectFromDB models.Project, err error)
	GetTask(task models.Task) (taskFromBD models.Task, err error)
	CreateProjectParticipant(projectParticipant models.ProjectParticipant) (err error)
	UpdateUser(user models.User) (err error)
	UpdateProject(project models.Project) (err error)
	UpdateTask(task models.Task) (err error)
}

func NewRepository(conn *pgx.Conn, logger *logrus.Logger) RepositoryInterface {
	return &Repository{
		Conn:   conn,
		Logger: logger,
	}
}
