package rabbitmq

import "log"

// Debug debug log flag
var Debug bool

func debug(args ...interface{}) {
	if !Debug {
		return
	}
	log.Print(args...)
}

func debugf(format string, args ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(format, args...)
}
