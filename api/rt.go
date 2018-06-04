package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetRtRecord(record *v1.ResourceRecord) dns.RR {
	rt := record.GetRt()
	return &dns.RT{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeRT,
			Class:  dns.ClassINET,
			Ttl:    0},
		Host:       rt.Host,
		Preference: uint16(rt.Preference),
	}
}
