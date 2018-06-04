package api

import (
	"net"

	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetL32Record(record *v1.ResourceRecord) dns.RR {
	l32 := record.GetL32()
	return &dns.L32{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeL32,
			Class:  dns.ClassINET,
			Ttl:    0},
		Locator32:  net.ParseIP(l32.Locator32).To4(),
		Preference: uint16(l32.Preference),
	}
}
