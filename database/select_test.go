package database

import (
	"context"
	"testing"

	"github.com/spf13/cast"
	_ "modernc.org/sqlite"
)

func TestSelectToMapAny(t *testing.T) {
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
	_, err = SelectToMapAny(Context(context.Background(), nil), "SELECT * FROM users")
	if err == nil {
		t.Error("Expected error for nil querier")
	} else if err.Error() != "querier (db/tx/conn) is nil" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test successful query
	result, err := SelectToMapAny(Context(context.Background(), db), "SELECT * FROM users ORDER BY id ASC")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}

	row := result[1]

	if cast.ToInt(row["id"]) != 2 {
		t.Errorf("Expected id 2, got %v", row["id"])
	}

	if row["name"] != "Bob" {
		t.Errorf("Expected name 'Bob', got '%v'", row["name"])
	}

	if row["email"] != "bob@example.com" {
		t.Errorf("Expected email 'bob@example.com', got '%v'", row["email"])
	}

	// Test query with no results
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to delete data: %v", err)
	}
	result, err = SelectToMapAny(Context(context.Background(), db), "SELECT * FROM users")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 rows, got %d", len(result))
	}

	// Test query with error
	_, err = SelectToMapAny(Context(context.Background(), db), "INVALID SQL")
	if err == nil {
		t.Error("Expected error for invalid SQL")
	}
}

func TestSelectToMapString(t *testing.T) {
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
	_, err = SelectToMapString(Context(context.Background(), nil), "SELECT * FROM users")
	if err == nil {
		t.Error("Expected error for nil querier")
	} else if err.Error() != "querier (db/tx/conn) is nil" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test successful query
	result, err := SelectToMapString(Context(context.Background(), db), "SELECT * FROM users ORDER BY id ASC")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}

	row := result[1]

	if row["id"] != "2" {
		t.Errorf("Expected id 2, got %v", row["id"])
	}

	if row["name"] != "Bob" {
		t.Errorf("Expected name 'Bob', got '%v'", row["name"])
	}

	if row["email"] != "bob@example.com" {
		t.Errorf("Expected email 'bob@example.com', got '%v'", row["email"])
	}

	// Test query with no results
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to delete data: %v", err)
	}
	result, err = SelectToMapString(Context(context.Background(), db), "SELECT * FROM users")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 rows, got %d", len(result))
	}

	// Test query with error
	_, err = SelectToMapString(Context(context.Background(), db), "INVALID SQL")
	if err == nil {
		t.Error("Expected error for invalid SQL")
	}
}

func createUserTableAndInserTesttData(db QueryableInterface) error {
	// Create a test table
	_, err := db.ExecContext(context.Background(), "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		return err
	}

	sql := "INSERT INTO users (name, email) VALUES ('Alice', 'alice@example.com'), ('Bob', 'bob@example.com'), ('Charlie', 'charlie@example.com')"

	// Insert some test data
	_, err = db.ExecContext(context.Background(), sql)

	if err != nil {
		return err
	}

	return nil
}
