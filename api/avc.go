package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

// Avc

func GetAvcRecord(record *v1.ResourceRecord) dns.RR {
	Avc := record.GetAvc()
	return &dns.AVC{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeAVC,
			Class:  dns.ClassINET,
			Ttl:    0},
		Txt: Avc.Txt,
	}
}
