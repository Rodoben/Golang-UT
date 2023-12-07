package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_app_ConnectToDatabase(t *testing.T) {
	_, err := app.connectToDatabase()
	if err != nil {
		t.Error("test failed to connect to databse")
	}
}

func Test_app_TableOpenDatabase(t *testing.T) {
	invalidDsn := fmt.Sprintf("host=%s port=%v user=%s passwod=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
	e := errors.New("failed to connect to `host=localhost user=postgres database=users`: failed SASL auth (FATAL: password authentication failed for user \"postgres\" (SQLSTATE 28P01))")

	testDB := []struct {
		name      string
		dsn       string
		errorCode error
	}{
		{name: "valid credentials", dsn: app.DSN, errorCode: nil},
		{name: "invalid credentials", dsn: invalidDsn, errorCode: e},
	}
	for _, test := range testDB {
		_, err := openDB(test.dsn)
		// Check if the error code matches the expected result
		if err != nil && err.Error() != test.errorCode.Error() {
			t.Errorf("Test case '%s' failed: expected error '%v', got '%v'", test.name, test.errorCode, err)
		}
	}

}
