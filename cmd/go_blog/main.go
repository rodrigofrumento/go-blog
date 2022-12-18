package main

import (
	"github.com/rodrigofrumento/go-blog/internal/conf"
	"github.com/rodrigofrumento/go-blog/internal/server"
)

func main() {

	server.Start(conf.NewConfig())
}
