package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetCnameRecord(record *v1.ResourceRecord) dns.RR {
	cname := record.GetCname()
	return &dns.CNAME{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeCNAME,
			Class:  dns.ClassINET,
			Ttl:    0},
		Target: cname.Target,
	}
}
