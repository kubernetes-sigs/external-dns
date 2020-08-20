package testutils

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/internal/config"
)

func init() {
	config.FastPoll = true
	if os.Getenv("DEBUG") == "" {
		logrus.SetOutput(ioutil.Discard)
		log.SetOutput(ioutil.Discard)
	} else {
		if level, err := logrus.ParseLevel(os.Getenv("DEBUG")); err == nil {
			logrus.SetLevel(level)
		}
	}
}
