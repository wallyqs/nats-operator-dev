package natsoperator

import "fmt"

// Logger is the minimal interface that a logger
// has to implement.
type Logger interface {
	Log(keyvals ...interface{}) error
}

// Noticef can be used to log events information.
func (op *Operator) Noticef(text string, params ...interface{}) {
	op.logger.Log("level", "info", "msg", fmt.Sprintf(text, params...))
}

// Fatalf can be used for logging events where a critical conditions
// after which component cannot longer continue.
func (op *Operator) Fatalf(text string, params ...interface{}) {
	op.logger.Log("level", "fatal", "msg", fmt.Sprintf(text, params...))
}

// Errorf can be used to logging errors and warnings events.
func (op *Operator) Errorf(text string, params ...interface{}) {
	op.logger.Log("level", "error", "msg", fmt.Sprintf(text, params...))
}

// Debugf can be used for logging debug events.
func (op *Operator) Debugf(text string, params ...interface{}) {
	if op.debug {
		op.logger.Log("level", "debug", "msg", fmt.Sprintf(text, params...))
	}
}

// Tracef can be used for logging trace events.
func (op *Operator) Tracef(text string, params ...interface{}) {
	// Use it to track the requests sent to Kubernetes back and forth.
	if op.trace {
		op.logger.Log("level", "trace", "msg", fmt.Sprintf(text, params...))
	}
}

// emptyLogger is a logger that does not log.
type emptyLogger struct{}

func (*emptyLogger) Log(keyvals ...interface{}) error {
	return nil
}
