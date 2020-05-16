package testutils

import (
	"io/ioutil"
	"os"

	"log"

	"github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/internal/config"
)

func init() {
	config.FAST_POLL = true
	if os.Getenv("DEBUG") == "" {
		logrus.SetOutput(ioutil.Discard)
		log.SetOutput(ioutil.Discard)
	} else {
		if level, err := logrus.ParseLevel(os.Getenv("DEBUG")); err == nil {
			logrus.SetLevel(level)
		}
	}
}
