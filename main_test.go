package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {

	router := newRouter()

	mockServer := httptest.NewServer(router)

	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok but got %v", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "Hi, the world is beautiful"
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestRouterNonExistentRoute(t *testing.T) {
	router := newRouter()

	mockServer := httptest.NewServer(router)

	resp, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.StatusCode)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405 but got %v", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
