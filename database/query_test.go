package database

import (
	"context"
	"testing"
)

func TestQuery(t *testing.T) {
	db, err := initSqliteDB()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	err = createUserTableAndInserTesttData(db)
	if err != nil {
		t.Fatal(err)
	}

	// Test nil querier error
	_, err = Query(Context(context.Background(), nil), "SELECT * FROM users")

	if err == nil {
		t.Error("Expected error for nil querier")
	} else if err.Error() != "querier (db/tx/conn) is nil" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test successful query
	rows, err := Query(Context(context.Background(), db), "SELECT * FROM users")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// Iterate over rows and check data
	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			t.Fatalf("Failed to scan row: %v", err)
		}
		// Add assertions to check the values of id and name
	}

	// Test query with no results
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to delete data: %v", err)
	}
	rows, err = Query(Context(context.Background(), db), "SELECT * FROM users")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if rows.Next() {
		t.Error("Expected no rows, but got one")
	}

	// Test query with error (invalid SQL)
	_, err = Query(Context(context.Background(), db), "INVALID SQL")
	if err == nil {
		t.Error("Expected error for invalid SQL")
	}
}
