package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetUidRecord(record *v1.ResourceRecord) dns.RR {
	uid := record.GetUid()
	return &dns.UID{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeUID,
			Class:  dns.ClassINET,
			Ttl:    0},
		Uid: uid.Uid,
	}
}
