package timeutil

import (
	"time"
	"github.com/gin-gonic/gin"
)

const TimeLayout  = "2006-1-2"

func ParseDate(dateString string) (time.Time, error)  {
	return time.Parse(TimeLayout, dateString)
}

func ToString(t time.Time) string  {
	return t.Format(TimeLayout)
}

func LastYearOfToday() time.Time {
	now := time.Now()
	lastYear := now.AddDate(-1,0,0)
	return lastYear
}
