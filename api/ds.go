package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetDsRecord(record *v1.ResourceRecord) dns.RR {
	ds := record.GetDs()
	return &dns.DS{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeDS,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:  uint8(ds.Algorithm),
		Digest:     ds.Digest,
		DigestType: uint8(ds.DigestType),
		KeyTag:     uint16(ds.KeyTag),
	}
}
