package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetUriRecord(record *v1.ResourceRecord) dns.RR {
	uri := record.GetUri()
	return &dns.URI{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeURI,
			Class:  dns.ClassINET,
			Ttl:    0},
		Priority: uint16(uri.Priority),
		Target:   uri.Target,
		Weight:   uint16(uri.Weight),
	}
}
