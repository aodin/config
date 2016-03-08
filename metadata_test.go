package config

import "testing"

func TestMetadata(t *testing.T) {
	data := Metadata{"version": "1.0.0"}

	if len(data.Keys()) != 1 {
		t.Fatalf("Unexpected length of metadata keys")
	}

	if data.Keys()[0] != "version" {
		t.Errorf("The only metadata key should be 'version'")
	}

	if len(data.Values()) != 1 {
		t.Fatalf("Unexpected length of metadata keys")
	}

	if data.Values()[0] != "1.0.0" {
		t.Errorf("The only metadata value should be '1.0.0'")
	}

	if !data.Has("version") {
		t.Errorf("The data should contain the key 'version'")
	}

	if data.Get("version") != "1.0.0" {
		t.Errorf(
			"Unexpected value of 'version': %s != 1.0.0", data.Get("version"),
		)
	}
}
