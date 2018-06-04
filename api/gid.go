package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetGidRecord(record *v1.ResourceRecord) dns.RR {
	gid := record.GetGid()
	return &dns.GID{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeGID,
			Class:  dns.ClassINET,
			Ttl:    0},
		Gid: gid.Gid,
	}
}
