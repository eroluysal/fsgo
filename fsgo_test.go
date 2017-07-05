package main

import (
	"testing"
)

func TestFNameAndExt(t *testing.T) {
	n, e := fNameAndExt("foo.txt")
	if n != "foo" {
		t.Fatal("name must be `foo`!")
	}
	if e != ".txt" {
		t.Fatal("ext must be `.txt`!")
	}
}
