package main

import (
	"fmt"
	"github.com/cb-migrate/connection/cluster"
	"github.com/cb-migrate/models"
	"github.com/cb-migrate/utility"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	source, target := models.Cluster{}, models.Cluster{}

	app := cli.NewApp()
	app.Name = "cb-migrate"
	app.Usage = "Lorem ipsum"
	app.Version = fmt.Sprintf("%s-alpha", time.Now().Local().Format("00.00.00"))

	app.Flags = []cli.Flag{

		// source
		cli.StringFlag{
			Name:        "source",
			Usage:       "The source from where to get the buckets",
			Destination: &source.DBHost,
		},
		cli.StringFlag{
			Name:        "source_user",
			Usage:       "Source Cluster Username",
			Destination: &source.DBUser,
		},
		cli.StringFlag{
			Name:        "source_pass",
			Usage:       "Source Cluster Password",
			Destination: &source.DBPassword,
		},
		cli.StringFlag{
			Name:        "source_port",
			Usage:       "Port the source is running",
			Destination: &source.DBPort,
		},

		// target
		cli.StringFlag{
			Name:        "target",
			Usage:       "The location where to post the buckets from source",
			Destination: &target.DBHost,
		},
		cli.StringFlag{
			Name:        "target_user",
			Usage:       "Destination Cluster Username",
			Destination: &target.DBUser,
		},
		cli.StringFlag{
			Name:        "target_pass",
			Usage:       "Destination Cluster Password",
			Destination: &target.DBPassword,
		},
		cli.StringFlag{
			Name:        "target_port",
			Usage:       "Port the target is running",
			Destination: &target.DBPort,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run the cli app",
			Action: func(c *cli.Context) {

				for {
					fmt.Println(source)
					if utility.CheckFlags(&source, "Source") {
						break
					}
				}

				for {
					if utility.CheckFlags(&target, "target") {
						break
					}
				}

				sCtrl, err := cluster.GetServer(&source)
				if err != nil {
					panic(err)
				} // todo: error handle

				tCtrl, err := cluster.GetServer(&target)
				if err != nil {
					panic(err)
				} // todo: error handle

				fmt.Println(sCtrl)
				fmt.Println(tCtrl)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}

