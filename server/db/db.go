package db

import (
	"database/sql"
	"net/url"

	// postgres driver
	_ "github.com/lib/pq"
)

// Connection structure handles data needed to connect to database
type Connection struct {
	DbName         string
	User, Password string
	Host           string
	DisableSSL     bool
}

// ConnectionURL forms an database URL from given data
func (c *Connection) ConnectionURL() string {
	dbURL := &url.URL{
		Scheme: "postgres",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbURL.RawQuery = url.Values{
			"sslmode": []string{"disable"},
		}.Encode()
	}
	return dbURL.String()
}

// Open connects to database
func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.ConnectionURL())
}
