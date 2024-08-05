package simplehosting

import (
	"github.com/go-gandi/go-gandi/internal/client"
)

// SimpleHosting is the API client to the Gandi v5 Simple Hosting API
type SimpleHosting struct {
	client client.Gandi
}

type Datacenter struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

// Database represents the type of a Simple Hosting database
type Database struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Language represents the type of a Simple Hosting database
type Language struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Instance struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Size       string      `json:"size"`
	Status     string      `json:"status"`
	Database   *Database   `json:"database"`
	Language   *Language   `json:"language"`
	Datacenter *Datacenter `json:"datacenter"`
}

type InstanceType struct {
	Database *Database `json:"database"`
	Language *Language `json:"language"`
}

type CreateInstanceRequest struct {
	Location string        `json:"location"`
	Type     *InstanceType `json:"type"`
	Name     string        `json:"name"`
	Size     string        `json:"size"`
}

type ErrorResponse struct {
	Cause   string `json:"cause,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Object  string `json:"object,omitempty"`
}

type LinkedDNSZone struct {
	AllowAlteration   bool   `json:"allow_alteration"`
	LastCheckedStatus string `json:"last_checked_status"`
}

type Vhost struct {
	CreatedAt     string         `json:"created_at"`
	FQDN          string         `json:"fqdn"`
	IsATestVhost  bool           `json:"is_a_test_vhost"`
	LinkedDNSZone *LinkedDNSZone `json:"linked_dns_zone"`
	Status        string         `json:"status"`
	Application   *Application   `json:"application,omitempty"`
}

type LinkedDNSZoneRequest struct {
	AllowAlteration        bool `json:"allow_alteration"`
	AlowAlterationOverride bool `json:"allow_alteration_override,omitempty"`
}

type Application struct {
	Name string `json:"name"`
}

type CreateVhostRequest struct {
	FQDN          string                `json:"fqdn"`
	LinkedDNSZone *LinkedDNSZoneRequest `json:"linked_dns_zone,omitempty"`
	Application   *Application          `json:"application,omitempty"`
}

type PatchVhostRequest struct {
	Application *Application `json:"application,omitempty"`
}

type PatchVhostResponse struct {
	FQDN   string `json:"fqdn"`
	Status string `json:"status"`
}
