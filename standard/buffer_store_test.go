package standard_test

import (
	"testing"
)

func TestBB(t *testing.T) {
	bb := make([]byte, 2)
	t.Logf("%p:%p:%p", bb, &bb, &bb[0])
}
