package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/rodrigofrumento/go-blog/internal/logging"
)

func usage() {
	fmt.Println(`This program runs Go-Blog backend server.
	
	Usage:
	blog [arguments]

	Supported arguments:
	`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetGinLogToFile()
	}
	return *env
}
