package database

import (
	"SUP/pkg/config"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func NewDatabase(
	config *config.Config,
	logger *logrus.Logger,
) *pgx.Conn {
	con, err := pgx.Connect(context.Background(), config.PostgresURL)

	if err != nil {
		logger.Error(err)
		panic(err)
	}

	return con
}
