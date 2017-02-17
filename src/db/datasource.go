package db

import "gopkg.in/mgo.v2"

type DataSource interface {
	GetMongoDataSource() (*mgo.Database)
}