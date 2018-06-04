package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMfRecord(record *v1.ResourceRecord) dns.RR {
	mf := record.GetMf()
	return &dns.MF{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMF,
			Class:  dns.ClassINET,
			Ttl:    0},
		Mf: mf.Mf,
	}
}
