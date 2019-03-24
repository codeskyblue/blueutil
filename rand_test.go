package blueutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringN(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := RandStringN(i)
		t.Log(s)
		assert.Equal(t, i, len(s))
	}
}
