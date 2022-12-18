package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/rodrigofrumento/go-blog/internal/conf"
)

func NewDBOptions(cfg conf.Config) *pg.Options {
	return &pg.Options{
		Addr:     cfg.DbHost + ":" + cfg.Port,
		Database: cfg.DbName,
		User:     cfg.DbUser,
		Password: cfg.DbPassword,
	}
}
