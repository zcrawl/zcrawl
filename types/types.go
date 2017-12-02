package types

import "gopkg.in/mgo.v2/bson"

// Project represents a project.
type Project struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
}

// Crawler represents a crawler.
type Crawler struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	ProjectID   bson.ObjectId `bson:"project_id" json:"project_id"`
	Jobs        []Job         `bson:",omitempty" json:",omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
}

// Job represents a job.
type Job struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	Project   *Project      `bson:",omitempty" json:",omitempty"`
	CrawlerID bson.ObjectId `bson:"crawler_id" json:"crawler_id"`
}

// StoreRecord wraps a dynamic data structure.
type StoreRecord struct {
	ID             bson.ObjectId `bson:"_id" json:"_id"`
	JobID          bson.ObjectId `bson:"job_id" json:"job_id"`
	CollectionName string        `bson:"collection_name" json:"collection_name"`
	Data           map[string]interface{}
}

// User represents a user.
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
}
