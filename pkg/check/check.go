package check

import (
	"net"
)

func IsIp(ip string) bool {
	// Check is ipv4/ipv6
	err := net.ParseIP(ip)
	if err == nil {
		return false
	}
	return true
}

func IsDomain(domain string) bool {
	// Check is a legal domain using regex

	return true
}
