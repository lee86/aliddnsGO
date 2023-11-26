package aliddns

import "os"

import log "github.com/sirupsen/logrus"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
