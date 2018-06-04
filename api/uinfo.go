package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetUinfoRecord(record *v1.ResourceRecord) dns.RR {
	uinfo := record.GetUinfo()
	return &dns.UINFO{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeUINFO,
			Class:  dns.ClassINET,
			Ttl:    0},
		Uinfo: uinfo.Uinfo,
	}
}
