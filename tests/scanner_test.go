package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-restful/app/model"
)

func TestScanStruct(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	db, err := sql.Open("mysql", os.Getenv("DATABASE_TEST_URL"))
	if err != nil {
		t.Error(err)
	}

	email := faker.Email()
	lastName := faker.LastName()
	_, err = db.Exec("INSERT INTO users(first_name, last_name, email, password, created_at,updated_at) VALUES (?, ?, ?, ?, now(), now())", faker.FirstName(), lastName, email, "password")

	if err != nil {
		t.Error(err)
	}

	rows, err := db.Query("SELECT last_name,email FROM users where email = ? limit 1", email)

	if err != nil {
		t.Error(err)
	}

	var userStruct model.User
	if rows.Next() {
		model.ScanRow(&userStruct, rows)
	}

	if userStruct.Email != email {
		t.Errorf("Expected %s, but got %s", email, userStruct.Email)
	}

	if userStruct.LastName.String != lastName {
		t.Errorf("Expected %s, but got %s", lastName, userStruct.LastName.String)
	}
}
func TestScanMap(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	db, err := sql.Open("mysql", os.Getenv("DATABASE_TEST_URL"))
	if err != nil {
		t.Error(err)
	}

	email := faker.Email()
	lastName := faker.LastName()
	_, err = db.Exec("INSERT INTO users(first_name, last_name, email, password, created_at,updated_at) VALUES (?, ?, ?, ?, now(), now())", faker.FirstName(), lastName, email, "password")

	if err != nil {
		t.Error(err)
	}

	rows, err := db.Query("SELECT id,last_name,email FROM users where email = ? limit 1", email)

	if err != nil {
		t.Error(err)
	}

	var userMap = make(map[string]interface{})
	if rows.Next() {
		model.ScanRow(&userMap, rows)
	}

	if e, _ := userMap["email"]; e.(string) != email {
		t.Errorf("Expected %s, but got %s", email, e)
	}
}