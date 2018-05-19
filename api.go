package main

import (
	"context"
	"log"
	"net"

	"bitbucket.ema.emoneyadvisor.com/cp/lemondns/lib/go/v1"
	impl "bitbucket.ema.emoneyadvisor.com/cp/lemondns/store"
	"bitbucket.ema.emoneyadvisor.com/cp/lemondns/store/database"
	"github.com/miekg/dns"
)

type addrHandle struct {
	Addr   []string
	Handle func(w dns.ResponseWriter, req *dns.Msg)
	rr     *[]dns.RR
}

type server struct {
	store impl.Store
}

//  *********     IMPORTANT      ************
//  This a one off not intended for long term use
//  I will only support A records here with limited options
//  *****************************************
type Request struct {
	Name    string
	Address []string `json:",omitempty"`
	// TTL  int `json:",omitempty"`
	// Class int `json:",omitempty"`
	// RecordType `json:"type,omitempty"`

}

func (s *server) ServeDNS(w dns.ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)

	q, err := s.store.ReadRecord(context.Background(), req.Question[0].Name)
	if err != nil {
		log.Println("There was an error reading the resource records during a DNS lookup: ", err)
	}
	m.Answer = make([]dns.RR, len(q.RR))
	for index, rr := range q.RR {
		m.Answer[index] = &rr
	}
	w.WriteMsg(m)
}

func NewServer(dbaddr string) *server {
	var s server
	s.store = database.NewMongoDBStore(dbaddr)
	return &s
}

func newResourceRecordForName(name string, address string) *dns.A {
	rr := &dns.A{
		Hdr: dns.RR_Header{
			Name:   name + ".",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    0},
		A: net.ParseIP(address).To4(),
	}
	return rr
}

func (s *server) PostRecord(ctx context.Context, entry *v1.Entry) (*v1.Entry, error) {
	a := newResourceRecordForName(entry.Name, entry.Address)
	s.store.UpdateRecord(ctx, *a)
	return entry, nil
}

func (s *server) GetRecord(ctx context.Context, entry *v1.Entry) (*v1.Entries, error) {
	var ret v1.Entries

	q, err := s.store.ReadRecord(ctx, entry.Name)
	if err == nil {
		for _, a := range q.RR {
			ret.Entries = append(ret.Entries, &v1.Entry{
				Name:    a.Header().Name,
				Address: a.A.To4().String(),
			})
		}
	}
	return &ret, nil
}

//Replace
func (s *server) PutRecord(ctx context.Context, entry *v1.Entry) (*v1.Entry, error) {
	rr := newResourceRecordForName(entry.Name, entry.Address)
	s.store.CreateRecord(ctx, *rr)
	return entry, nil
}

func (s *server) DeleteRecord(ctx context.Context, entry *v1.Entry) (*v1.Entry, error) {
	q, err := s.store.ReadRecord(ctx, entry.Name)
	if err != nil {
		log.Println(err)
		return &v1.Entry{}, nil
	}
	for _, a := range q.RR {
		// A record
		if entry.Address == a.A.To4().String() {
			_, err = s.store.DeleteRecord(ctx, a)
			if err != nil {
				log.Println("Failed to delete record: ", err)
			}
		}

	}
	return &v1.Entry{}, nil
}
