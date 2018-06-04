package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetDhcidRecord(record *v1.ResourceRecord) dns.RR {
	dhcid := record.GetDhcid()
	return &dns.DHCID{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeDHCID,
			Class:  dns.ClassINET,
			Ttl:    0},
		Digest: dhcid.Digest,
	}
}
