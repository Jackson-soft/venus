package standard_test

import (
	"testing"

	"github.com/Jackson-soft/venus/standard"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := standard.NewSet()
	assert.NotNil(t, s)

	aa := s.List()
	t.Log(aa)

	assert.True(t, s.Empty())

	s.Insert("e")
	assert.Equal(t, s.Size(), 1)
	assert.False(t, s.Empty())

	s.Insert("b")
	assert.Equal(t, s.Size(), 2)

	array := s.List()
	assert.Equal(t, s.Size(), len(array))
	t.Log(array)

	sarray := s.SortList()
	t.Log(sarray)

	s.Erase("b")
	assert.Equal(t, s.Size(), 1)

	s.Clear()
	assert.True(t, s.Empty())
}
