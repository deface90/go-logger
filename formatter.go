package logrussentry

import (
	log "github.com/sirupsen/logrus"
)

// formatter adds default fields to each log entry.
type Formatter struct {
	Fields log.Fields
	LF     log.Formatter
}

// Format satisfies the logrus.Formatter interface.
func (f *Formatter) Format(e *log.Entry) ([]byte, error) {
	for k, v := range f.Fields {
		e.Data[k] = v
	}
	return f.LF.Format(e)
}
