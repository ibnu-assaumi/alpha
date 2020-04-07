package errorutil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		e := NewError("test")
		assert.NotNil(t, e)
	})
}

func TestError_Error(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		e := NewError("test")
		assert.Equal(t, e.Error(), "map[test:[]]")
	})
}

func TestError_AppendMessage(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		e := NewError("test")
		e.AppendMessage("testing")
		assert.Equal(t, e.Message["test"], []string{"testing"})
	})
}

func TestIsSystemError(t *testing.T) {
	t.Run("CUSTOM_ERROR_TYPE", func(t *testing.T) {
		e := NewError("test")
		e.AppendMessage("testing")
		val, system := IsSystemError(e)
		assert.NotNil(t, val)
		assert.Equal(t, system, false)
	})

	t.Run("SYSTEM_ERROR_TYPE", func(t *testing.T) {
		err := errors.New("system error")
		val, system := IsSystemError(err)
		assert.Nil(t, val)
		assert.Equal(t, system, true)
	})
}

func TestError_GetErrorMessage(t *testing.T) {
	e := NewError("test")
	t.Run("SUCCESS", func(t *testing.T) {
		assert.Equal(t, e.GetErrorMessage(), []string{})
	})
}

func TestError_HasError(t *testing.T) {
	t.Run("HAS_ERROR", func(t *testing.T) {
		e := NewError("test")
		e.AppendMessage("has error")
		assert.Equal(t, e.HasError(), true)
	})

	t.Run("HAS_NO_ERROR", func(t *testing.T) {
		e := NewError("test")
		assert.Equal(t, e.HasError(), false)
	})
}
