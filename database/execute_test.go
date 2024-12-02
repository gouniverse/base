package database

import (
	"context"
	"testing"
)

func TestExecute(t *testing.T) {
	db, err := initSqliteDB()

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// Test nil querier error
	_, err = Execute(Context(context.Background(), nil), "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	if err == nil {
		t.Error("Expected error for nil querier")
	} else if err.Error() != "querier (db/tx/conn) is nil" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test successful execution
	_, err = Execute(Context(context.Background(), db), "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
