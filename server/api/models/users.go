package models

import (
	"errors"

	"github.com/zcrawl/zcrawl/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	usersCollectionName = "users"
)

// User is an alias for types.Users
type User types.User

// UsersCollection is an alias for a collection of users
type UsersCollection []types.User

// GetAll retrieves all users
func (p *UsersCollection) GetAll() error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(usersCollectionName)
	return collection.Find(bson.M{}).All(p)
}

// Get retrieves a user item.
func (u *User) Get(id string) error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(usersCollectionName)
	// TODO: handle error
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}
	objectID := bson.ObjectIdHex(id)
	err := collection.FindId(objectID).One(u)
	if err != nil {
		return err
	}
	return nil
}

// Save stores a user item.
func (u *User) Save() error {
	session := mongoSession.Clone()
	defer session.Close()
	collection := session.DB(mongoDialInfo.Database).C(usersCollectionName)
	u.ID = bson.NewObjectId()
	return collection.Insert(u)
}
