package log

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMsgToLong(t *testing.T) {
	log := New("logger-test:", "")
	log.Log("it works! to long not to long to long not to long but it is too long or it isnt't na it isn't that long :)")
	log.Log("here we go")
	time.Sleep(time.Microsecond * 1)
	// sleep above is to simulate the cost of doing somthing, sending to syslog
	// has a bit of startup overhead so if there must be a little work to be
	// done otherwise the message won't get send since the however still negletable
	// in normal apps.
	log.Close()
}

func TestFileLogger(t *testing.T) {
	log := New("logger-test:", "local.log")
	log.Log("It works!")
	log.Close()
	os.Remove("local.log")
}

func TestFileLoggerManyMsg(t *testing.T) {
	log := New("logger-test:", "manylocal.log")
	for i := 0; i < 33; i++ {
		log.Log("It works! :" + strconv.Itoa(i))
	}
	log.Close()
	os.Remove("manylocal.log")
}
