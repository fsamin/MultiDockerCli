package main

import (
	"./cli"
	"os"
)

func main() {
	app := cli.NewCli()
	app.Run(os.Args)
}
