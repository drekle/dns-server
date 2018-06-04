package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNsec3ParamRecord(record *v1.ResourceRecord) dns.RR {
	nsec3param := record.GetNsec3Param()
	return &dns.NSEC3PARAM{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNSEC3PARAM,
			Class:  dns.ClassINET,
			Ttl:    0},
		Flags:      uint8(nsec3param.Flags),
		Hash:       uint8(nsec3param.Hash),
		Iterations: uint16(nsec3param.Iterations),
		Salt:       nsec3param.Salt,
		SaltLength: uint8(nsec3param.SaltLength),
	}
}
