package repository

import (
	"SUP/pkg/errors"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) CheckUserById(userID int) (err error) {
	var idFromDB int

	row := repo.Conn.QueryRow(context.Background(), `
			select id from users 
			where id = $1`, userID)

	err = row.Scan(&idFromDB)

	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
