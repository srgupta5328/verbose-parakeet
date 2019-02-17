package helpers

import "testing"

func TestGenerateUUID(t *testing.T) {
	t.Run("Generating a UUID", func(t *testing.T) {
		got := generateUUID()
		if got != "" {
			t.Logf("Generated a uuid successfully: %s", got)
		}
	})
}
