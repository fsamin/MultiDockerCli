package cli

import (
    "github.com/codegangsta/cli"
    "log"
    "os")

var multiDockerCommand *DockerCommand

func init() {
    commands, err := NewDockerCommand()
    if err != nil {
        log.Fatal("Cannot initialize CLI", err)
        os.Exit(1)
    }
    multiDockerCommand = commands
}

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
            Action: multiDockerCommand.ListImages,
            Flags: []cli.Flag{
                cli.BoolFlag{
                    Name: "all, a",
                    Usage: "List all images (by default filter out the intermediate image layers)",
                },
                cli.BoolFlag{
                    Name: "size, s",
                    Usage: "Show size",
                },
            },
        }, {
            Name:   "ps",
            Usage:  "List containers",
            Action: multiDockerCommand.ListContainers,
            Flags: []cli.Flag{
                cli.BoolFlag{
                    Name: "all, a",
                    Usage: "List all containers. Only running containers are shown by default.",
                },
                cli.BoolFlag{
                    Name: "size, s",
                    Usage: "Show size",
                },
            },
        }, {
            Name: "pull",
            Usage: "Pull an image or a repository from the registry. Set argument to IMAGE[:TAG]",
            Action: multiDockerCommand.PullImage,
        },
    }
    return app
}

