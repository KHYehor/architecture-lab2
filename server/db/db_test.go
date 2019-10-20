package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "tablets",
		User:       "ykh2k",
		Password:   "1111",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://ykh2k:1111@localhost/tablets?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
