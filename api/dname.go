package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetDnameRecord(record *v1.ResourceRecord) dns.RR {
	dname := record.GetDname()
	return &dns.DNAME{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeDNAME,
			Class:  dns.ClassINET,
			Ttl:    0},
		Target: dname.Target,
	}
}
