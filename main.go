package main

import (
	"github.com/urfave/cli"
	"fmt"
	"time"
	"os"
)

type cluster struct {
	address string
	username string
	password string
}

func main() {
	source, target := cluster{}, cluster{}

	app := cli.NewApp()
	app.Name = "cb-migrate"
	app.Usage = "Lorem ipsum"
	app.Version = fmt.Sprintf("%s-alpha", time.Now().Local().Format("00.00.00"))

	app.Flags = []cli.Flag {

		// source
		cli.StringFlag {
			Name: "source",
			Value: "localhost:8091",
			Usage: "The source from where to get the buckets",
			Destination: &source.address,
		},
		cli.StringFlag {
			Name: "source_user",
			Value: "Administrator",
			Usage: "Source cluster Username",
			Destination: &source.username,
		},
		cli.StringFlag {
			Name: "source_pass",
			Value: "password",
			Usage: "Source cluster Password",
			Destination: &source.address,
		},

		// target
		cli.StringFlag{
			Name: "target",
			Value: "localhost:8091",
			Usage: "The location where to post the buckets from source",
			Destination: &target.address,
		},
		cli.StringFlag {
			Name: "target_user",
			Value: "Administrator",
			Usage: "Destination cluster Username",
			Destination: &target.username,
		},
		cli.StringFlag {
			Name: "target_pass",
			Value: "password",
			Usage: "Destination cluster Password",
			Destination: &target.address,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run the cli app",
			Action: func(c *cli.Context) {

			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
