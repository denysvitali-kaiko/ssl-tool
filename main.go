package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type argsType struct {
	Certificate *CertificateCmd `arg:"subcommand:cert"`
	Client      *ClientCmd      `arg:"subcommand:client"`
	LogLevel    string          `arg:"-l,--log-level" default:"warning" help:"log level (debug, info, warning, error)"`
}

func (argsType) Description() string {
	return "ssl-tool inspects and retrieves SSL/TLS certificates."
}

func (argsType) Version() string {
	return "ssl-tool 1.0.0"
}

var args argsType

const specifySubcommand = "please specify a subcommand"

func main() {
	arg.MustParse(&args)

	switch strings.ToLower(args.LogLevel) {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	}

	if args.Certificate != nil {
		doCertificateCmd(args.Certificate)
	} else if args.Client != nil {
		doClientCmd()
	} else {
		fmt.Fprintln(os.Stderr, specifySubcommand)
		os.Exit(1)
	}
}
