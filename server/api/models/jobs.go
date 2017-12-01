package models

import (
	"errors"

	"github.com/zcrawl/zcrawl/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	jobsCollectionName = "jobs"
)

// Job is an alias for types.Job
type Job types.Job

// Get retrieves a project item.
func (c *Job) Get(id string) error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(jobsCollectionName)
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
func (c *Job) Save() error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(jobsCollectionName)
	c.ID = bson.NewObjectId()
	return collection.Insert(c)
}
