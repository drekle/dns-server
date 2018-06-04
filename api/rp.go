package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetRpRecord(record *v1.ResourceRecord) dns.RR {
	rp := record.GetRp()
	return &dns.RP{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeRP,
			Class:  dns.ClassINET,
			Ttl:    0},
		Mbox: rp.Mbox,
		Txt:  rp.Txt,
	}
}
