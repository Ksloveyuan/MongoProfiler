package router_test

import (
	"github.com/stretchr/testify/mock"
	"time"
	"model"
	"testing"
	"router"
	"github.com/stretchr/testify/assert"
	"timeutil"
	"errors"
)

type MockedProfiler struct{
	mock.Mock
}

const ProfileMethodName  =  "Profile"

func(mockerProfiler MockedProfiler) Profile(groupMethod string, startDate time.Time) ([]model.ProfileSummary, error){
	var result []model.ProfileSummary

	args := mockerProfiler.Called(groupMethod, startDate)
	firstArg := args.Get(0)

	if firstArg == nil {
		result = nil
	}else {
		result = firstArg.([]model.ProfileSummary)
	}

	return result, args.Error(1)
}

func TestProfileByGroupMethod_WithWrongTimeFormat_ReturnError(t *testing.T) {
	groupMehtod, wrongTimeFormat := "bymonth", "2011:55:44"
	result, err := router.ProfileByGroupMethod(groupMehtod, wrongTimeFormat, nil)

	assert.Error(t, err, "Invalid date or unsupported date format, e.g. valid time is 2006-01-02")
	assert.Nil(t, result)
}

func TestProfileByGroupMethod_IfProfilerMetError_ReturnError(t *testing.T) {
	groupMehtod, dateString := "bymonth", "2011-1-2"
	date,_ := timeutil.ParseDate(dateString)
	fakedProfiler := new(MockedProfiler)
	fakedProfiler.On(ProfileMethodName, groupMehtod, date).Return(nil, errors.New("Some error happens"))
	result, err := router.ProfileByGroupMethod(groupMehtod, dateString, fakedProfiler)

	assert.Error(t, err,"Some error happens")
	assert.Nil(t, result)
}

func TestProfileByGroupMethod_ReturnRightResult(t *testing.T) {
	groupMehtod, dateString, fakedReults := "bymonth", "2011-1-2", make([]model.ProfileSummary, 1)
	date,_ := timeutil.ParseDate(dateString)
	fakedProfiler := new(MockedProfiler)
	fakedProfiler.On(ProfileMethodName, groupMehtod, date).Return(fakedReults, nil)
	result, err := router.ProfileByGroupMethod(groupMehtod, dateString, fakedProfiler)

	assert.NoError(t, err)
	assert.Equal(t, fakedReults, result)
}

