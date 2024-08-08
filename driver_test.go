package writer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Run("duplicate driver name", func(t *testing.T) {
		assert.Panics(t, func() {
			Register("test", new(mockDriver))
		})
	})

	t.Run("nil driver", func(t *testing.T) {
		assert.Panics(t, func() {
			Register("mock", nil)
		})
	})

	t.Run("success", func(t *testing.T) {
		Register("mock", new(mockDriver))
		assert.True(t, drivers.ContainsKey("mock"))
	})
}

func TestDrivers(t *testing.T) {
	assert.ElementsMatch(t, []string{"mock", "test"}, Drivers())
}

func TestOpen(t *testing.T) {
	t.Run("unregistered driver", func(t *testing.T) {
		w, err := Open("unregistered", nil)
		assert.Error(t, err)
		assert.Nil(t, w)
	})

	t.Run("registered driver", func(t *testing.T) {
		w, err := Open("test", map[string]any{
			"name": "test",
		})
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			assert.IsType(t, new(mockWriter), w)
		}
	})
}
