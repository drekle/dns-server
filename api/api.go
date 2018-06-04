package api

import (
	"context"
	"log"
	"net"

	"github.com/drekle/dns-server/pkg/lib/go/v1"
	impl "github.com/drekle/dns-server/store"
	"github.com/drekle/dns-server/store/database"
	"github.com/miekg/dns"
)

type addrHandle struct {
	Addr   []string
	Handle func(w dns.ResponseWriter, req *dns.Msg)
	rr     *[]dns.RR
}

type Server struct {
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

func (s *Server) ServeDNS(w dns.ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)

	rrList, err := s.store.ReadRecord(context.Background(), req.Question[0].Name)
	if err != nil {
		log.Println("There was an error reading the resource records during a DNS lookup: ", err)
	}
	m.Answer = make([]dns.RR, len(rrList))
	for index, rr := range rrList {
		m.Answer[index] = rr
	}
	w.WriteMsg(m)
}

func NewServer(dbaddr string) *Server {
	var s Server
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

func (s *Server) PostRecord(ctx context.Context, record *v1.ResourceRecord) (*v1.ResourceRecord, error) {
	rr := getRRfromAPIRecord(record)
	s.store.UpdateRecord(ctx, rr)
	return record, nil
}

func (s *Server) GetRecord(ctx context.Context, header *v1.Header) (*v1.ResourceRecords, error) {
	var ret v1.ResourceRecords

	rrList, err := s.store.ReadRecord(ctx, header.Name)
	if err != nil {
		log.Println(err)
	}
	_ = rrList

	for _, rr := range rrList {
		//TODO: Reflection
		if aRecord, ok := (rr).(*dns.A); ok {
			var rec v1.ResourceRecord
			rec.Header = &v1.Header{
				Name: aRecord.Header().Name,
			}
			rec.Record = &v1.ResourceRecord_A{
				A: &v1.A{
					Address: aRecord.A.To4().String(),
				},
			}

			ret.Records = append(ret.Records, &rec)
		} else if aRecord, ok := (rr).(*dns.AAAA); ok {
			var rec v1.ResourceRecord
			rec.Header = &v1.Header{
				Name: aRecord.Header().Name,
			}
			rec.Record = &v1.ResourceRecord_Aaaa{
				Aaaa: &v1.AAAA{
					Address: aRecord.AAAA.To16().String(),
				},
			}

			ret.Records = append(ret.Records, &rec)

		}
	}
	return &ret, nil
}

//Replace
func (s *Server) PutRecord(ctx context.Context, record *v1.ResourceRecord) (*v1.ResourceRecord, error) {

	rr := getRRfromAPIRecord(record)
	s.store.CreateRecord(ctx, rr)
	return record, nil
}

func (s *Server) DeleteRecord(ctx context.Context, record *v1.ResourceRecord) (*v1.ResourceRecord, error) {
	// q, err := s.store.ReadRecord(ctx, entry.Name)
	// if err != nil {
	// 	log.Println(err)
	// 	return &v1.Entry{}, nil
	// }
	//TODO:
	// for _, a := range q.RR {
	// 	// A record
	// 	if entry.Address == a.A.To4().String() {
	// 		_, err = s.store.DeleteRecord(ctx, a)
	// 		if err != nil {
	// 			log.Println("Failed to delete record: ", err)
	// 		}
	// 	}

	// }
	return record, nil
}
