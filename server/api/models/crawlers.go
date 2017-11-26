package models

import (
	"errors"

	"github.com/zcrawl/zcrawl/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	crawlersCollectionName = "crawlers"
)

// Crawler is an alias for types.Crawler
type Crawler types.Crawler

// Get retrieves a project item.
func (c *Crawler) Get(id string) error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(crawlersCollectionName)
	// TODO: handle error
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}
	objectID := bson.ObjectIdHex(id)
	err := collection.FindId(objectID).One(c)
	if err != nil {
		return err
	}
	return nil
}

// Save stores a project item.
func (c *Crawler) Save() error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(crawlersCollectionName)
	c.ID = bson.NewObjectId()
	return collection.Insert(c)
}
