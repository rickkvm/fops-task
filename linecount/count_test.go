package linecount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFeature(T *testing.T) {
	data := []byte("this\nis\na\ntest")
	result := Count(data)
	assert.Equal(T, 4, result)
}
