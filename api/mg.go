package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMgRecord(record *v1.ResourceRecord) dns.RR {
	mg := record.GetMg()
	return &dns.MG{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMG,
			Class:  dns.ClassINET,
			Ttl:    0},
		Mg: mg.Mg,
	}
}
