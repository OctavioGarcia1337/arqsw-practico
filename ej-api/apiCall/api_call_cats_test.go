package apiCall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	result, _ := ApiCall()
	assert.NotEmpty(t, result)
	for _, r := range result {
		assert.NotNil(t, r.URL)
		assert.NotNil(t, r.Height)
	}
}

func TestOldURL(t *testing.T) {
	resp, _ := ApiCall()
	if resp != nil {
		return
	}
}
