package app

import "testing"

func TestInitialize(t *testing.T) {
	testRouter := Initialize()
	if testRouter == nil {
		t.Errorf("Error: '%v', Router didn't initialize", testRouter)
	}
}
