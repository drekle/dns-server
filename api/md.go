package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMdRecord(record *v1.ResourceRecord) dns.RR {
	md := record.GetMd()
	return &dns.MD{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMD,
			Class:  dns.ClassINET,
			Ttl:    0},
		Md: md.Md,
	}
}
