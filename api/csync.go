package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetCsyncRecord(record *v1.ResourceRecord) dns.RR {
	csync := record.GetCsync()
	typeBitMap := make([]uint16, len(csync.TypeBitMap))
	for i, val := range csync.TypeBitMap {
		typeBitMap[i] = uint16(val)
	}
	return &dns.CSYNC{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeCSYNC,
			Class:  dns.ClassINET,
			Ttl:    0},
		Flags:      uint16(csync.Flags),
		Serial:     csync.Serial,
		TypeBitMap: typeBitMap,
	}
}
