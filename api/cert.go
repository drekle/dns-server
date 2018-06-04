package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetCertRecord(record *v1.ResourceRecord) dns.RR {
	cert := record.GetCert()
	return &dns.CERT{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeCERT,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:   uint8(cert.Algorithm),
		Certificate: cert.Certificate,
		KeyTag:      uint16(cert.KeyTag),
		Type:        uint16(cert.Type),
	}
}
