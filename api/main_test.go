package main

import (
	"errors"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_MAIN", func(t *testing.T) {
		go main()
		time.Sleep(1 * time.Second)
		wg.Done()
	})

	t.Run("NEGATIVE_MAIN", func(t *testing.T) {
		errorLoadENV = errors.New("test")

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		main()
	})
}
