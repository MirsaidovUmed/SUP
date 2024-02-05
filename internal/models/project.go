package models

import "time"

type Project struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      Status    `json:"status_id"`
	Manager     User      `json:"manager_id"`
	StartDate   time.Time `json:"start_date"`
	DeadLine    time.Time `json:"dead_line"`
}

type ProjectStatus struct {
	Project      Project `json:"project"`
	Status       Status  `json:"status_name"`
	ManagerEmail User    `json:"manager_email"`
}
