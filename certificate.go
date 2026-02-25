package main

import (
	"fmt"
	"os"
)

type CertificateCmd struct {
	Parse *ParseCmd `arg:"subcommand:parse"`
}

func doCertificateCmd(cmd *CertificateCmd) {
	if cmd == nil {
		logger.Fatalf("%s", "cmd cannot be nil")
	}

	if cmd.Parse != nil {
		doParseCertificate(cmd.Parse)
		return
	}

	fmt.Fprintln(os.Stderr, specifySubcommand)
	os.Exit(1)
}
