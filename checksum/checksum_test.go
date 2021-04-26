package checksum

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleFile []byte = []byte("1\n")

func TestChecksumMD5(T *testing.T) {
	testFile := bytes.NewReader(sampleFile)
	result, e := ChecksumHash(testFile, MD5)
	assert.Nil(T, e)
	assert.Equal(T, "b026324c6904b2a9cb4b88d6d61c81d1", result)
}

func TestChecksumSHA256(T *testing.T) {
	testFile := bytes.NewReader(sampleFile)
	result, e := ChecksumHash(testFile, SHA256)
	assert.Nil(T, e)
	assert.Equal(T, "4355a46b19d348dc2f57c046f8ef63d4538ebb936000f3c9ee954a27460dd865", result)
}

func TestChecksumSHA1(T *testing.T) {
	testFile := bytes.NewReader(sampleFile)
	result, e := ChecksumHash(testFile, SHA1)
	assert.Nil(T, e)
	assert.Equal(T, "e5fa44f2b31c1fb553b6021e7360d07d5d91ff5e", result)
}
