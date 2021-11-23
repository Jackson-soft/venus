package fsnotify_test

import (
	"path/filepath"
	"testing"

	"github.com/Jackson-soft/venus/fsnotify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotify(t *testing.T) {
	pp, err := filepath.Abs("./fsnotify.go")
	assert.NoError(t, err)
	t.Log(pp)
}

func TestEx(t *testing.T) {
	watcher_, err := fsnotify.New()
	require.NoError(t, err)

	file := "./fsnotify.go"
	err = watcher_.Mark(file, fsnotify.MarkAdd)
	require.NoError(t, err)

	go func() {
		var ev fsnotify.Event

		for {
			if ev, err = watcher_.Wait(); err == nil {
				if ev.Op_ == fsnotify.Write {
					// reload something
					t.Log("file writed....")
				}
			}
		}
	}()
}
