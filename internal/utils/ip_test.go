package utils

import "testing"

func TestIsIPv4(t *testing.T) {
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", true}, {"IPV4", true}, {"ipv4", true}, {"IpV4", true},
		{"IPV6", false}, {"ipv6", false}, {"bogus", false},
	} {
		if got := IsIPv4(c.in); got != c.want {
			t.Errorf("IsIPv4(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {"IPV4", false}, {"IPV6", true}, {"ipv6", true}, {"bogus", false},
	} {
		if got := IsIPv6(c.in); got != c.want {
			t.Errorf("IsIPv6(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestRecordType(t *testing.T) {
	for _, c := range []struct {
		in   string
		want string
	}{
		{"", IPTypeA}, {"ipv4", IPTypeA}, {"IPV6", IPTypeAAAA}, {"ipv6", IPTypeAAAA}, {"bogus", IPTypeA},
	} {
		if got := RecordType(c.in); got != c.want {
			t.Errorf("RecordType(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
