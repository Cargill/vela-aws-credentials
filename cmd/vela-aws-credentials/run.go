package main

import (
	"github.com/Cargill/vela-aws-credentials/pkg/plugin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// run executes the plugin based off the configuration provided.
func run(c *cli.Context) error {
	// create a new, empty sirupsen/logrus logger
	logger := logrus.New()

	// set the log format for the plugin
	switch c.String(plugin.FlagLogFormat) {
	case "text", "Text", "TEXT":
		logger.SetFormatter(&logrus.TextFormatter{})
	case "json", "Json", "JSON":
		fallthrough
	default:
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	// set the log level for the plugin
	switch c.String(plugin.FlagLogLevel) {
	case "t", "trace", "Trace", "TRACE":
		logger.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logger.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logger.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.WithFields(logrus.Fields{
		"code":     "https://github.com/Cargill/vela-aws-credentials",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/pipeline/aws-credentials",
		"registry": "https://hub.docker.com/r/cargill/vela-aws-credentials",
	}).Info("Vela AWS Credentials Config")

	// create the plugin
	p := plugin.FromCLIContext(c, logger.WithField("version", version))

	// validate the plugin
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
