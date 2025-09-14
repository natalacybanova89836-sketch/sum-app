package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit test для Add
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %d; want 5", got)
	}
}

// Handler test используя httptest
func TestSumHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/sum?a=2&b=3", nil)
	w := httptest.NewRecorder()

	sumHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d; want 200; body: %s", resp.StatusCode, string(body))
	}

	var m map[string]int
	if err := json.Unmarshal(body, &m); err != nil {
		t.Fatalf("json decode error: %v; body: %s", err, string(body))
	}
	if m["sum"] != 5 {
		t.Fatalf("sum = %d; want 5", m["sum"])
	}
}
