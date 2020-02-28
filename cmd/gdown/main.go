package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	// "io/ioutil"
)

func init() {
	// log.SetOutput(ioutil.Discard)
	log.SetLevel(log.DebugLevel)
}

const Version = "0.1.0"

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name: "print-version", Aliases: []string{"V"},
		Usage: "print only the version",
	}

	app := &cli.App{
		Name:    "gdown",
		Version: Version,
		Usage:   "Download files directly from googledrive",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				Usage:   "`\"FILEURL\"` to download from google drive ",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("url") != "" {
				cliDownload(c.String("url"))
			}
			if c.NArg() > 0 {
				link := c.Args().Get(0)
				cliDownload(link)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gdown --help")
}
