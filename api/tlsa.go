package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetTlsaRecord(record *v1.ResourceRecord) dns.RR {
	tlsa := record.GetTlsa()
	return &dns.TLSA{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeTLSA,
			Class:  dns.ClassINET,
			Ttl:    0},
		Certificate:  tlsa.Certificate,
		MatchingType: uint8(tlsa.MatchingType),
		Selector:     uint8(tlsa.Selector),
		Usage:        uint8(tlsa.Usage),
	}
}
