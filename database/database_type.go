package database

import (
	"database/sql"
	"reflect"
	"strings"
	"unsafe"
)

// DatabaseType finds the driver name from database
//
// It returns the type of the database in the following way:
//
//   - "mysql" for MySQL
//   - "postgres" for PostgreSQL
//   - "sqlite" for SQLite
//   - "mssql" for Microsoft SQL Server
//   - the full name of the driver otherwise
//
// The function is useful when you want to find the type of the database,
// without knowing it during compilation.
//
// Parameters:
// - db *sql.DB: the database connection
//
// Returns:
// - string: the type of the database
func DatabaseType(q Queryable) string {
	var db *sql.DB

	// check if q is sql.DB or sql.Tx or sql.Conn
	if qdb, ok := q.(*sql.DB); ok {
		db = qdb
	}

	if tx, ok := q.(*sql.Tx); ok {
		v := reflect.ValueOf(tx).Elem()
		dbField := v.FieldByName("db")
		dbFieldElem := reflect.NewAt(dbField.Type(), unsafe.Pointer(dbField.UnsafeAddr())).Elem()
		dbAny := dbFieldElem.Interface()
		db = dbAny.(*sql.DB)
	}

	if conn, ok := q.(*sql.Conn); ok {
		v := reflect.ValueOf(conn).Elem()
		dbField := v.FieldByName("db")
		dbFieldElem := reflect.NewAt(dbField.Type(), unsafe.Pointer(dbField.UnsafeAddr())).Elem()
		dbAny := dbFieldElem.Interface()
		db = dbAny.(*sql.DB)
	}

	driverFullName := reflect.ValueOf(db.Driver()).Type().String()

	if strings.Contains(driverFullName, DATABASE_TYPE_MYSQL) {
		return DATABASE_TYPE_MYSQL
	}

	if strings.Contains(driverFullName, DATABASE_TYPE_POSTGRES) || strings.Contains(driverFullName, "pq") {
		return DATABASE_TYPE_POSTGRES
	}

	if strings.Contains(driverFullName, DATABASE_TYPE_SQLITE) {
		return DATABASE_TYPE_SQLITE
	}

	if strings.Contains(driverFullName, DATABASE_TYPE_MSSQL) {
		return DATABASE_TYPE_MSSQL
	}

	return driverFullName
}
