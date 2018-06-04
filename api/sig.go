package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetSigRecord(record *v1.ResourceRecord) dns.RR {
	sig := record.GetSig()
	return &dns.SIG{
		RRSIG: dns.RRSIG{
			Hdr: dns.RR_Header{
				Name:   record.Header.Name + ".",
				Rrtype: dns.TypeSIG,
				Class:  dns.ClassINET,
				Ttl:    0},
			Algorithm:   uint8(sig.Algorithm),
			Expiration:  sig.Expiration,
			Inception:   sig.Inception,
			KeyTag:      uint16(sig.Keytag),
			Labels:      uint8(sig.Labels),
			OrigTtl:     sig.Origttl,
			Signature:   sig.Signature,
			SignerName:  sig.SignerName,
			TypeCovered: uint16(sig.TypeCovered),
		},
	}
}
