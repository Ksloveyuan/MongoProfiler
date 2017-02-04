package model

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

const (
	collectionProfile = "system.profile"
)

type ProfileTime struct {
	Year *int32 `bson:"year" json:",omitempty"`
	Month *int32 `bson:"month,omitempty" json:",omitempty"`
	Day *int32 `bson:"day,omitempty" json:",omitempty"`
}

type ProfileSummary struct {
	ID ProfileTime `bson:"_id"`
	TotalMS int32 `bson:"totalMS"`
	AvgMS float32 `bson:"avgMS"`
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

func GetGroupID(groupMethod string) (bson.M, error){
	var err error

	groupIDMap := map[string] bson.M {
		"byyear": {"year": "$year"},
		"bymonth": {"year": "$year", "month": "$month"},
		"byday": {"day":"$day", "month": "$month", "year": "$year"},
	}

	id,ok := groupIDMap[groupMethod]
	if !ok {
		err = errors.New("The group method is not supported")
	}

	return id,err
}

func Profile(db *mgo.Database, startDate time.Time, groupID bson.M) ([]ProfileSummary, error) {

	pipeline := [] bson.M{matchGreaterThan(startDate), project(), groupBy(groupID)}

	c := db.C(collectionProfile)

	var result []ProfileSummary

	err := c.Pipe(pipeline).All(&result)

	return result, err
}