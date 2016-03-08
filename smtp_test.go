package config

import "testing"

func TestSMTP(t *testing.T) {
	test := SMTP{From: "test", Host: "l", Port: 1234}

	if test.Address() != "l:1234" {
		t.Errorf("Unexpected address: %s != l:1234", test.Address())
	}

	if test.FromAddress() != "<test>" {
		t.Errorf("Unexpected from address: %s != <test>", test.FromAddress())
	}

	test.Alias = "alias"
	if test.FromAddress() != `"alias" <test>` {
		t.Errorf(
			"Unexpected from address with alias: %s != \"alias\" <test>",
			test.FromAddress(),
		)
	}
}
