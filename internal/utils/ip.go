package utils

import (
	"net"
	"strings"

	dnsResolver "github.com/TimothyYe/godns/pkg/resolver"
	"github.com/miekg/dns"
)

// IsIPv4 reports whether the configured ip_type selects IPv4.
// An empty ip_type defaults to IPv4, matching how IPHelper decides
// which IP to fetch.
func IsIPv4(ipType string) bool {
	return ipType == "" || strings.ToUpper(ipType) == IPV4
}

// IsIPv6 reports whether the configured ip_type selects IPv6.
func IsIPv6(ipType string) bool {
	return strings.ToUpper(ipType) == IPV6
}

// RecordType returns the DNS record type ("A" or "AAAA") for the
// configured ip_type. Anything that is not IPv6 maps to "A".
func RecordType(ipType string) string {
	if IsIPv6(ipType) {
		return IPTypeAAAA
	}
	return IPTypeA
}

// ResolveDNS will query DNS for a given hostname.
func ResolveDNS(hostname, resolver, ipType string) (string, error) {
	dnsType := dns.TypeA
	if IsIPv6(ipType) {
		dnsType = dns.TypeAAAA
	}

	// If no DNS server is set in config file, falls back to default resolver.
	if resolver == "" {
		dnsAddress, err := net.LookupHost(hostname)
		if err != nil {
			return "<nil>", err
		}

		return dnsAddress[0], nil
	}
	res := dnsResolver.New([]string{resolver})
	// In case of i/o timeout
	res.RetryTimes = 5

	ip, err := res.LookupHost(hostname, dnsType)
	if err != nil {
		return "<nil>", err
	}

	return ip[0].String(), nil
}
