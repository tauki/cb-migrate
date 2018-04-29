package main

import (
	"fmt"
	"github.com/cb-migrate/connection"
	"github.com/cb-migrate/models"
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
			Value:       "localhost",
			Usage:       "The source from where to get the buckets",
			Destination: &source.DBHost,
		},
		cli.StringFlag{
			Name:        "source_user",
			Value:       "Administrator",
			Usage:       "Source Cluster Username",
			Destination: &source.DBUser,
		},
		cli.StringFlag{
			Name:        "source_pass",
			Value:       "password",
			Usage:       "Source Cluster Password",
			Destination: &source.DBPassword,
		},
		cli.StringFlag{
			Name:        "source_port",
			Value:       "8091",
			Usage:       "Port the source is running",
			Destination: &source.DBPort,
		},

		// target
		cli.StringFlag{
			Name:        "target",
			Value:       "localhost",
			Usage:       "The location where to post the buckets from source",
			Destination: &target.DBHost,
		},
		cli.StringFlag{
			Name:        "target_user",
			Value:       "Administrator",
			Usage:       "Destination Cluster Username",
			Destination: &target.DBUser,
		},
		cli.StringFlag{
			Name:        "target_pass",
			Value:       "password",
			Usage:       "Destination Cluster Password",
			Destination: &target.DBPassword,
		},
		cli.StringFlag{
			Name:        "target_port",
			Value:       "8091",
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
					if checkFlags(&source, "Source") {
						break
					}
				}

				for {
					if checkFlags(&target, "target") {
						break
					}
				}

				sCtrl, err := connection.GetClusterServer(&source)
				if err != nil {
				} // todo: error handle

				tCtrl, err := connection.GetClusterServer(&target)
				if err != nil {
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

func checkFlags(data *models.Cluster, context string) (bool) {
	check := true

	if data.DBHost == "" {
		check = false
		fmt.Printf("Please input %s address: ", context)
		fmt.Scanf("%s", data.DBHost)
		for {
			if checkIfUrl(data.DBHost) {
				break
			}
		}
	}

	if data.DBPort == "" {
		check = false
		fmt.Printf("Please input %s port: ", context)
		fmt.Scanf("%s", data.DBPort)
	}

	if data.DBUser == "" {
		check = false
		fmt.Print("Please input DB Username: ")
		fmt.Scanf("%s", data.DBUser)
	}

	if data.DBPassword == "" {
		check = false
		fmt.Print("Please input DB password: ")
		fmt.Scanf("%s", data.DBPassword)
	}

	return check
}

func checkIfUrl(source string) (bool) {
	return true
}
