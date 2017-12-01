package models

import (
	"fmt"

	"github.com/zcrawl/zcrawl/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	storeCollectionName = "store"
)

// StoreRecord is an alias.
type StoreRecord types.StoreRecord

// Save stores a project item.
func (s *StoreRecord) Save() error {
	collectionName := fmt.Sprintf("%s_%s", s.JobID.Hex(), s.CollectionName)
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(collectionName)
	s.ID = bson.NewObjectId()
	return collection.Insert(s.Data)
}
