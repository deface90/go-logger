package defaultfield

import (
	"github.com/sirupsen/logrus"
)

type Hook struct {
	Fields map[string]string
}

func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *Hook) Fire(e *logrus.Entry) error {
	for key, value := range h.Fields {
		e.Data[key] = value
	}
	return nil
}

func NewHook(fields map[string]string) *Hook {
	hook := Hook{
		Fields: fields,
	}
	return &hook
}
