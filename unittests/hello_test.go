package main

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	expected := "hello"
	actual := Hello()

	if !strings.EqualFold(expected, actual) {
		t.Fail()
	}
}

func TestHelloError(t *testing.T) {
	expected := "error"
	actual := Hello()

	if !strings.EqualFold(expected, actual) {
		t.Fail()
	}
}
