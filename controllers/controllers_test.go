package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

type TestWriter struct {
	http.ResponseWriter
}

func TestGetVarToInt(t *testing.T) {
	writer := &TestWriter{}
	vars := map[string]string{
		"key": "123",
	}

	value, err := GetVarToInt(writer, vars, "key")

	assert.Equal(t, 123, value)
	assert.Assert(t, err == nil, err)
}

func init() {
	log.SetOutput(ioutil.Discard)
}
