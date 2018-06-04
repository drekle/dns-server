package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetLocRecord(record *v1.ResourceRecord) dns.RR {
	loc := record.GetLoc()
	return &dns.LOC{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeLOC,
			Class:  dns.ClassINET,
			Ttl:    0},
		Altitude:  loc.Altitude,
		HorizPre:  uint8(loc.HorizPre),
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Size:      uint8(loc.Size),
		Version:   uint8(loc.Version),
		VertPre:   uint8(loc.VertPre),
	}
}
