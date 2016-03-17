package config

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

type mockWriter struct {
	status int
	header http.Header
}

// Header
func (mw *mockWriter) Header() http.Header {
	return mw.header
}

// Write will write nothing
func (mw *mockWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// WriteHeader does nothing with the code
func (mw *mockWriter) WriteHeader(code int) {
	mw.status = code
}

func newMockWriter() *mockWriter {
	return &mockWriter{header: http.Header{}}
}

func TestCookie_Set(t *testing.T) {
	jan2014 := time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)

	// Mock a writer
	mw := newMockWriter()

	// Write the default cookie config and session
	DefaultCookie.Set(mw, "KEY", jan2014)

	value, ok := mw.header["Set-Cookie"]
	if !ok {
		t.Fatalf("auth: no Set-Cookie header was set on the writer")
	}
	expected := []string{
		"sessionid=KEY; Path=/; Expires=Wed, 01 Jan 2014 00:00:00 GMT",
	}
	if !reflect.DeepEqual(value, expected) {
		t.Errorf(
			"unexpected cookie value: %+v != %+v",
			value, expected,
		)
	}
}
