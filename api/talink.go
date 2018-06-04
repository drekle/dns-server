package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetTalinkRecord(record *v1.ResourceRecord) dns.RR {
	talink := record.GetTalink()
	return &dns.TALINK{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeTALINK,
			Class:  dns.ClassINET,
			Ttl:    0},
		NextName:     talink.NextName,
		PreviousName: talink.PreviousName,
	}
}
