package defaultfield

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFilenameHook(t *testing.T) {
	hook := NewHook(map[string]string{
		"test": "test",
	})

	buff := new(bytes.Buffer)
	logrus.SetOutput(buff)
	logrus.AddHook(hook)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Info("Test")

	require.NotEqual(t, "", buff.String())
	require.Equal(t,
		"test",
		gjson.Get(buff.String(), "test").Str,
	)
}
