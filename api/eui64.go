package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetEui64Record(record *v1.ResourceRecord) dns.RR {
	eui64 := record.GetEui64()
	return &dns.EUI64{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeEUI64,
			Class:  dns.ClassINET,
			Ttl:    0},
		Address: eui64.Address,
	}
}
