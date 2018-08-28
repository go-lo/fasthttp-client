package fasthttpclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"hello": "world"}`)
	}))
	defer ts.Close()

	c := New()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
}
