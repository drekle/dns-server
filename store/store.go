package store

import (
	"context"
	"log"

	"encoding/json"

	"github.com/miekg/dns"
)

type Query struct {
	Name string `json:"name"`
	//Cannot use dns.RR here as interfaces cannot serialize
	RR []dns.RR `json:"rr"`
}

//Interface that describes the base implementation to store things in various data sinks
type Store interface {
	CreateRecord(context.Context, dns.RR) (dns.RR, error)
	ReadRecord(context.Context, string) ([]dns.RR, error)
	UpdateRecord(context.Context, dns.RR) (dns.RR, error)
	DeleteRecord(context.Context, dns.RR) (dns.RR, error)
}

func PrintQuery(q *Query) {
	jsonData, err := json.MarshalIndent(q, "", "\t")
	if err == nil {
		log.Println(string(jsonData))
		log.Println()
	}
}
