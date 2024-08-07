package file

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDriver(t *testing.T) {
	t.Run("valid options", func(t *testing.T) {
		var file = filepath.ToSlash(filepath.Join(os.TempDir(), "gopi-frame-test", "writer", "driver", "file", fmt.Sprintf("log.%d.txt", time.Now().Unix())))
		var options = map[string]any{
			"file": file,
		}
		w, err := Open(options)
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			_, err = w.Write([]byte("hello world"))
			if err != nil {
				assert.FailNow(t, err.Error())
			}
			err := w.Close()
			if err != nil {
				assert.FailNow(t, err.Error())
			}
			content, err := os.ReadFile(file)
			if err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.Equal(t, "hello world", string(content))
		}
	})

	t.Run("invalid options", func(t *testing.T) {
		var options = map[string]any{
			"file": nil,
		}
		w, err := Open(options)
		assert.Error(t, err)
		assert.Nil(t, w)
	})
}
