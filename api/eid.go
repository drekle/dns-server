package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetEidRecord(record *v1.ResourceRecord) dns.RR {
	eid := record.GetEid()
	return &dns.EID{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeEID,
			Class:  dns.ClassINET,
			Ttl:    0},
		Endpoint: eid.Endpoint,
	}
}
