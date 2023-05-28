package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDB(t *testing.T) {

	t.Run("create new mock db", func(t *testing.T) {
		_, _, err := sqlmock.New()
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)

		}
	})

	t.Run("create new db with valid path", func(t *testing.T) {

		db := NewDB()
		err := db.Connect(t.TempDir() + "test.db")
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)

		}
	})

	// t.Run("create new db with invalid path", func(t *testing.T) {

	// 	db := NewDB()
	// 	err := db.Connect(t.TempDir() + "/DD/test.db")
	// 	if err == nil {
	// 		t.Errorf("an error was expected, instead got '%s'", err)

	// 	}
	// })

}
