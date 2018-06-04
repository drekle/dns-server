package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNinfoRecord(record *v1.ResourceRecord) dns.RR {
	ninfo := record.GetNinfo()
	return &dns.NINFO{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNINFO,
			Class:  dns.ClassINET,
			Ttl:    0},
		ZSData: ninfo.ZSDATA,
	}
}
