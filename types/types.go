package types

import "gopkg.in/mgo.v2/bson"

// Project represents a project.
type Project struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Jobs        []Job         `bson:",omitempty" json:",omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
}

// Job represents a job.
type Job struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Project *Project      `bson:",omitempty" json:",omitempty"`
}
