package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetX25Record(record *v1.ResourceRecord) dns.RR {
	x25 := record.GetX25()
	return &dns.X25{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeX25,
			Class:  dns.ClassINET,
			Ttl:    0},
		PSDNAddress: x25.PSDNAddress,
	}
}
