package config

import (
	"net/url"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	conf, err := ParsePath("./testdata/example.json")
	if err != nil {
		t.Fatalf("ParsePath should not error")
	}
	if conf.Address() != "localhost:9001" {
		t.Errorf("Unexpected address: %s != localhost:9001", conf.Address())
	}
	if conf.URL().String() != "http://localhost:9001" {
		t.Errorf(
			"Unexpected URL: %s != http://localhost:9001", conf.Address())
	}
	if conf.URL().String() != conf.FullAddress() {
		t.Errorf("conf URL should equal the full address", conf.FullAddress())
	}
	if conf.Metadata.Get("version") != "1.0.0" {
		t.Errorf(
			"Unexpected metadata 'version': %s != 1.0.0",
			conf.Metadata.Get("version"),
		)
	}

	conf.ProxyDomain = "example.com"
	conf.ProxyPort = 3000

	if conf.FullAddress() != "http://example.com:3000" {
		t.Errorf(
			"Unexpected full address: %s != http://example.com:3000",
			conf.FullAddress(),
		)
	}

	expect := &url.URL{Scheme: "http", Host: "example.com:3000"}
	if conf.URL().String() != expect.String() {
		t.Errorf("Unexpected URL: %s != %s", conf.URL(), expect)
	}

	conf.ProxyPort = 80
	conf.HTTPS = true

	expect = &url.URL{Scheme: "https", Host: "example.com"}
	if conf.FullAddress() != "https://example.com" {
		t.Errorf(
			"Unexpected full address with proxy: %s != https://example.com",
			conf.FullAddress(),
		)
	}
	if conf.URL().String() != expect.String() {
		t.Errorf("Unexpected URL with proxy: %s != %s", conf.URL(), expect)
	}

	// StaticURL test
	conf.StaticURL = "static/"
	if conf.StaticAddress() != "https://example.com/static/" {
		t.Errorf(
			"unexpected static address: %s != https://example.com/static/",
			conf.StaticAddress(),
		)
	}

	conf.StaticURL = "/static/"
	if conf.StaticAddress() != "https://example.com/static/" {
		t.Errorf(
			"unexpected static address: %s != https://example.com/static/",
			conf.StaticAddress(),
		)
	}

	// Test Database
	db := "host=localhost port=5432 dbname=db user=pg password=pass"
	if _, dbconf := conf.Database.Credentials(); dbconf != db {
		t.Errorf(
			"Unexpected DB credentials: %s != %s",
			dbconf, db,
		)
	}

	// Test Cookie
	if conf.Cookie.Age != 336*time.Hour {
		t.Errorf(
			"Unexpected cookie age: %d != %d", conf.Cookie.Age, 336*time.Hour,
		)
	}
	if conf.Cookie.Domain != "" {
		t.Errorf("Cookie domain should be empty")
	}
	if conf.Cookie.HttpOnly {
		t.Errorf("Cookie should be HTTP Only = false")
	}
	if conf.Cookie.Name != "sessionid" {
		t.Errorf("Unexpected cookie name: %s != sessionid", conf.Cookie.Name)
	}
	if conf.Cookie.Path != "/" {
		t.Errorf("Unexpected cookie path: %s != /", conf.Cookie.Path)
	}
	if conf.Cookie.Secure {
		t.Errorf("Cookie should be Secure = false")
	}
}
