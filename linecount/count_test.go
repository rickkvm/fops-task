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

func TestCountWithNewlineEnd(T *testing.T) {
	data := []byte("this\nis\na\ntest\n")
	result := Count(data)
	assert.Equal(T, 4, result)
}

func TestCountWithEmptyFile(T *testing.T) {
	data := []byte("")
	result := Count(data)
	assert.Equal(T, 0, result)
}
