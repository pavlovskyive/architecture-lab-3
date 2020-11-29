package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "db1",
		User:       "vsevolodpavlovskyi",
		Password:   "admin",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://vsevolodpavlovskyi:admin@localhost/db1?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
