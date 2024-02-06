package models

type ProjectParticipant struct {
	Id          int     `json:"id"`
	Participant User    `json:"participant_id"`
	Project     Project `json:"project_id"`
}
