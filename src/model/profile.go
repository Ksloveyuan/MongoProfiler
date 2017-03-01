package model

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/pkg/errors"
)

const (
	collectionProfile = "system.profile"
)

type Profiler interface {
	Profile(groupMethod string, startDate time.Time) ([]ProfileSummary, error)
}


type MongoDataSource struct {
	db *mgo.Database
}

type ProfileTime struct {
	Year *int32 `bson:"year" json:",omitempty"`
	Month *int32 `bson:"month,omitempty" json:",omitempty"`
	Day *int32 `bson:"day,omitempty" json:",omitempty"`
}

type ProfileSummary struct {
	ID      ProfileTime `bson:"_id"`
	TotalMS int32 `bson:"totalMS"`
	AvgMS   float32 `bson:"avgMS"`
}

func NewMongoDataSource(db *mgo.Database) MongoDataSource  {
	return MongoDataSource{db}
}

func (ds MongoDataSource) Profile(groupMethod string, startDate time.Time) ([]ProfileSummary, error)  {
	var groupId bson.M
	var err error
	var result []ProfileSummary
	if groupId, err = getGroupID(groupMethod); err!=nil {
		return  result, err
	}

	return profile(ds.db, startDate, groupId)

}

func getGroupID(groupMethod string) (bson.M, error){
	var err error

	groupIDMap := map[string] bson.M {
		"byyear": {"year": "$year"},
		"bymonth": {"year": "$year", "month": "$month"},
		"byday": {"day":"$day", "month": "$month", "year": "$year"},
	}

	id,ok := groupIDMap[groupMethod]

	if !ok {
		err = errors.Errorf("The group method(%s) is not supported.", groupMethod)
	}

	return id,err
}

func matchGreaterThan(startDate time.Time) bson.M{
	return bson.M{"$match": bson.M{"ts": bson.M{"$gt": startDate}}}
}

func project() bson.M {
	return bson.M{
		"$project": bson.M{
			"year": bson.M{"$year": "$ts"},
			"month": bson.M{"$month": "$ts"},
			"day": bson.M{"$dayOfMonth": "$ts"},
			"millis": "$millis",
		},
	}
}

func groupBy(groupID bson.M) bson.M {
	return bson.M{
		"$group": bson.M{
			"_id" : groupID,
			"totalMS": bson.M{"$sum": "$millis"},
			"avgMS": bson.M{"$avg": "$millis"},
		},
	}
}

func profile(db *mgo.Database, startDate time.Time, groupID bson.M) ([]ProfileSummary, error) {

	pipeline := [] bson.M{matchGreaterThan(startDate), project(), groupBy(groupID)}

	c := db.C(collectionProfile)

	var result []ProfileSummary

	err := c.Pipe(pipeline).All(&result)

	if err != nil {
		err = errors.Wrapf(err, "Some unexpected errrors happen, pleasea contact the webstie administrator for support.")
	}

	return result, err
}