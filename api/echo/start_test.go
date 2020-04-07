package echo

import (
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	t.Run("NEGATIVE_START", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		os.Setenv("SERVER_PORT", "ABCD")
		Start()
	})
}
