package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getAndUnmarshalJson(t *testing.T, path string, v interface{}) {
	content, ts := getContents(t, path)
	defer ts.Close()
	unmarshalJson(t, content, v)
}

func getContents(t *testing.T, path string) ([]byte, *httptest.Server) {
	ts := httptest.NewServer(defineRoutes())
	res, err := http.Get(ts.URL + path)
	if err != nil {
		t.Errorf("Unable to GET URL %s", path)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Unable to read body contents")
	}

	return content, ts
}

func unmarshalJsonFromResponse(t *testing.T, res *http.Response, v interface{}) {
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Unable to read body contents")
	}
	unmarshalJson(t, content, v)
}

func unmarshalJson(t *testing.T, content []byte, v interface{}) {
	err := json.Unmarshal(content, v)
	if err != nil {
		t.Errorf("Unable to unmarshal result, got: [%s] %s", content, err)
	}
}

func assertIntEquals(t *testing.T, expected, result int) {
	if expected != result {
		t.Errorf("Expected [%d], got [%d]", expected, result)
	}

}

func assertEquals(t *testing.T, expected, result interface{}) {
	if expected != result {
		t.Errorf("Expected [%s], got [%s]", expected, result)
	}
}
