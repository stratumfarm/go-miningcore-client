package miningcore

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testServer *httptest.Server
)

func TestBuildRequestUrl(t *testing.T) {
	url, err := buildRequestURL("http://localhost:8080", "/api/pools", map[string]string{"id": "1"})
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/api/pools?id=1", url)

	url, err = buildRequestURL("http://localhost:8080", "/api/pools", map[string]string{"id": "1", "name": "2"})
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/api/pools?id=1&name=2", url)

	url, err = buildRequestURL("http://localhost:8080", "/api/pools", map[string]string{"i": "Day"})
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/api/pools?i=Day", url)

}

func TestMain(m *testing.M) {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/pools", poolsReq)
	handler.HandleFunc("/api/pools/eth", poolReq)
	handler.HandleFunc("/api/pools/mock", poolMock)

	testServer = httptest.NewServer(handler)
	defer testServer.Close()

	code := m.Run()
	os.Exit(code)
}

func newClient() *Client {
	return New(testServer.URL, WithoutTLSVerfiy())
}

func poolsReq(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("testdata/pools.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func poolReq(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("testdata/pool_eth.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func poolMock(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}

func TestPools(t *testing.T) {
	pools, code, err := newClient().GetPools(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, 1, len(pools))
}

func TestCustomJSON(t *testing.T) {
	t.Cleanup(func() {
		newClient() // reset the client
	})
	client := New(testServer.URL, WithJSONEncoder(json.Marshal), WithJSONDecoder(json.Unmarshal))
	_, code, err := client.GetPools(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, code)
}

func TestPool(t *testing.T) {
	pool, code, err := newClient().GetPool(context.Background(), "eth")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "eth", pool.ID)
}

func TestPoolMock(t *testing.T) {
	pool, code, err := newClient().GetPool(context.Background(), "mock")
	assert.Error(t, err)
	assert.Equal(t, http.StatusForbidden, code)
	assert.Nil(t, pool)
}
