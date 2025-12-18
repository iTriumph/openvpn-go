package openvpn

import "log"

type Logger interface {
	Debugf(format string, args ...any)
}

var logger Logger = nil

func SetLogger(l Logger) {
	logger = l
}

type SimpleLogger struct {
}

func (_ *SimpleLogger) Debugf(format string, args ...any) {
	log.Printf(format, args)
}
