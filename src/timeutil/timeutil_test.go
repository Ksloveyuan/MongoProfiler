package timeutil_test

import ("testing"
	"timeutil"
	"time"

)

func TestParseDate(t *testing.T)  {
	dateString := "2017-01-02"
	expectedMonth := time.January

	result, err := timeutil.ParseDate(dateString)


	if err != nil {
		t.Error("Failed")
	}else if result.Month() != expectedMonth{
		t.Error("Not expected month")
	}
}
