package glog

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLogger(t *testing.T) {
	l := &logrus.Logger{
		Out:   ioutil.Discard,
		Level: logrus.PanicLevel,
	}
	SetLogger(l)

	assert.Equal(t, l, logger)
}

func TestV(t *testing.T) {
	assert.True(t, bool(V(0)))
}

func TestInfo(t *testing.T) {
	formatter := &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
	var buf bytes.Buffer
	l := &logrus.Logger{
		Out:       &buf,
		Formatter: formatter,
		Level:     logrus.InfoLevel,
	}
	SetLogger(l)

	testcases := []struct {
		Name     string
		Format   string
		Args     []interface{}
		Expected string
	}{{
		Name:     "String",
		Format:   "foo",
		Args:     []interface{}{},
		Expected: `{"msg":"foo", "func":"Infof", "level":"info"}`,
	}, {
		Name:     "Int",
		Format:   "%d",
		Args:     []interface{}{42},
		Expected: `{"msg":"42", "func":"Infof", "level":"info"}`,
	}, {
		Name:     "Bool",
		Format:   "%t",
		Args:     []interface{}{true},
		Expected: `{"msg":"true", "func":"Infof", "level":"info"}`,
	}, {
		Name:     "StringInt",
		Format:   "%s=%d",
		Args:     []interface{}{"foo", 42},
		Expected: `{"msg":"foo=42", "func":"Infof", "level":"info"}`,
	}}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			Infof(tc.Format, tc.Args...)
			assert.JSONEq(t, tc.Expected, buf.String())
			buf.Reset()
		})
	}
}
