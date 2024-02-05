package models

type ProjectParticipant struct {
	Id            int     `json:"id"`
	ParticipantId User    `json:"participant_id"`
	ProjectId     Project `json:"project_id"`
}
