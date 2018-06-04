package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetSmimeaRecord(record *v1.ResourceRecord) dns.RR {
	smimea := record.GetSmimea()
	return &dns.SMIMEA{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeSMIMEA,
			Class:  dns.ClassINET,
			Ttl:    0},
		Certificate:  smimea.Certificate,
		MatchingType: uint8(smimea.MatchingType),
		Selector:     uint8(smimea.Selector),
		Usage:        uint8(smimea.Usage),
	}
}
