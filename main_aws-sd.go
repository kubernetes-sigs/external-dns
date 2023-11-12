//go:build all || awssd
// +build all awssd

package main

import (
	sd "github.com/aws/aws-sdk-go/service/servicediscovery"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/aws"
	"sigs.k8s.io/external-dns/provider/awssd"
)

func init() {
	var err error
	awsSession, err = aws.NewSession(
		aws.AWSSessionConfig{
			AssumeRole:           cfg.AWSAssumeRole,
			AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
			APIRetries:           cfg.AWSAPIRetries,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	// Check that only compatible Registry is used with AWS-SD
	if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
		log.Infof("Registry \"%s\" cannot be used with AWS Cloud Map. Switching to \"aws-sd\".", cfg.Registry)
		cfg.Registry = "aws-sd"
	}
	p, err := awssd.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.DryRun, cfg.AWSSDServiceCleanup, cfg.TXTOwnerID, sd.New(awsSession))
	if err != nil {
		log.Fatal(err)
	}
	providerMap["aws-sd"] = p
}
