package models

import "time"

type User struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       Role      `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserRole struct {
	User     User   `json:"user"`
	RoleName string `json:"role_name"`
}
