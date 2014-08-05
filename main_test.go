package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {

	correct_result := "Hello World"
	str := HelloWorld()

	if str != correct_result {
		t.Error("Hello World Error")
	}
}
