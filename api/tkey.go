package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetTkeyRecord(record *v1.ResourceRecord) dns.RR {
	tkey := record.GetTkey()
	return &dns.TKEY{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeTKEY,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm:  tkey.Algorithm,
		Error:      uint16(tkey.Error),
		Expiration: tkey.Expiration,
		Inception:  tkey.Inception,
		Key:        tkey.Key,
		KeySize:    uint16(tkey.KeySize),
		Mode:       uint16(tkey.Mode),
		OtherData:  tkey.OtherData,
		OtherLen:   uint16(tkey.OtherLen),
	}
}
