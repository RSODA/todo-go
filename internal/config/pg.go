package config

import (
	"errors"
	"os"
)

const (
	dsnPG = "DSN_PG"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv("PG_CONFIG_DSN")
	if len(dsn) == 0 {
		return nil, errors.New("PG_CONFIG_DSN environment variable not set")
	}

	return &pgConfig{dsn: dsn}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
