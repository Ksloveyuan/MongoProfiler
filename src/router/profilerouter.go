package router

import (
	"model"
	"timeutil"
)

func ProfileByGroupMethod(groupMethod string, startDateString string, profiler model.Profiler) ([]model.ProfileSummary, error) {
	var result []model.ProfileSummary

	startDate, err := timeutil.ParseDate(startDateString)

	if  err == nil {
		result, err = profiler.Profile(groupMethod, startDate)
	}

	return result, err
}