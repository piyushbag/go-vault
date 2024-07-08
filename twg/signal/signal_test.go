package signal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler.
	// We don't have any query parameters for now, so we'll pass 'nil' as the third parameter.
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %v; want nil", err)
	}
	Handler(w, r)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != 200 {
		t.Fatalf("w.Result().StatusCode = %v; want 200", resp.StatusCode)
	}

	// Check the content type is what we expect.
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type = %v; want application/json", contentType)
	}

	// Read the body and check it is what we expect.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("io.ReadAll() err = %v; want nil", err)
	}
	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal() err = %v; want nil", err)
	}
	if p.Age != 25 {
		t.Errorf("p.Age = %v; want 25", p.Age)
	}
	if p.Name != "John Doe" {
		t.Errorf("p.Name = %v; want John Doe", p.Name)
	}
}
