package server

import (
	"github.com/rodrigofrumento/go-blog/internal/conf"
	"github.com/rodrigofrumento/go-blog/internal/database"
	"github.com/rodrigofrumento/go-blog/internal/store"
)

func Start(cfg conf.Config) {
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	router.Run(":8080")
}
