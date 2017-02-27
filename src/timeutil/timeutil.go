package timeutil

import (
	"time"
	"github.com/pkg/errors"
)

const TimeLayout  = "2006-1-2"

func ParseDate(dateString string) (time.Time, error)  {
	date,err := time.Parse(TimeLayout, dateString)
	if err != nil{
		err = errors.Wrapf(err,"{s} is an invalid date or it is in a unsupported date format, e.g. valid example is 2006-01-02", dateString)
	}

	return date, err
}

func ToString(t time.Time) string  {
	return t.Format(TimeLayout)
}

func LastYearOfToday() time.Time {
	now := time.Now()
	lastYear := now.AddDate(-1,0,0)
	return lastYear
}
