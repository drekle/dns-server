package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetCdsRecord(record *v1.ResourceRecord) dns.RR {
	cds := record.GetCds()
	return &dns.CDS{
		DS: dns.DS{
			Hdr: dns.RR_Header{
				Name:   record.Header.Name + ".",
				Rrtype: dns.TypeCDS,
				Class:  dns.ClassINET,
				Ttl:    0},
			Algorithm:  uint8(cds.Algorithm),
			Digest:     cds.Digest,
			DigestType: uint8(cds.DigestType),
			KeyTag:     uint16(cds.KeyTag),
		},
	}
}
