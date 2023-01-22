package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()

	// API からのレスポンスを用いてなんならの処理をする

	if res.StatusCode != 200 {
		t.Errorf("GET %s: expected status code = %d; got %d", ts.URL, 200, res.StatusCode)
	}
	if string(got) != "Hello, client\n" {
		t.Errorf("expected body %v; got %v", "Hello, client", string(got))
	}
}