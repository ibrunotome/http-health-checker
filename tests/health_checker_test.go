package tests

import (
	"testing"
)

func TestValidHealthCheck(t *testing.T) {
	test.Get("/").
		SetHeader("Accept", "application/json").
		SetQuery("link", "https://google.com").
		Expect(t).
		Status(200).
		Done()
}

func TestInvalidHealthCheck(t *testing.T) {
	test.Get("/").
		SetHeader("Accept", "application/json").
		SetQuery("link", "https://github.com/invaliduser/invalidrepo").
		Expect(t).
		Status(404).
		Type("plain").
		Done()
}
