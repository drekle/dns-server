package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetAfsdbRecord(record *v1.ResourceRecord) dns.RR {
	Afsdb := record.GetAfsdb()
	return &dns.AFSDB{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeAAAA,
			Class:  dns.ClassINET,
			Ttl:    0},
		Hostname: Afsdb.Hostname,
		Subtype:  (uint16)(Afsdb.Subtype),
	}
}
