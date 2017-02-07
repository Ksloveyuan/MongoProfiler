package router

import (
	"model"
	"timeutil"
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"errors"
	"fmt"
)


func  ProfileByGroupMethod(groupMethod string, startDateString string, db *mgo.Database) ([]model.ProfileSummary, error) {
	var groupID bson.M
	var startDate time.Time
	var result []model.ProfileSummary
	var err error

	if groupID, err = model.GetGroupID(groupMethod); err != nil {
		return nil, errors.New(err.Error())
	}

	if startDate, err = timeutil.ParseDate(startDateString); err != nil {
		return nil, errors.New("Invalid date or unsupported date format, e.g. valid time is 2006-01-02")
	}

	if result,err = model.Profile(db, startDate, groupID); err != nil {
		fmt.Errorf("Failed to get profile result. Details are %", err.Error())
		return nil, errors.New("Some unexpected errrors happen, pleasea contact the webstie administrator for support.")
	}

	return result, nil
}