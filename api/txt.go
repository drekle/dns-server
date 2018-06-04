package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetTxtRecord(record *v1.ResourceRecord) dns.RR {
	txt := record.GetTxt()
	return &dns.TXT{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeTXT,
			Class:  dns.ClassINET,
			Ttl:    0},
		Txt: txt.Txt,
	}
}
