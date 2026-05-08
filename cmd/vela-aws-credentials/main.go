// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"time"

	"github.com/Cargill/vela-aws-credentials/pkg/plugin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

var (
	version = "dev"
)

func main() {
	// create new CLI application
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-aws-credentials"
	app.HelpName = "vela-aws-credentials"
	app.Usage = "Vela AWS credentials plugin for temporary AWS credentials."

	// Plugin Metadata

	app.Action = run
	app.Compiled = time.Now()
	app.Version = version

	// Plugin Flags

	app.Flags = plugin.Flags

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
