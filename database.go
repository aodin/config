package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Database contains the fields needed to connect to a database.
type Database struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
}

// Address returns a domain:port pair
func (db Database) Address() string {
	return fmt.Sprintf("%s:%d", db.Host, db.Port)
}

// Credentials with return the driver and credentials appropriate for Go's
// sql.Open function as strings
// TODO Return the full address instead of a connection string?
func (db Database) Credentials() (string, string) {
	// Only add the key if there is a value
	var values []string
	if db.Host != "" {
		values = append(values, fmt.Sprintf("host=%s", db.Host))
	}
	if db.Port != 0 {
		values = append(values, fmt.Sprintf("port=%d", db.Port))
	}
	if db.Name != "" {
		values = append(values, fmt.Sprintf("dbname=%s", db.Name))
	}
	if db.User != "" {
		values = append(values, fmt.Sprintf("user=%s", db.User))
	}
	if db.Password != "" {
		values = append(values, fmt.Sprintf("password=%s", db.Password))
	}
	if db.SSLMode != "" {
		values = append(values, fmt.Sprintf("sslmode=%s", db.SSLMode))
	}
	return db.Driver, strings.Join(values, " ")
}

// URL returns the db URL as a net.URL
func (db Database) URL() (u *url.URL) {
	u = &url.URL{
		Scheme: db.Driver,
		User:   url.UserPassword(db.User, db.Password),
		Host:   db.Address(),
		Path:   db.Name,
	}
	if db.SSLMode != "" {
		u.RawQuery = (url.Values{"sslmode": []string{db.SSLMode}}).Encode()
	}
	return u
}

// FullAddress returns the String output of the db URL
func (db Database) FullAddress() string {
	return db.URL().String()
}

// ParseDatabasePath will create a Database using the given filepath.
func ParseDatabasePath(path string) (c Database, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&c)
	return
}

// ParseDatabaseURL will create a Database from the given raw URL
func ParseDatabaseURL(rawurl string) (db Database, err error) {
	var parsed *url.URL
	parsed, err = url.Parse(rawurl)
	if err != nil {
		return
	}

	// Remove the forward slashes
	db.Name = strings.Trim(parsed.Path, "/")

	db.Driver = parsed.Scheme
	if parsed.User != nil {
		db.User = parsed.User.Username()
		db.Password, _ = parsed.User.Password()
	}

	// Split the port from the host
	parts := strings.SplitN(parsed.Host, ":", 2)
	if len(parts) < 2 {
		db.Host = parsed.Host
	} else {
		db.Host = parts[0]
		if db.Port, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return
		}
	}
	db.SSLMode = parsed.Query().Get("sslmode")
	return
}
