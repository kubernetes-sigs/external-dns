//go:build all || oci
// +build all oci

package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/oci"
)

func init() {
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	var err error
	var config *oci.OCIConfig
	// if the instance-principals flag was set, and a compartment OCID was provided, then ignore the
	// OCI config file, and provide a config that uses instance principal authentication.
	if cfg.OCIAuthInstancePrincipal {
		if len(cfg.OCICompartmentOCID) == 0 {
			err = fmt.Errorf("instance principal authentication requested, but no compartment OCID provided")
		} else {
			authConfig := oci.OCIAuthConfig{UseInstancePrincipal: true}
			config = &oci.OCIConfig{Auth: authConfig, CompartmentID: cfg.OCICompartmentOCID}
		}
	} else {
		config, err = oci.LoadOCIConfig(cfg.OCIConfigFile)
	}

	if err == nil {
		p, err := oci.NewOCIProvider(*config, domainFilter, zoneIDFilter, cfg.OCIZoneScope, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap["oci"] = p
	} else {
		log.Fatal(err)
	}
}
