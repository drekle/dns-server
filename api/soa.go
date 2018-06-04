package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetSoaRecord(record *v1.ResourceRecord) dns.RR {
	soa := record.GetSoa()
	return &dns.SOA{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeSOA,
			Class:  dns.ClassINET,
			Ttl:    0},
		Expire:  soa.Expire,
		Mbox:    soa.Mbox,
		Minttl:  soa.Minttl,
		Ns:      soa.Ns,
		Refresh: soa.Refresh,
		Retry:   soa.Retry,
		Serial:  soa.Serial,
	}
}
