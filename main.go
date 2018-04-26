package main

import (
	"github.com/urfave/cli"
	"fmt"
	"time"
	"os"
	"errors"
	"github.com/cb-migrate/models"
	"github.com/cb-migrate/connection"
)

func main() {
	source, target := models.Cluster{}, models.Cluster{}

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
			Destination: &source.Address,
		},
		cli.StringFlag {
			Name: "source_user",
			Value: "Administrator",
			Usage: "Source Cluster Username",
			Destination: &source.Username,
		},
		cli.StringFlag {
			Name: "source_pass",
			Value: "password",
			Usage: "Source Cluster Password",
			Destination: &source.Password,
		},

		// target
		cli.StringFlag{
			Name: "target",
			Value: "localhost:8091",
			Usage: "The location where to post the buckets from source",
			Destination: &target.Address,
		},
		cli.StringFlag {
			Name: "target_user",
			Value: "Administrator",
			Usage: "Destination Cluster Username",
			Destination: &target.Username,
		},
		cli.StringFlag {
			Name: "target_pass",
			Value: "password",
			Usage: "Destination Cluster Password",
			Destination: &target.Password,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run the cli app",
			Action: func(c *cli.Context) {
				err := checkFlags(&source, &target)
				if err != nil {
					panic(err.Error())
				}

				err = checkIfUrl(source.Address, target.Address)
				if err != nil {
					panic(err.Error())
				}

				sCtrl, err := connection.GetBucketServer(&source)
				if err != nil {} // todo: error handle
				tCtrl, err := connection.GetBucketServer(&target)
				if err != nil {} // todo: error handle

			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func checkFlags(source, target *models.Cluster) (error) {
	if source.Address == "" || source.Username == "" || source.Password == "" ||
		target.Address == "" || target.Username == "" || target.Password == "" {
			return errors.New("check flags")
	}
	return nil
}

func checkIfUrl(source, target string) (error) {
	return nil
}