package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetEui48Record(record *v1.ResourceRecord) dns.RR {
	eui48 := record.GetEui48()
	return &dns.EUI48{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeEUI48,
			Class:  dns.ClassINET,
			Ttl:    0},
		Address: eui48.Address,
	}
}
