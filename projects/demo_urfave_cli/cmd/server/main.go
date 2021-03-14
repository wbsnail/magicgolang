package main

import (
	"app/pkg/log"
	"app/pkg/server"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"os"
)

func main() {
	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "addr",
			Value: "0.0.0.0:8080",
			Usage: "listen address",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "db-driver",
			Usage: "database driver (mysql, sqlite)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "db-args",
			Usage: "database args",
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:  "gin-debug",
			Usage: "enable gin debug mode",
		}),
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to config file (yaml)",
		},
	}
	cliApp := cli.NewApp()
	cliApp.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	cliApp.Flags = flags
	cliApp.Action = func(c *cli.Context) error {
		log.Infof("config file: %s", c.String("config"))

		app, err := server.SetupApp(&server.Options{
			Addr:     c.String("addr"),
			DBDriver: c.String("db-driver"),
			DBArgs:   c.String("db-args"),
			GinDebug: c.Bool("gin-debug"),
		})
		if err != nil {
			return errors.Wrap(err, "setup app error")
		}

		return app.Serve()
	}

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
