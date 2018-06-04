package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetTsigRecord(record *v1.ResourceRecord) dns.RR {
	tsig := record.GetTsig()
	return &dns.TSIG{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeTSIG,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:  tsig.Algorithm,
		Error:      uint16(tsig.Error),
		Fudge:      uint16(tsig.Fudge),
		MAC:        tsig.MAC,
		MACSize:    uint16(tsig.MACSize),
		OrigId:     uint16(tsig.OrigID),
		OtherData:  tsig.OtherData,
		OtherLen:   uint16(tsig.OtherLen),
		TimeSigned: tsig.TimeSigned,
	}
}
