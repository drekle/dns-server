package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNsRecord(record *v1.ResourceRecord) dns.RR {
	ns := record.GetNs()
	return &dns.NS{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNS,
			Class:  dns.ClassINET,
			Ttl:    0},
		Ns: ns.Ns,
	}
}
