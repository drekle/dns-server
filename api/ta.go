package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetTaRecord(record *v1.ResourceRecord) dns.RR {
	ta := record.GetTa()
	return &dns.TA{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeTA,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:  uint8(ta.Algorithm),
		Digest:     ta.Digest,
		DigestType: uint8(ta.DigestType),
		KeyTag:     uint16(ta.KeyTag),
	}
}
