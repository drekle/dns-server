package database

import (
	"context"
	"log"

	impl "github.com/drekle/dns-server/store"

	"github.com/miekg/dns"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	impl.Store
	address string
}

var tableName = "dns"
var collectionName = "record"

func NewMongoDBStore(address string) *MongoDB {
	var ret MongoDB
	ret.address = address

	session, err := mgo.Dial(ret.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", ret.address)
	}
	defer session.Close()

	return &ret
}

func (store *MongoDB) CreateRecord(ctx context.Context, rr dns.RR) (dns.RR, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	collection := session.DB(tableName).C(collectionName)
	bulk := collection.Bulk()
	bulk.RemoveAll(bson.M{"hdr.name": rr.Header().Name})
	bulk.Insert(&rr)
	_, err = bulk.Run()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rr, nil
}

func (store *MongoDB) ReadRecord(ctx context.Context, name string) ([]dns.RR, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	var qs []dns.RR
	collection := session.DB(tableName).C(collectionName)
	query := collection.Find(bson.M{"hdr.name": name})
	if name == "" {
		query = collection.Find(nil)
	}

	//cant RR interface here because of marshalling
	//Use any records to get the headers
	var any []*dns.ANY
	err = query.All(&any)
	if err != nil {
		return qs, err
	}

	types := make(map[uint16]bool, 0)
	for _, record := range any {
		types[record.Header().Rrtype] = true
	}

	//iterate through all types and query for each record
	for recordType := range types {
		switch recordType {
		case dns.TypeA:
			var records []*dns.A
			query.All(&records)
			for _, record := range records {
				qs = append(qs, record)
			}
		case dns.TypeAAAA:
			var records []*dns.AAAA
			query.All(&records)
			for _, record := range records {
				qs = append(qs, record)
			}
			// case dns.TypeAFSDB:
			// 	breakoutRecords(query, []*dns.AFSDB{}, "", "\t")
			// case dns.TypeANY:
			// 	breakoutRecords(query, []*dns.ANY{}, "", "\t")
			// case dns.TypeAVC:
			// 	breakoutRecords(query, []*dns.AVC{}, "", "\t")
			// case dns.TypeCAA:
			// 	breakoutRecords(query, []*dns.CAA{}, "", "\t")
			// case dns.TypeCDNSKEY:
			// 	breakoutRecords(query, []*dns.CDNSKEY{}, "", "\t")
			// case dns.TypeCDS:
			// 	breakoutRecords(query, []*dns.CDS{}, "", "\t")
			// case dns.TypeCERT:
			// 	breakoutRecords(query, []*dns.CERT{}, "", "\t")
			// case dns.TypeCNAME:
			// 	breakoutRecords(query, []*dns.CNAME{}, "", "\t")
			// case dns.TypeCSYNC:
			// 	breakoutRecords(query, []*dns.CSYNC{}, "", "\t")
			// case dns.TypeDHCID:
			// 	breakoutRecords(query, []*dns.DHCID{}, "", "\t")
			// case dns.TypeDLV:
			// 	breakoutRecords(query, []*dns.DLV{}, "", "\t")
			// case dns.TypeDNAME:
			// 	breakoutRecords(query, []*dns.DNAME{}, "", "\t")
			// case dns.TypeDNSKEY:
			// 	breakoutRecords(query, []*dns.DNSKEY{}, "", "\t")
			// case dns.TypeDS:
			// 	breakoutRecords(query, []*dns.DS{}, "", "\t")
			// case dns.TypeEID:
			// 	breakoutRecords(query, []*dns.EID{}, "", "\t")
			// case dns.TypeEUI48:
			// 	breakoutRecords(query, []*dns.EUI48{}, "", "\t")
			// case dns.TypeEUI64:
			// 	breakoutRecords(query, []*dns.EUI64{}, "", "\t")
			// case dns.TypeGID:
			// 	breakoutRecords(query, []*dns.GID{}, "", "\t")
			// case dns.TypeGPOS:
			// 	breakoutRecords(query, []*dns.GPOS{}, "", "\t")
			// case dns.TypeHINFO:
			// 	breakoutRecords(query, []*dns.HINFO{}, "", "\t")
			// case dns.TypeHIP:
			// 	breakoutRecords(query, []*dns.HIP{}, "", "\t")
			// case dns.TypeKEY:
			// 	breakoutRecords(query, []*dns.KEY{}, "", "\t")
			// case dns.TypeKX:
			// 	breakoutRecords(query, []*dns.KX{}, "", "\t")
			// case dns.TypeL32:
			// 	breakoutRecords(query, []*dns.L32{}, "", "\t")
			// case dns.TypeL64:
			// 	breakoutRecords(query, []*dns.L64{}, "", "\t")
			// case dns.TypeLOC:
			// 	breakoutRecords(query, []*dns.LOC{}, "", "\t")
			// case dns.TypeLP:
			// 	breakoutRecords(query, []*dns.LP{}, "", "\t")
			// case dns.TypeMB:
			// 	breakoutRecords(query, []*dns.MB{}, "", "\t")
			// case dns.TypeMD:
			// 	breakoutRecords(query, []*dns.MD{}, "", "\t")
			// case dns.TypeMF:
			// 	breakoutRecords(query, []*dns.MF{}, "", "\t")
			// case dns.TypeMG:
			// 	breakoutRecords(query, []*dns.MG{}, "", "\t")
			// case dns.TypeMINFO:
			// 	breakoutRecords(query, []*dns.MINFO{}, "", "\t")
			// case dns.TypeMR:
			// 	breakoutRecords(query, []*dns.MR{}, "", "\t")
			// case dns.TypeMX:
			// 	breakoutRecords(query, []*dns.MX{}, "", "\t")
			// case dns.TypeNAPTR:
			// 	breakoutRecords(query, []*dns.NAPTR{}, "", "\t")
			// case dns.TypeNID:
			// 	breakoutRecords(query, []*dns.NID{}, "", "\t")
			// case dns.TypeNIMLOC:
			// 	breakoutRecords(query, []*dns.NIMLOC{}, "", "\t")
			// case dns.TypeNINFO:
			// 	breakoutRecords(query, []*dns.NINFO{}, "", "\t")
			// case dns.TypeNS:
			// 	breakoutRecords(query, []*dns.NS{}, "", "\t")
			// case dns.TypeNSAPPTR:
			// 	breakoutRecords(query, []*dns.NSAPPTR{}, "", "\t")
			// case dns.TypeNSEC:
			// 	breakoutRecords(query, []*dns.NSEC{}, "", "\t")
			// case dns.TypeNSEC3:
			// 	breakoutRecords(query, []*dns.NSEC3{}, "", "\t")
			// case dns.TypeNSEC3PARAM:
			// 	breakoutRecords(query, []*dns.NSEC3PARAM{}, "", "\t")
			// case dns.TypeOPENPGPKEY:
			// 	breakoutRecords(query, []*dns.OPENPGPKEY{}, "", "\t")
			// case dns.TypeOPT:
			// 	breakoutRecords(query, []*dns.OPT{}, "", "\t")
			// case dns.TypePTR:
			// 	breakoutRecords(query, []*dns.PTR{}, "", "\t")
			// case dns.TypePX:
			// 	breakoutRecords(query, []*dns.PX{}, "", "\t")
			// case dns.TypeRKEY:
			// 	breakoutRecords(query, []*dns.RKEY{}, "", "\t")
			// case dns.TypeRP:
			// 	breakoutRecords(query, []*dns.RP{}, "", "\t")
			// case dns.TypeRRSIG:
			// 	breakoutRecords(query, []*dns.RRSIG{}, "", "\t")
			// case dns.TypeRT:
			// 	breakoutRecords(query, []*dns.RT{}, "", "\t")
			// case dns.TypeSIG:
			// 	breakoutRecords(query, []*dns.SIG{}, "", "\t")
			// case dns.TypeSMIMEA:
			// 	breakoutRecords(query, []*dns.SMIMEA{}, "", "\t")
			// case dns.TypeSOA:
			// 	breakoutRecords(query, []*dns.SOA{}, "", "\t")
			// case dns.TypeSPF:
			// 	breakoutRecords(query, []*dns.SPF{}, "", "\t")
			// case dns.TypeSRV:
			// 	breakoutRecords(query, []*dns.SRV{}, "", "\t")
			// case dns.TypeSSHFP:
			// 	breakoutRecords(query, []*dns.SSHFP{}, "", "\t")
			// case dns.TypeTA:
			// 	breakoutRecords(query, []*dns.TA{}, "", "\t")
			// case dns.TypeTALINK:
			// 	breakoutRecords(query, []*dns.TALINK{}, "", "\t")
			// case dns.TypeTKEY:
			// 	breakoutRecords(query, []*dns.TKEY{}, "", "\t")
			// case dns.TypeTLSA:
			// 	breakoutRecords(query, []*dns.TLSA{}, "", "\t")
			// case dns.TypeTSIG:
			// 	breakoutRecords(query, []*dns.TSIG{}, "", "\t")
			// case dns.TypeTXT:
			// 	breakoutRecords(query, []*dns.TXT{}, "", "\t")
			// case dns.TypeUID:
			// 	breakoutRecords(query, []*dns.UID{}, "", "\t")
			// case dns.TypeUINFO:
			// 	breakoutRecords(query, []*dns.UINFO{}, "", "\t")
			// case dns.TypeURI:
			// 	breakoutRecords(query, []*dns.URI{}, "", "\t")
			// case dns.TypeX25:
			// 	breakoutRecords(query, []*dns.X25{}, "", "\t")
		}
	}

	return qs, nil
}

func AcontainsA(collection []dns.A, rr dns.A) bool {
	for _, collectionRR := range collection {
		if collectionRR.A.To4().Equal(rr.A.To4()) {
			return true
		}
	}
	return false
}

func (store *MongoDB) UpdateRecord(ctx context.Context, rr dns.RR) (dns.RR, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	collection := session.DB(tableName).C(collectionName)
	err = collection.Update(bson.M{"name": rr.Header().Name}, &rr)
	if err != nil {
		if err == mgo.ErrNotFound {
			return store.CreateRecord(ctx, rr)
		} else {
			return nil, err
		}
	}
	return rr, nil
}

func (store *MongoDB) DeleteRecord(ctx context.Context, rr dns.RR) (dns.RR, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	collection := session.DB(tableName).C(collectionName)
	err = collection.Remove(bson.M{"name": rr.Header().Name})

	return rr, err
}
