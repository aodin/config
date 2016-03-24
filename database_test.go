package config

import "testing"

func TestDatabase(t *testing.T) {
	db := Database{
		Driver:   "postgres",
		User:     "admin",
		Password: "secret",
		Host:     "localhost",
		Port:     5432,
		Name:     "what",
	}

	// Without SSLMode
	if db.Address() != "localhost:5432" {
		t.Errorf("Unexpected DB address: %s != localhost:5432", db.Address())
	}

	expected := "postgres://admin:secret@localhost:5432/what"
	if db.FullAddress() != expected {
		t.Errorf(
			"Unexpected DB full address: %s != %s", db.FullAddress(), expected,
		)
	}

	// With SSLMode
	db.SSLMode = "disable"
	expected = "postgres://admin:secret@localhost:5432/what?sslmode=disable"
	if db.FullAddress() != expected {
		t.Errorf(
			"Unexpected DB full address with SSLMode: %s != %s",
			db.FullAddress(), expected,
		)
	}

	// TODO without user/password?
}

func TestDatabase_ParseDatabaseURL(t *testing.T) {
	rawurl := "postgres://admin:secret@localhost:5432/what?sslmode=disable"
	db, err := ParseDatabaseURL(rawurl)
	if err != nil {
		t.Fatalf("failed to parse database URL: %s", err)
	}

	expected := Database{
		Driver:   "postgres",
		User:     "admin",
		Password: "secret",
		Host:     "localhost",
		Port:     5432,
		Name:     "what",
		SSLMode:  "disable",
	}
	if expected != db {
		t.Errorf("unexpected database config: %s != %s", expected, db)
	}
}
