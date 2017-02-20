package router

import (
	"model"
	"timeutil"
	"time"
	"errors"
	"fmt"
)


func ProfileByGroupMethod(groupMethod string, startDateString string, profiler model.Profiler) ([]model.ProfileSummary, error) {
	var startDate time.Time
	var result []model.ProfileSummary
	var err error

	if startDate, err = timeutil.ParseDate(startDateString); err != nil {
		return nil, errors.New("Invalid date or unsupported date format, e.g. valid time is 2006-01-02")
	}

	if result,err = profiler.Profile(groupMethod, startDate); err != nil {
		fmt.Errorf("Failed to get profile result. Details are %", err.Error())
		return nil, errors.New("Some unexpected errrors happen, pleasea contact the webstie administrator for support.")
	}

	return result, nil
}