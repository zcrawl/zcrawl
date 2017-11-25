package models

import (
	"gopkg.in/mgo.v2"
)

const (
	defaultDatabaseName = "zcrawl"
)

var (
	mongoSession  *mgo.Session
	mongoDialInfo *mgo.DialInfo
)

// DialMongo establishes the database connection.
func DialMongo(mongoURL string) error {
	var err error
	mongoDialInfo, err = mgo.ParseURL(mongoURL)
	if err != nil {
		return err
	}
	if mongoDialInfo.Database == "" {
		mongoDialInfo.Database = defaultDatabaseName
	}
	mongoSession, err = mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		return err
	}
	return nil
}
