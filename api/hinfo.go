package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetHinfoRecord(record *v1.ResourceRecord) dns.RR {
	hinfo := record.GetHinfo()
	return &dns.HINFO{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeHINFO,
			Class:  dns.ClassINET,
			Ttl:    0},
		Cpu: hinfo.Cpu,
		Os:  hinfo.Os,
	}
}
