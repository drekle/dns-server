package api

import (
	"net"

	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetAaaaRecord(record *v1.ResourceRecord) dns.RR {
	Aaaa := record.GetAaaa()
	return &dns.AAAA{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeAAAA,
			Class:  dns.ClassINET,
			Ttl:    0},
		AAAA: net.ParseIP(Aaaa.Address).To16(),
	}
}
