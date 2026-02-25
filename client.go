package main

import (
	"fmt"
	"os"
)

type ClientCmd struct {
	GetCertificate *GetCertificateCmd `arg:"subcommand:get-certificate"`
}

func doClientCmd() {
	if args.Client.GetCertificate != nil {
		doGetCertificateCmd()
	} else {
		fmt.Fprintln(os.Stderr, specifySubcommand)
		os.Exit(1)
	}
}
