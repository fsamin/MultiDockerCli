package cli

import (
    "github.com/codegangsta/cli"
)

func NewCli() *cli.App {
    app := cli.NewApp()
    app.Name = "multidocker"
    app.Usage = "fight the loneliness!"
    app.Action = func(c *cli.Context) {
        println("HELP !")
    }
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "debug, d",
            Usage: "Debug verbose mode",
        },
    }
    app.Commands = []cli.Command{
        {
            Name:  "images",
            Usage: "List images",
            Action: func(c *cli.Context) {
                println("list images: ", c.Args().First())
            },
        }, {
            Name:   "ps",
            Usage:  "List containers",
            Action: listContainers,
            Flags: []cli.Flag{
                cli.BoolFlag{
                    Name: "all, a",
                    Usage: "List all containers",
                },
                cli.BoolFlag{
                    Name: "size, s",
                    Usage: "Show size",
                },
            },
        },
    }
    return app
}
