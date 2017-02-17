package router

import (
	"model"
	"timeutil"
	"time"
	"errors"
	"fmt"
	"db"
)


func  ProfileByGroupMethod(groupMethod string, startDateString string, dataSource db.DataSource) ([]model.ProfileSummary, error) {
	var startDate time.Time
	var result []model.ProfileSummary
	var err error

	mongoDataSource := model.NewMongoDataSource(dataSource.GetMongoDataSource())

	if startDate, err = timeutil.ParseDate(startDateString); err != nil {
		return nil, errors.New("Invalid date or unsupported date format, e.g. valid time is 2006-01-02")
	}

	if result,err = mongoDataSource.Profile(groupMethod, startDate); err != nil {
		fmt.Errorf("Failed to get profile result. Details are %", err.Error())
		return nil, errors.New("Some unexpected errrors happen, pleasea contact the webstie administrator for support.")
	}

	return result, nil
}