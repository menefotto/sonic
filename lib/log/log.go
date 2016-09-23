// Package log provides a New function that given a prefix string and a file or
// lack of empty string "" builds either a syslog or a filename named local log.
// It provides 2 functions and one of them its called Log, which taken a strings
// logs it, if the string is longer than 80 characters is going to be cut at 80
// characters.
// The other one is Close which must be called when the logger isn't needed any
// longer, IT'S IMPORTANT otherwise message may get lost and never reach the log
// whether is a syslog or local file.
// The main difference between other log package or the standard library one is
// to provide a simpler to use interface, and "ELIMINATE THE COST usually
// associated with the standard logging methods". All of this is achieved by letting
// a background cooroutine to the actual logging so the every call to Log is as
// much expensive as a message sent to a go channel.

package log

import (
	"log"
	"log/syslog"
	"os"
	"time"
)

func New(prefix string, filename string) *Logger {

	var (
		stdlog  *log.Logger
		err     error
		flag    int = log.Lshortfile
		logfile *os.File
	)

	switch {
	case filename == "":
		priority := syslog.LOG_EMERG | syslog.LOG_USER
		stdlog, err = syslog.NewLogger(priority, flag)
		if err != nil {
			panic(err)
		}
	default:
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
		logfile, err = os.OpenFile(filename, flag, 0660)
		if err != nil {
			panic(err)
		}
		stdlog = log.New(logfile, prefix, flag)
	}

	stdlog.SetPrefix(prefix)

	log := &Logger{
		logger:   stdlog,
		Messages: make(chan string, 1024),
		Done:     make(chan struct{}, 1),
	}

	go func(log *Logger) {
		for {
			select {
			case <-log.Done:

				close(log.Messages)
				for msg := range log.Messages {
					log.log(msg)
				}

				time.Sleep(time.Millisecond * 250)

				defer func() {
					if logfile != nil {
						logfile.Close()
					}
				}()

				return
			case msg := <-log.Messages:
				log.log(msg)
			}
		}
	}(log)

	return log

}

type Logger struct {
	logger   *log.Logger
	Messages chan string
	Done     chan struct{}
}

// Log does what it says but it doesn't assure the message will be send.
func (s *Logger) Log(msg string) {
	s.toLog(msg)
}

// Typically called in defer log.Close() fashion for short lived objects otherwise
// must be called when the logger isn't needed any longer, once called if any
// messages are still left on the Messages channel, they will be sent to syslog before
// actually exting, any message sent after the close method is called will a result
// in a sent operation to a close channel.

func (s *Logger) Close() {
	s.Done <- struct{}{}
}

const (
	msgMaxLen = 79
)

// log not uppercase is used internaly to actually print to syslog
func (s *Logger) log(msg string) {
	s.logger.Output(2, msg)
}

// clients code should use uppercase Log to send messages to the listeining goroutine
func (s *Logger) toLog(msg string) {
	truncated := []byte(msg)

	if len(msg) > msgMaxLen {
		// msg should must not be over 79 characthers so panic
		truncated = truncated[:79]
	}
	s.Messages <- string(truncated)
}
