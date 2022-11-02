package cb

import (
	"errors"
	"testing"

	"github.com/sony/gobreaker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{})

	t.Run("should return result as value", func(t *testing.T) {
		wantR := 123
		gotR, err := Execute(cb, func() (int, error) {
			return wantR, nil
		})
		require.NoError(t, err)
		assert.Equal(t, wantR, gotR)
	})

	t.Run("should return result as pointer", func(t *testing.T) {
		wantR := 123
		gotR, err := Execute(cb, func() (*int, error) {
			return &wantR, nil
		})
		require.NoError(t, err)
		assert.Equal(t, wantR, *gotR)
	})

	t.Run("should return result as pointer with nil value", func(t *testing.T) {
		gotR, err := Execute(cb, func() (*int, error) {
			return nil, nil
		})
		require.NoError(t, err)
		assert.Nil(t, gotR)
	})

	t.Run("should return error when f fails", func(t *testing.T) {
		wantErr := errors.New("some error")
		_, err := Execute(cb, func() (int, error) {
			return 0, wantErr
		})
		require.ErrorIs(t, err, wantErr)
	})
}
