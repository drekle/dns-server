package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetOptRecord(record *v1.ResourceRecord) dns.RR {
	opt := record.GetOpt()
	_ = opt
	return &dns.OPT{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeOPT,
			Class:  dns.ClassINET,
			Ttl:    0},
		//TODO: Option: dns.EDNS0
	}
}
