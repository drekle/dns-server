package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetDnskeyRecord(record *v1.ResourceRecord) dns.RR {
	dnskey := record.GetDnskey()
	return &dns.DNSKEY{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeDNSKEY,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm: uint8(dnskey.Algorithm),
		Flags:     uint16(dnskey.Flags),
		Protocol:  uint8(dnskey.Protocol),
		PublicKey: dnskey.PublicKey,
	}
}
