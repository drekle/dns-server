package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNsecRecord(record *v1.ResourceRecord) dns.RR {
	nsec := record.GetNsec()
	typeBitMap := make([]uint16, len(nsec.TypeBitMap))
	for i, val := range nsec.TypeBitMap {
		typeBitMap[i] = uint16(val)
	}
	return &dns.NSEC{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNSEC,
			Class:  dns.ClassINET,
			Ttl:    0},
		NextDomain: nsec.NextDomain,
		TypeBitMap: typeBitMap,
	}
}
