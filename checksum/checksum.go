package checksum

import "hash"
import "bufio"
import "crypto/md5"
import "crypto/sha1"
import "crypto/sha256"
import "encoding/hex"
import "io"

type HashAlgorithm int
type HashSum hash.Hash

const (
	_ = iota
	MD5
	SHA1
	SHA256
)

func getHashReader(alg HashAlgorithm) HashSum {
	switch alg {
	case MD5:
		return md5.New()
	case SHA256:
		return sha256.New()
	case SHA1:
		return sha1.New()
	default:
		return md5.New()
	}
}

func ChecksumHash(file io.Reader, alg HashAlgorithm) (string, error) {
	hash := getHashReader(alg)

	buff := make([]byte, 2048)
	reader := bufio.NewReader(file)
	for {
		n, e := reader.Read(buff)
		if e == io.EOF {
			break
		}
		if e != nil {
			return "", e
		}
		hash.Write(buff[:n])
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
