package api

import (
	"net"

	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetARecord(record *v1.ResourceRecord) dns.RR {
	A := record.GetA()
	return &dns.A{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    0},
		A: net.ParseIP(A.Address).To4(),
	}
}
