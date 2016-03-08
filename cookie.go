package config

import (
	"encoding/json"
	"os"
	"time"
)

// Cookie contains the fields needed to set and retrieve cookies.
// Cookie names are valid tokens as defined by RFC 2616 section 2.2:
// http://tools.ietf.org/html/rfc2616#section-2.2
// TL;DR: Any non-control or non-separator character.
type Cookie struct {
	Age      time.Duration `json:"age"`
	Domain   string        `json:"domain"`
	HttpOnly bool          `json:"http_only"`
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Secure   bool          `json:"secure"`
}

// DefaultCookie is a default CookieConfig implementation. It expires after
// two weeks and is not very secure.
var DefaultCookie = Cookie{
	Age:      14 * 24 * time.Hour, // Two weeks
	Domain:   "",
	HttpOnly: false,
	Name:     "sessionid",
	Path:     "/",
	Secure:   false,
}

// ParseCookiePath will create a Cookie using the given filepath.
func ParseCookiePath(path string) (c Cookie, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&c)
	return
}
