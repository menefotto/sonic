package futils

import "github.com/sonic/lib/log"

type FileMoverEntry struct {
	Source, Destination string
}

func NewFileMoverEntry(s, d string) *FileMoverEntry {
	return &FileMoverEntry{Source: s, Destination: d}
}

type FileMover struct {
	Entries chan *FileMoverEntry
	Done    chan bool
	Log     *log.Logger
}

func NewFileMover(log *log.Logger) *FileMover {

	mover := &FileMover{
		Entries: make(chan *FileMoverEntry, 32),
		Done:    make(chan bool, 1),
		Log:     log,
	}

	go func(m *FileMover) {
		finalize := func(m *FileMover) {
			close(m.Entries)
			for e := range m.Entries {
				err := MoveFile(e.Source, e.Destination)
				if err != nil {
					log.Log(err.Error())
				}
			}

		}
		for {
			select {
			case entry := <-m.Entries:
				err := MoveFile(entry.Source, entry.Destination)
				if err != nil {
					log.Log(err.Error())
				}
			case <-m.Done:
				finalize(m)
				return
			}
		}
	}(mover)

	return mover
}

func (m *FileMover) Close() {
	m.Done <- true
}

func (m *FileMover) Send(origin, destination string) {
	m.Entries <- NewFileMoverEntry(origin, destination)
}
