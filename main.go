package main

import (
	"fmt"
	cb "github.com/cb-migrate/connection"
	"github.com/cb-migrate/models"
	"github.com/cb-migrate/utility"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	source, target := models.Data{}, models.Data{}

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
			Usage:       "Source Data Username",
			Destination: &source.DBUser,
		},
		cli.StringFlag{
			Name:        "source_pass",
			Usage:       "Source Data Password",
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
			Usage:       "Destination Data Username",
			Destination: &target.DBUser,
		},
		cli.StringFlag{
			Name:        "target_pass",
			Usage:       "Destination Data Password",
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

				sCtrl, err := cb.GetServer(&source)
				if err != nil {
					panic(err)
				} // todo: error handle

				tCtrl, err := cb.GetServer(&target)
				if err != nil {
					panic(err)
				} // todo: error handle

				if len(*sCtrl.Data.Buckets) != len(*tCtrl.Data.Buckets) {
					for _, bucket := range *sCtrl.Data.Buckets {
						if !tCtrl.BucketExists(bucket.Name) {
							tCtrl.CreateBucket(bucket.Name, bucket)
						}
					}
				}

				for _, bucket := range *sCtrl.Data.Buckets {
					err := tCtrl.Copy(bucket)
					if err != nil {
					} // todo: error handle
				}
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
