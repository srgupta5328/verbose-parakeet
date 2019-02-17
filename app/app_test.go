package app

import "testing"

func TestInitialize(t *testing.T) {
	testRouter := Initialize()
	if testRouter == nil {
		t.Errorf("Error: '%v', Router didn't initialize", testRouter)
	}
}

func TestInitDB(t *testing.T) {
	t.Run("No Error starting the db", func(t *testing.T) {
		_, err := InitDB()
		if err != nil {
			t.Errorf("Error initializing the db")
		}
	})

}
