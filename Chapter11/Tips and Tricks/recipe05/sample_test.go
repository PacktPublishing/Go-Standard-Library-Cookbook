package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const cookieName = "X-Cookie"

func HandlerUnderTest(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Domain:  "localhost",
		Expires: time.Now().Add(3 * time.Hour),
		Name:    cookieName,
	})
	r.ParseForm()
	username := r.FormValue("username")
	fmt.Fprintf(w, "Hello %s!", username)
}

func TestHttpRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "http://unknown.io?username=John", nil)
	w := httptest.NewRecorder()
	HandlerUnderTest(w, req)

	var res *http.Cookie
	for _, c := range w.Result().Cookies() {
		if c.Name == cookieName {
			res = c
		}
	}

	if res == nil {
		t.Fatal("Cannot find " + cookieName)
	}

	content, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal("Cannot read response body")
	}

	if string(content) != "Hello John!" {
		t.Fatal("Content not matching expected value")
	}
}
