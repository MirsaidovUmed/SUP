package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status_id"`
	Controller  User      `json:"controller_id"`
	Executor    User      `json:"executor_id"`
	Project     Project   `json:"project_id"`
	StartDate   time.Time `json:"start_date"`
	DeadLine    time.Time `json:"dead_line"`
}

type TaskStatus struct {
	Task   Task   `json:"task"`
	Status Status `json:"status_name"`
}
