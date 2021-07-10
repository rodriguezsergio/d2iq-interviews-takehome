package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rodriguezsergio/d2iq-interviews-takehome/lru"
)

func TestServer_Add_GetKey(t *testing.T) {
	sharedServer := &Server{
		MaxItems: 1,
		Cache:    make(map[string]string),
		LRU:      &lru.DoublyLinkedList{},
	}

	addTests := map[string]struct {
		Input      []byte
		Response   string
		Server     *Server
		StatusCode int
	}{
		"400_1": {
			Input:    []byte(`{"malformed": input"}`),
			Response: `{"error":"bad input provided"}`,
			Server: &Server{
				Cache:    make(map[string]string),
				LRU:      &lru.DoublyLinkedList{},
				MaxItems: 1,
			},
			StatusCode: 400,
		},
		"400_2": {
			Input:    []byte(`{"numbers_not_supported": 123}`),
			Response: `{"error":"bad input provided"}`,
			Server: &Server{
				Cache:    make(map[string]string),
				LRU:      &lru.DoublyLinkedList{},
				MaxItems: 1,
			},
			StatusCode: 400,
		},
		"200_1": {
			Input:      []byte(`{"foo": "bar"}`),
			Response:   `{"status":"key(s) successfully added"}`,
			Server:     sharedServer,
			StatusCode: 200,
		},
		"200_2": {
			Input:      []byte(`{"fizz": "buzz"}`),
			Response:   `{"status":"key(s) successfully added"}`,
			Server:     sharedServer,
			StatusCode: 200,
		},
	}

	for name, test := range addTests {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/key", bytes.NewReader(test.Input))

			test.Server.AddKey(w, req)
			res := w.Result()
			if res.StatusCode != test.StatusCode {
				t.Errorf("expected %v got %v", test.StatusCode, res.StatusCode)
			}

			if w.Body.String() != test.Response {
				t.Errorf("expected %v got %v", test.Response, w.Body.String())
			}
		})
	}

	getTests := map[string]struct {
		Key        string
		Response   string
		Server     *Server
		StatusCode int
	}{
		"200": {
			Key:        "fizz",
			Response:   `{"key":"fizz","value":"buzz"}`,
			Server:     sharedServer,
			StatusCode: 200,
		},
		"404": {
			Key:        "foo",
			Response:   `{"error":"key not found"}`,
			Server:     sharedServer,
			StatusCode: 404,
		},
	}

	for name, test := range getTests {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/key/%v", test.Key), nil)
			// if this is not included mux.Vars(r) will return nothing
			req = mux.SetURLVars(req, map[string]string{"key": test.Key})

			test.Server.GetKey(w, req)
			res := w.Result()
			if res.StatusCode != test.StatusCode {
				t.Errorf("expected %v got %v", test.StatusCode, res.StatusCode)
			}

			if w.Body.String() != test.Response {
				t.Errorf("expected %v got %v", test.Response, w.Body.String())
			}
		})
	}
}
