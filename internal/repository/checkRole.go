package repository

import (
	"SUP/pkg/errors"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) RoleExists(roleName string) (err error) {
	var idFromDB int

	row := repo.Conn.QueryRow(context.Background(), `SELECT id FROM roles where name = $1`, roleName)

	err = row.Scan(&idFromDB)
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
