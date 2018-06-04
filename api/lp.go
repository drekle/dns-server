package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetLpRecord(record *v1.ResourceRecord) dns.RR {
	lp := record.GetLp()
	return &dns.LP{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeLP,
			Class:  dns.ClassINET,
			Ttl:    0},
		Fqdn:       lp.FQDN,
		Preference: uint16(lp.Preference),
	}
}
