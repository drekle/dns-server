package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetCaaRecord(record *v1.ResourceRecord) dns.RR {
	Caa := record.GetCaa()
	return &dns.CAA{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeCAA,
			Class:  dns.ClassINET,
			Ttl:    0},
		Flag:  uint8(Caa.Flag),
		Tag:   Caa.Tag,
		Value: Caa.Value,
	}
}
