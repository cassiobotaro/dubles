package examplerequests

import (
	"context"
	"testing"

	"github.com/carlmjohnson/requests"
	"github.com/carlmjohnson/requests/reqtest"
	"github.com/stretchr/testify/assert"
)

func Test_getFromDB(t *testing.T) {
	ctx := context.Background()
	// record a request to the file system
	var s1, s2 string
	err := requests.URL("http://example.com").
		Transport(reqtest.Record(nil, "somedir")).
		ToString(&s1).
		Fetch(ctx)
	assert.NoError(err)

	// now replay the request in tests
	err = requests.URL("http://example.com").
		Transport(reqtest.Replay("somedir")).
		ToString(&s2).
		Fetch(ctx)
	assert.NoError(err)
	assert.True(s1 == s2)
}
