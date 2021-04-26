package linecount

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

// bash command `wc -l` count "this\nis\na\ntest" as 3 lines
func TestCountFeatureWithoutNewlineEnd(T *testing.T) {
	data := []byte("this\nis\na\ntest")
	result, e := Count(bytes.NewReader(data))
	assert.Nil(T, e)
	assert.Equal(T, 3, result)
}

func TestCountWithNewlineEnd(T *testing.T) {
	data := []byte("this\nis\na\ntest\n")
	result, e := Count(bytes.NewReader(data))
	assert.Nil(T, e)
	assert.Equal(T, 4, result)
}

//  bash command `wc -l` count empty files as 0 lines
func TestCountWithEmptyFile(T *testing.T) {
	data := []byte("")
	result, e := Count(bytes.NewReader(data))
	assert.Nil(T, e)
	assert.Equal(T, 0, result)
}

func TestCountWithBinaryFile(T *testing.T) {
	data := make([]byte, 1024)
	result, e := Count(bytes.NewReader(data))
	assert.NotNil(T, e)
	assert.Equal(T, ErrCountingBinary, e)
	assert.Equal(T, 0, result)
}
