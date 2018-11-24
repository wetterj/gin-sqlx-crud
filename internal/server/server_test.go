// +build integration

package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/wetterj/gin-sqlx-crud/internal/models"
)

var srv *Server

// These are integration tests. There should be a test DB running.
func TestServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var err error
	srv, err = NewServer()
	if err != nil {
		t.Fatalf("cannot create server: %v", err)
	}

	t.Run("Create", createPerson)
	t.Run("Update", updatePerson)
	t.Run("Get", getPerson)
	t.Run("Delete", deletePerson)
}

const (
	testFirstName = "john"
	testLastName  = "doe"
	testAddress   = "58 downing st"
	testAge       = 22
)

var testPersonID string

func createPerson(t *testing.T) {
	var body []byte
	var req *http.Request
	var w *httptest.ResponseRecorder
	var person models.Person
	var err error

	body, _ = json.Marshal(gin.H{
		"firstName": testFirstName,
	})

	req, _ = http.NewRequest("POST", "/person", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	srv.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatal("cannot create person")
	}

	err = json.Unmarshal(w.Body.Bytes(), &person)
	if err != nil {
		t.Fatal("cannot unmarshal person")
	}
	if person.FirstName != testFirstName {
		t.Fatal("name not set correctly")
	}
	if person.LastName.Valid {
		t.Fatal("name not set correctly")
	}

	testPersonID = person.ID
}

func updatePerson(t *testing.T) {
	var body []byte
	var req *http.Request
	var w *httptest.ResponseRecorder

	body, _ = json.Marshal(gin.H{
		"firstName": testFirstName,
		"lastName":  testLastName,
		"address":   testAddress,
		"age":       testAge,
	})

	// Put a missing person
	req, _ = http.NewRequest("PUT", "/person/fdsadfa", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	srv.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Log(w)
		t.Fatal("updated a missing person")
	}

	// Put the real thing
	req, _ = http.NewRequest("PUT", "/person/"+testPersonID, bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	srv.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("cannot update person")
	}
}

func getPerson(t *testing.T) {
	var req *http.Request
	var w *httptest.ResponseRecorder
	var person models.Person
	var err error

	// Get a missing person
	req, _ = http.NewRequest("GET", "/person/fdsadfa", nil)
	w = httptest.NewRecorder()
	srv.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("got a missing person")
	}

	// Get the real thing
	req, _ = http.NewRequest("GET", "/person/"+testPersonID, nil)
	w = httptest.NewRecorder()
	srv.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("cannot get person")
	}

	err = json.Unmarshal(w.Body.Bytes(), &person)
	if err != nil {
		t.Fatal("cannot unmarshal person")
	}
	if person.FirstName != testFirstName {
		t.Fatal("name not set correctly")
	}
	if person.LastName != null.StringFrom(testLastName) {
		t.Fatal("name not set correctly")
	}
	if person.Address != null.StringFrom(testAddress) {
		t.Fatal("address not set correctly")
	}
	if person.Age != null.IntFrom(int64(testAge)) {
		t.Fatal("age not set correctly")
	}
}

func deletePerson(t *testing.T) {
	var req *http.Request
	var w *httptest.ResponseRecorder

	// Get a missing person
	req, _ = http.NewRequest("DELETE", "/person/"+testPersonID, nil)
	w = httptest.NewRecorder()
	srv.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("cannot delete a person")
	}
}
