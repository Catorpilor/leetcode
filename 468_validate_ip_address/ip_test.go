package ip

import "testing"

func TestValidateIP(t *testing.T) {
	st := []struct {
		name string
		ip   string
		exp  string
	}{
		{"not a valid ipv4", "256.0.0.0", "Neither"},
		{"valid ipv4", "255.255.0.0", "IPV4"},
		{"valid ipv6", "2001:0db8:85a3:0:0:8A2E:0370:7334", "IPV6"},
		{"invalid ipv6", "2001:0db8:85a3::8A2E:0370:7334", "Neither"},
		{"failed1", "2001:0db8:85a3:0:0:8A2E:0370:7334:", "Neither"},
		{"failed2", "20EE:FGb8:85a3:0:0:8A2E:0370:7334", "Neither"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := validateIP(tt.ip)
			if out != tt.exp {
				t.Fatalf("with input ip:%s wanted %s but got %s", tt.ip, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
