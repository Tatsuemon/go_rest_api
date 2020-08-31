package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("root", "", "go_rest_api_db")
	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}
 const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	age INT NOT NULL
)
 `

 // "/users" test
 func TestEmptyTable(t *testing.T) {
	 clearTable()

	 req, _ := http.NewRequest("GET", "/users", nil)
	 response := executeRequest(req)

	 checkResponseCode(t, http.StatusOK, response.Code)

	 if body := response.Body.String(); body != "[]" {
		 t.Errorf("Expected an empty array. Got %s", body)
	 }
 }

 func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	 rr := httptest.NewRecorder()
	 a.Router.ServeHTTP(rr, req)

	 return rr
 }

 func checkResponseCode(t *testing.T, expected, actual int) {
	 if expected != actual {
		 t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	 }
 }

 // userが存在しない時のテスト
 func TestGenNonExistentUser(t *testing.T) {
	 clearTable()

	 req, _ := http.NewRequest("GET", "/users/45", nil)
	 response := executeRequest(req)

	 checkResponseCode(t, http.StatusNotFound, response.Code)

	 var m map[string]string
	 json.Unmarshal(response.Body.Bytes(), &m)
	 if m["error"] != "User not found" {
		 t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	 } 
 }
 
