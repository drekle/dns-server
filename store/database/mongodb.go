package database

import (
	"context"
	"log"

	impl "bitbucket.ema.emoneyadvisor.com/cp/lemondns/store"

	"github.com/miekg/dns"
	"gopkg.in/mgo.v2"
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

func (store *MongoDB) CreateRecord(ctx context.Context, rr dns.A) (*impl.Query, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	var q impl.Query
	q.Name = rr.Header().Name
	q.RR = []dns.A{rr}

	collection := session.DB(tableName).C(collectionName)
	bulk := collection.Bulk()
	bulk.Remove(bson.M{"name": q.Name})
	bulk.Insert(&q)
	_, err = bulk.Run()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &q, nil
}

func (store *MongoDB) ReadRecord(ctx context.Context, name string) (*impl.Query, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	var qs []*impl.Query
	collection := session.DB(tableName).C(collectionName)
	if name == "" {
		//All records
		collection.Find(nil).All(&qs)
	} else {
		err = collection.Find(bson.M{"name": name}).All(&qs)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	var q impl.Query
	q.Name = name
	for _, query := range qs {
		q.RR = append(q.RR, query.RR...)
	}
	return &q, nil
}

func AcontainsA(collection []dns.A, rr dns.A) bool {
	for _, collectionRR := range collection {
		if collectionRR.A.To4().Equal(rr.A.To4()) {
			return true
		}
	}
	return false
}

func (store *MongoDB) UpdateRecord(ctx context.Context, rr dns.A) (*impl.Query, error) {

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	collection := session.DB(tableName).C(collectionName)
	oldq, err := store.ReadRecord(ctx, rr.Header().Name)

	if err != nil {
		if err == mgo.ErrNotFound {
			err = nil
			//Just add
			return store.CreateRecord(ctx, rr)
		} else {
			return nil, err
		}
	}

	// update the old collection
	if !AcontainsA(oldq.RR, rr) {
		oldq.RR = append(oldq.RR, rr)
	}

	err = collection.Update(bson.M{"name": oldq.Name}, &oldq)
	if err != nil {
		if err == mgo.ErrNotFound {
			return store.CreateRecord(ctx, rr)
		} else {
			return nil, err
		}
	}
	return oldq, nil
}

func (store *MongoDB) DeleteRecord(ctx context.Context, rr dns.A) (*impl.Query, error) {
	name := rr.Header().Name

	session, err := mgo.Dial(store.address)
	if err != nil {
		log.Fatalf("Could not establish store connection: %s", store.address)
	}
	defer session.Close()

	q, err := store.ReadRecord(ctx, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var remove []int
	for index, dbrr := range q.RR {
		if rr.A.To4().String() == dbrr.A.To4().String() {
			remove = append(remove, index)
		}
	}
	var new impl.Query
	new.Name = name
	for index, rr := range q.RR {
		keep := true
		for _, removeIndex := range remove {
			if index == removeIndex {
				keep = false
			}
		}
		if keep {
			new.RR = append(new.RR, rr)
		}
	}

	collection := session.DB(tableName).C(collectionName)
	if len(new.RR) == 0 {
		err = collection.Remove(bson.M{"name": name})
		return q, err
	}
	for index, rr := range new.RR {
		if index == 0 {
			_, err = store.CreateRecord(ctx, rr)
		} else {
			_, err = store.UpdateRecord(ctx, rr)
		}
	}

	return q, err
}
