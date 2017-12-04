package models

import (
	"errors"

	"github.com/zcrawl/zcrawl/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	projectsCollectionName = "projects"
)

// Project is an alias for types.Project
type Project types.Project

// ProjectsCollection is an alias for a collection of projects
type ProjectsCollection []types.Project

// GetAll retrieves all projects
func (p *ProjectsCollection) GetAll() error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(projectsCollectionName)
	err := collection.Find(bson.M{}).All(p)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a project item.
func (p *Project) Get(id string) error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(projectsCollectionName)
	// TODO: handle error
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}
	objectID := bson.ObjectIdHex(id)
	err := collection.FindId(objectID).One(p)
	if err != nil {
		return err
	}
	return nil
}

// Save stores a project item.
func (p *Project) Save() error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(projectsCollectionName)
	p.ID = bson.NewObjectId()
	return collection.Insert(p)
}
