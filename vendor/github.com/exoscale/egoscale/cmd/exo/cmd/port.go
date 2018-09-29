package cmd

import (
	"strings"
)

//go:generate stringer -type=port

type port uint16

const (
	// Daytime of the night
	Daytime port = 13
	// FTP protocol port number
	FTP port = 21
	// SSH protocol port number
	SSH port = 22
	// Telnet protocol port number
	Telnet port = 23
	// SMTP protocol port number
	SMTP port = 25
	// Time protocol port number
	Time port = 37
	// Whois protocol port number
	Whois port = 43
	// DNS protocol port number
	DNS port = 53
	// TFTP protocol port number
	TFTP port = 69
	// Gopher protocol port number
	Gopher port = 70
	// HTTP protocol port number
	HTTP port = 80
	// Kerberos protocol port number
	Kerberos port = 88
	// Nic protocol port number
	Nic port = 101
	// SFTP protocol port number
	SFTP port = 115
	// NTP protocol port number
	NTP port = 123
	// IMAP protocol port number
	IMAP port = 143
	// SNMP protocol port number
	SNMP port = 161
	// IRC protocol port number
	IRC port = 194
	// HTTPS protocol port number
	HTTPS port = 443
	// Docker protocol port number
	Docker port = 2376
	// RDP protocol port number
	RDP port = 3389
	// Minecraft protocol port number
	Minecraft port = 25565
)

func (i port) StringFormatted() string {
	res := i.String()
	if strings.HasPrefix(res, "port(") {
		return ""
	}
	return res
}
