package database

import (
	"database/sql"
	"reflect"
	"strings"
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
func DatabaseType(db *sql.DB) string {
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
