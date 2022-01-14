package torion

import (
	"runtime"
	"time"
)

func (config *Configuration) New() {
	runtime.LockOSThread()
}

func Default() *Configuration {
	return &Configuration{
		BootstrapTimeout: time.Time{},
		DirectAddress:    "",
		CustomTorrc:      "",
		UseMirror:        false,
		StdoutLogger:     false,
	}
}
