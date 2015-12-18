package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	Root(w, r)
	if w.Code != http.StatusOK {
		t.Fail()
	}
}
