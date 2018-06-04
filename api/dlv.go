package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetDlvRecord(record *v1.ResourceRecord) dns.RR {
	dlv := record.GetDlv()
	return &dns.DLV{
		DS: dns.DS{
			Hdr: dns.RR_Header{
				Name:   record.Header.Name + ".",
				Rrtype: dns.TypeDLV,
				Class:  dns.ClassINET,
				Ttl:    0},
			Algorithm:  uint8(dlv.Algorithm),
			Digest:     dlv.Digest,
			DigestType: uint8(dlv.DigestType),
			KeyTag:     uint16(dlv.KeyTag),
		},
	}
}
