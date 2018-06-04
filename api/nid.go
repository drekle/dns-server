package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNidRecord(record *v1.ResourceRecord) dns.RR {
	nid := record.GetNid()
	return &dns.NID{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNID,
			Class:  dns.ClassINET,
			Ttl:    0},
		NodeID:     nid.Nodeid,
		Preference: uint16(nid.Preference),
	}
}
