package repository

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) CheckProjectByName(project models.Project) (projectFromDB models.Project, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
			select id from projects 
			where name = $1`, project.Name)

	err = row.Scan(&projectFromDB.Id)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
