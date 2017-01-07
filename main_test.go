package main

import (
	"testing"
)

func TestPassing(t *testing.T) {
	//This test will pass
}

func TestFailing(t *testing.T) {
	//This test will fail
	t.Error("Test failed because: ", "reasons")
}
