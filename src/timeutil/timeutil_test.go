package timeutil_test

import ("testing"
	"timeutil"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestParseDate(t *testing.T)  {
	dateString := "2017-01-02"

	result, err := timeutil.ParseDate(dateString)

	assert.Equal(t, nil, err)
	assert.Equal(t, 2017, result.Year())
	assert.Equal(t, time.January, result.Month())
	assert.Equal(t, 2, result.Day())
}

func TestToString(t *testing.T) {
	date := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	resultString := timeutil.ToString(date)

	assert.Equal(t, "2009-11-10", resultString)
}