package misc

import (
	"fmt"
	"time"
)

func GetDate() string {

	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()

	zone, _ := time.Now().Zone()
	return fmt.Sprintf("%d %s %d, %d:%d:%d %s",
		day, month.String(), year, hour, min, sec, zone)

}
