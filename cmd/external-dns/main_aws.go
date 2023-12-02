//go:build all || aws
// +build all aws

package main

import (
	"github.com/aws/aws-sdk-go/service/route53"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/aws"
)

func init() {
	if cfg.Provider == "aws" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		zoneTypeFilter := provider.NewZoneTypeFilter(cfg.AWSZoneType)
		zoneTagFilter := provider.NewZoneTagFilter(cfg.AWSZoneTagFilter)
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
		p, err := aws.NewAWSProvider(
			aws.AWSConfig{
				DomainFilter:         domainFilter,
				ZoneIDFilter:         zoneIDFilter,
				ZoneTypeFilter:       zoneTypeFilter,
				ZoneTagFilter:        zoneTagFilter,
				BatchChangeSize:      cfg.AWSBatchChangeSize,
				BatchChangeInterval:  cfg.AWSBatchChangeInterval,
				EvaluateTargetHealth: cfg.AWSEvaluateTargetHealth,
				PreferCNAME:          cfg.AWSPreferCNAME,
				DryRun:               cfg.DryRun,
				ZoneCacheDuration:    cfg.AWSZoneCacheDuration,
			},
			route53.New(awsSession))
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
