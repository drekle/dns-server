package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetGposRecord(record *v1.ResourceRecord) dns.RR {
	gpos := record.GetGpos()
	return &dns.GPOS{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeGPOS,
			Class:  dns.ClassINET,
			Ttl:    0},
		Altitude:  gpos.Altitude,
		Latitude:  gpos.Latitude,
		Longitude: gpos.Longitude,
	}
}
