package database

import "github.com/go-pg/pg/v10"

func NewDBOptions() *pg.Options {
	return &pg.Options{
		Addr:     "localhost:5432",
		Database: "go_blog",
		User:     "postgres",
		Password: "postgres",
	}
}
