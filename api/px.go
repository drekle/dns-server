package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetPxRecord(record *v1.ResourceRecord) dns.RR {
	px := record.GetPx()
	return &dns.PX{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypePX,
			Class:  dns.ClassINET,
			Ttl:    0},
		Map822:     px.Map822,
		Mapx400:    px.Mapx400,
		Preference: uint16(px.Preference),
	}
}
