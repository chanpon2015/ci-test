package main

import "testing"

func TestOutput(t *testing.T) {
  if "test" != Output() {
    t.Fail("error")
  }
}
