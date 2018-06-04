package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetSshfpRecord(record *v1.ResourceRecord) dns.RR {
	sshfp := record.GetSshfp()
	return &dns.SSHFP{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeSSHFP,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:   uint8(sshfp.Algorithm),
		FingerPrint: sshfp.Fingerprint,
		Type:        uint8(sshfp.Type),
	}
}
