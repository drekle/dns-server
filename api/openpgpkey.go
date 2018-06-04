package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetOpenpgpkeyRecord(record *v1.ResourceRecord) dns.RR {
	openpgpkey := record.GetOpenpgpkey()
	return &dns.OPENPGPKEY{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeOPENPGPKEY,
			Class:  dns.ClassINET,
			Ttl:    0},
		PublicKey: openpgpkey.PublicKey,
	}
}
