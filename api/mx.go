package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMxRecord(record *v1.ResourceRecord) dns.RR {
	mx := record.GetMx()
	return &dns.MX{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMX,
			Class:  dns.ClassINET,
			Ttl:    0},
		Mx:         mx.Mx,
		Preference: uint16(mx.Preference),
	}
}
