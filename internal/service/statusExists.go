package service

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
)

func (s *Service) StatusExists(status models.Status) (err error) {
	_, err = s.Repo.GetStatusByName(status)
	if err != nil {
		return errors.ErrDataNotFound
	}
	return
}
