package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMinfoRecord(record *v1.ResourceRecord) dns.RR {
	minfo := record.GetMinfo()
	return &dns.MINFO{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMINFO,
			Class:  dns.ClassINET,
			Ttl:    0},
		Email: minfo.Email,
		Rmail: minfo.Rmail,
	}
}
