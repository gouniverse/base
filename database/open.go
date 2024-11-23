package database

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

// Open opens the database
//
// Note:
//   - drivers are not included to this package to prevent size bloat
//   - you must add only the required database driver
//
// Drivers:
// - sqlite add the following includes:
// ```
// _ "modernc.org/sqlite"
// ```
// - mysql add the following includes:
// ```
// _ "github.com/go-sql-driver/mysql"
// ```
// - postgres add the following includes:
// ```
// _ "github.com/lib/pq"
// ```
//
// Business logic:
//   - opens the database based on the driver name
//   - each driver has its own set of parameters
//
// Parameters:
// - options openOptionsInterface
//
// Returns:
// - *sql.DB: the database connection
// - error: the error if any
func Open(options openOptionsInterface) (*sql.DB, error) {
	var db *sql.DB
	var err error

	err = options.Verify()

	if err != nil {
		return nil, err
	}

	databaseType := options.DatabaseType()
	host := options.DatabaseHost()
	port := options.DatabasePort()
	databaseName := options.DatabaseName()
	user := options.UserName()
	pass := options.Password()
	timezone := options.TimeZone()
	charset := options.Charset()

	dsn := dsn(databaseType, databaseName, user, pass, host, port, timezone, charset)

	db, err = sql.Open(databaseType, dsn)

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("database for driver " + databaseType + " could not be intialized")
	}

	if databaseType == DATABASE_TYPE_MYSQL || databaseType == DATABASE_TYPE_POSTGRES {
		// Maximum Idle Connections
		db.SetMaxIdleConns(5)
		// Maximum Open Connections
		db.SetMaxOpenConns(5)
		// Idle Connection Timeout
		db.SetConnMaxIdleTime(5 * time.Second)
		// Connection Lifetime
		db.SetConnMaxLifetime(30 * time.Second)
	}

	err = db.Ping()

	if err != nil {
		return nil, errors.Join(errors.New("database for driver "+databaseType+" could not be pinged"), err)
	}

	return db, nil
}

func dsn(
	driver string,
	databaseName string,
	user string,
	pass string,
	host string,
	port string,
	timezone string,
	charset string,
) string {
	if strings.EqualFold(driver, DATABASE_TYPE_SQLITE) {
		return databaseName
	}

	if strings.EqualFold(driver, DATABASE_TYPE_MYSQL) {
		dsn := user + `:` + pass
		dsn += `@tcp(` + host + `:` + port + `)/` + databaseName
		dsn += `?charset=` + charset
		dsn += `&parseTime=True`
		dsn += `&loc=` + timezone
		return dsn
	}

	if strings.EqualFold(driver, DATABASE_TYPE_POSTGRES) {
		dsn := `host=` + host
		dsn += ` user=` + user
		dsn += ` password=` + pass
		dsn += ` dbname=` + databaseName
		dsn += ` port=` + port
		dsn += ` sslmode=disable`
		dsn += ` TimeZone=` + timezone
		return dsn
	}

	return ""
}

func Options() openOptionsInterface {
	return &openOptions{
		properties: make(map[string]interface{}),
	}
}

type openOptions struct {
	properties map[string]interface{}
}

func (o *openOptions) Verify() error {
	if !o.HasDatabaseType() {
		return errors.New(`database type is required`)
	}

	if o.DatabaseType() == "" {
		return errors.New(`database type cannot be empty`)
	}

	supportedDrivers := []string{DATABASE_TYPE_SQLITE, DATABASE_TYPE_MYSQL, DATABASE_TYPE_POSTGRES}

	if !strings.EqualFold(o.DatabaseType(), DATABASE_TYPE_SQLITE) &&
		!strings.EqualFold(o.DatabaseType(), DATABASE_TYPE_MYSQL) &&
		!strings.EqualFold(o.DatabaseType(), DATABASE_TYPE_POSTGRES) {
		msg := `driver ` + o.DatabaseType() + ` is not supported.`
		msg += ` Supported drivers: ` + strings.Join(supportedDrivers, ", ")
		return errors.New(msg)
	}

	if !o.HasDatabaseName() {
		return errors.New(`database name is required`)
	}

	if o.DatabaseName() == "" {
		return errors.New(`database name cannot be empty`)
	}

	if !o.HasDatabaseHost() {
		o.SetDatabaseHost("")
	}

	if o.DatabaseHost() == "" && o.DatabaseType() != DATABASE_TYPE_SQLITE {
		return errors.New(`database host is required`)
	}

	if !o.HasDatabasePort() {
		o.SetDatabasePort("")
	}

	if o.DatabasePort() == "" && o.DatabaseType() != DATABASE_TYPE_SQLITE {
		return errors.New(`database port is required`)
	}

	if !o.HasUserName() {
		o.SetUserName("")
	}

	if !o.HasPassword() {
		o.SetPassword("")
	}

	if !o.HasTimeZone() {
		o.SetTimeZone("UTC")
	}

	if !o.HasCharset() {
		if o.DatabaseType() == DATABASE_TYPE_MYSQL {
			o.SetCharset("utf8mb4")
		} else {
			o.SetCharset("")
		}
	}

	return nil
}

func (o *openOptions) DatabaseType() string {
	return o.get("database_type").(string)
}

func (o *openOptions) HasDatabaseType() bool {
	return o.has("database_type")
}

func (o *openOptions) SetDatabaseType(databaseType string) openOptionsInterface {
	o.set("database_type", databaseType)
	return o
}

func (o *openOptions) DatabaseHost() string {
	return o.get("database_host").(string)
}

func (o *openOptions) HasDatabaseHost() bool {
	return o.has("database_host")
}

func (o *openOptions) SetDatabaseHost(databaseHost string) openOptionsInterface {
	o.set("database_host", databaseHost)
	return o
}

func (o *openOptions) DatabasePort() string {
	return o.get("database_port").(string)
}

func (o *openOptions) HasDatabasePort() bool {
	return o.has("database_port")
}

func (o *openOptions) SetDatabasePort(databasePort string) openOptionsInterface {
	o.set("database_port", databasePort)
	return o
}

func (o *openOptions) DatabaseName() string {
	return o.get("database_name").(string)
}

func (o *openOptions) HasDatabaseName() bool {
	return o.has("database_name")
}

func (o *openOptions) SetDatabaseName(databaseName string) openOptionsInterface {
	o.set("database_name", databaseName)
	return o
}

func (o *openOptions) UserName() string {
	return o.get("user_name").(string)
}

func (o *openOptions) HasUserName() bool {
	return o.has("user_name")
}

func (o *openOptions) SetUserName(userName string) openOptionsInterface {
	o.set("user_name", userName)
	return o
}

func (o *openOptions) Password() string {
	return o.get("password").(string)
}

func (o *openOptions) HasPassword() bool {
	return o.has("password")
}

func (o *openOptions) SetPassword(password string) openOptionsInterface {
	o.set("password", password)
	return o
}

func (o *openOptions) Charset() string {
	return o.get("charset").(string)
}

func (o *openOptions) HasCharset() bool {
	return o.has("charset")
}

func (o *openOptions) SetCharset(charset string) openOptionsInterface {
	o.set("charset", charset)
	return o
}

func (o *openOptions) TimeZone() string {
	return o.get("time_zone").(string)
}

func (o *openOptions) HasTimeZone() bool {
	return o.has("time_zone")
}

func (o *openOptions) SetTimeZone(timeZone string) openOptionsInterface {
	o.set("time_zone", timeZone)
	return o
}

func (o *openOptions) has(key string) bool {
	_, ok := o.properties[key]
	return ok
}

func (o *openOptions) set(key string, value interface{}) {
	o.properties[key] = value
}

func (o *openOptions) get(key string) interface{} {
	return o.properties[key]
}

// func (o *openOptions) getOrDefault(key string, defaultValue interface{}) interface{} {
// 	if o.has(key) {
// 		return o.get(key)
// 	}
// 	return defaultValue
// }

type openOptionsInterface interface {
	DatabaseType() string
	HasDatabaseType() bool
	SetDatabaseType(string) openOptionsInterface
	DatabaseHost() string
	HasDatabaseHost() bool
	SetDatabaseHost(string) openOptionsInterface
	DatabasePort() string
	HasDatabasePort() bool
	SetDatabasePort(string) openOptionsInterface
	DatabaseName() string
	HasDatabaseName() bool
	SetDatabaseName(string) openOptionsInterface
	UserName() string
	HasUserName() bool
	SetUserName(string) openOptionsInterface
	Password() string
	HasPassword() bool
	SetPassword(string) openOptionsInterface
	Charset() string
	HasCharset() bool
	SetCharset(string) openOptionsInterface
	TimeZone() string
	HasTimeZone() bool
	SetTimeZone(string) openOptionsInterface
	Verify() error
}
