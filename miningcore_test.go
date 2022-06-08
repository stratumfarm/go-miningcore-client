package miningcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRequestUrl(t *testing.T) {
	url, err := buildRequestUrl("http://localhost:8080", "/api/pools", map[string]string{"id": "1"})
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/api/pools?id=1", url)

	url, err = buildRequestUrl("http://localhost:8080", "/api/pools", map[string]string{"id": "1", "name": "2"})
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/api/pools?id=1&name=2", url)

	url, err = buildRequestUrl("http://localhost:8080", "/api/pools", map[string]string{"i": "Day"})
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/api/pools?i=Day", url)

}
