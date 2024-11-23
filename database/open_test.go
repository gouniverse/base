package database

import (
	"strings"
	"testing"

	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func TestOpenWithUnsupportedDriver(t *testing.T) {
	db, err := Open(Options().
		SetDatabaseType("unsupported_driver").
		SetDatabaseHost("").
		SetDatabasePort("").
		SetDatabaseName(":memory:").
		SetUserName("").
		SetPassword(""))

	if err == nil {
		t.Fatal(`err MUST NOT be nil`)
	}

	if !strings.Contains(err.Error(), `driver unsupported_driver is not supported`) {
		t.Fatal(`err MUST contain 'unsupported_driver unsupported is not supported', found: `, err.Error())
	}

	if db != nil {
		t.Fatal(`db MUST be nil`)
	}
}

func TestOpen(t *testing.T) {
	db, err := Open(Options().
		SetDatabaseType(DATABASE_TYPE_SQLITE).
		SetDatabaseHost("").
		SetDatabasePort("").
		SetDatabaseName(":memory:").
		SetUserName("").
		SetPassword(""))

	if err != nil {
		t.Fatal(err)
	}

	if db == nil {
		t.Fatal(`db is nil`)
	}
}

func TestOpenDatabseIsRequired(t *testing.T) {
	db, err := Open(Options().
		SetDatabaseType(DATABASE_TYPE_SQLITE).
		SetDatabaseHost("").
		SetDatabasePort("").
		SetUserName("").
		SetPassword(""))

	if err == nil {
		t.Fatal(`err MUST NOT be nil`)
	}

	if !strings.Contains(err.Error(), `database name is required`) {
		t.Fatal(`err MUST contain, 'database name is required'`, `, found: `, err.Error())
	}

	if db != nil {
		t.Fatal(`db MUST be nil`)
	}
}

func TestOpenHostIsRequired(t *testing.T) {
	db, err := Open(Options().
		SetDatabaseType(DATABASE_TYPE_MYSQL).
		SetDatabaseHost("").
		SetDatabasePort("").
		SetDatabaseName(":memory:").
		SetUserName("").
		SetPassword(""))

	if err == nil {
		t.Fatal(`err MUST NOT be nil`)
	}

	if !strings.Contains(err.Error(), `database host is required`) {
		t.Fatal(`err MUST contain, 'database host is required'`, `, found: `, err.Error())
	}

	if db != nil {
		t.Fatal(`db MUST be nil`)
	}
}

func TestOpenPortIsRequired(t *testing.T) {
	db, err := Open(Options().
		SetDatabaseType(DATABASE_TYPE_MYSQL).
		SetDatabaseHost("localhost").
		SetDatabasePort("").
		SetDatabaseName(":memory:").
		SetUserName("").
		SetPassword(""))

	if err == nil {
		t.Fatal(`err MUST NOT be nil`)
	}

	if !strings.Contains(err.Error(), `database port is required`) {
		t.Fatal(`err MUST contain, 'database port is required'`, `, found: `, err.Error())
	}

	if db != nil {
		t.Fatal(`db MUST be nil`)
	}
}
