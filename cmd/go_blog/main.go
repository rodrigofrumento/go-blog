package main

import (
	"github.com/rodrigofrumento/go-blog/internal/cli"
	"github.com/rodrigofrumento/go-blog/internal/conf"
	"github.com/rodrigofrumento/go-blog/internal/server"
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}
