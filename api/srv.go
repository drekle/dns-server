package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetSrvRecord(record *v1.ResourceRecord) dns.RR {
	srv := record.GetSrv()
	return &dns.SRV{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeSRV,
			Class:  dns.ClassINET,
			Ttl:    0},
		Port:     uint16(srv.Port),
		Priority: uint16(srv.Priority),
		Target:   srv.Target,
		Weight:   uint16(srv.Weight),
	}
}
