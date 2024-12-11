package example

import (
	"example/mocks/example"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFromDB(t *testing.T) {
	mockDB := mocks.NewMockDB(t)
	mockDB.EXPECT().Get("ice cream").Return("chocolate").Once()
	flavor := getFromDB(mockDB)
	assert.Equal(t, "chocolate", flavor)
}
