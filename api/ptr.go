package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetPtrRecord(record *v1.ResourceRecord) dns.RR {
	ptr := record.GetPtr()
	return &dns.PTR{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypePTR,
			Class:  dns.ClassINET,
			Ttl:    0},
		Ptr: ptr.Ptr,
	}
}
