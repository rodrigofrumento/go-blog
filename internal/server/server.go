package server

import (
	"github.com/rodrigofrumento/go-blog/internal/database"
	"github.com/rodrigofrumento/go-blog/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()

	router.Run(":8080")
}
