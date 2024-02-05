package repository

import (
	"SUP/internal/models"
	"context"
)

func (repo *Repository) GetRoleByName(roleName string) (role models.Role, err error) {
	err = repo.Conn.QueryRow(context.Background(), `SELECT * FROM roles WHERE name=$1`, roleName).Scan(&role.Id, &role.Name)
	return
}
