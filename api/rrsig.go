package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetRrsigRecord(record *v1.ResourceRecord) dns.RR {
	rrsig := record.GetRrsig()
	return &dns.RRSIG{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeRRSIG,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:   uint8(rrsig.Algorithm),
		Expiration:  rrsig.Expiration,
		Inception:   rrsig.Inception,
		KeyTag:      uint16(rrsig.Keytag),
		Labels:      uint8(rrsig.Labels),
		OrigTtl:     rrsig.Origttl,
		Signature:   rrsig.Signature,
		SignerName:  rrsig.SignerName,
		TypeCovered: uint16(rrsig.TypeCovered),
	}
}
