package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNsec3Record(record *v1.ResourceRecord) dns.RR {
	nsec3 := record.GetNsec3()
	typeBitMap := make([]uint16, len(nsec3.TypeBitMap))
	for i, val := range nsec3.TypeBitMap {
		typeBitMap[i] = uint16(val)
	}
	return &dns.NSEC3{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNSEC3,
			Class:  dns.ClassINET,
			Ttl:    0},
		Flags:      uint8(nsec3.Flags),
		Hash:       uint8(nsec3.Hash),
		HashLength: uint8(nsec3.HashLength),
		Iterations: uint16(nsec3.Iterations),
		NextDomain: nsec3.NextDomain,
		Salt:       nsec3.Salt,
		SaltLength: uint8(nsec3.SaltLength),
		TypeBitMap: typeBitMap,
	}
}
