package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetHipRecord(record *v1.ResourceRecord) dns.RR {
	hip := record.GetHip()
	return &dns.HIP{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeHIP,
			Class:  dns.ClassINET,
			Ttl:    0},
		Hit:                hip.Hit,
		HitLength:          uint8(hip.HitLength),
		PublicKey:          hip.PublicKey,
		PublicKeyAlgorithm: uint8(hip.PublicKeyAlgorithm),
		PublicKeyLength:    uint16(hip.PublicKeyLength),
		RendezvousServers:  hip.RendezvousServers,
	}
}
