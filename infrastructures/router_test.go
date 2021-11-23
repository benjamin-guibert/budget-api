package infrastructures

import (
	"io/ioutil"
	"log"
	"testing"

	"gotest.tools/assert"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	assert.Assert(t, router != nil)
	assert.Assert(t, router.MuxRouter != nil)
}

func init() {
	log.SetOutput(ioutil.Discard)
}
