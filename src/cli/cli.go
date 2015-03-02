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
        }, {
            Name: "stop",
            Usage: "Stop running containers by sending SIGTERM and then SIGKILL after a grace period",
            Action: multiDockerCommand.StopContainers,
            Flags: []cli.Flag{
                cli.IntFlag{
                    Name: "time, t",
                    Value: 10,
                    Usage : "Number of seconds to wait for the container to stop before killing it. Default is 10 seconds.",
                }, cli.StringSliceFlag{
                    Name: "node",
                    Value: &cli.StringSlice{},
                    Usage: "Nodes aliases on which running containers will be stopped ",
                }, cli.StringSliceFlag{
                    Name: "image",
                    Value: &cli.StringSlice{},
                    Usage: "Stop all running container using thoses images",
                },
            },
        },
    }
    return app
}

