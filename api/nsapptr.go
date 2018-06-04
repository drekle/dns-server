package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNsapptrRecord(record *v1.ResourceRecord) dns.RR {
	nsapptr := record.GetNsapptr()
	return &dns.NSAPPTR{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNSAPPTR,
			Class:  dns.ClassINET,
			Ttl:    0},
		Ptr: nsapptr.Ptr,
	}
}
