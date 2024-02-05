package service

import "SUP/internal/models"

func (s *Service) StatusExists(status models.Status) (err error) {
	_, err = s.Repo.GetStatusByName(status)
	if err != nil {
		return err
	}
	return
}
