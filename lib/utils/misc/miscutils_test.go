package misc

import (
	"strings"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	date := GetDate()
	datesplit := strings.Split(date, " ")

	_, month, _ := time.Now().Date()

	if datesplit[1] != month.String() {
		t.Fatal("Month should not be different")
	}
}
