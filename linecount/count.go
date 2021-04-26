package linecount

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

var ErrCountingBinary = errors.New("Counting Binary")

func isText(data []byte) bool {
	contentType := http.DetectContentType(data)
	return strings.Contains(contentType, "text")
}

func Count(file io.Reader) (int, error) {
	data := make([]byte, 512)
	reader := bufio.NewReader(file)
	sep := []byte("\n")
	lines := 0

	n, e := reader.Read(data)

	// empty file
	if e == io.EOF {
		return lines, nil
	}
	if e != nil {
		return 0, e
	}

	// is binary
	if !isText(data[:n]) {
		return 0, ErrCountingBinary
	}

	for {
		if e == io.EOF {
			return lines, nil
		}

		if e != nil {
			return lines, e
		}

		lines += bytes.Count(data[:n], sep)

		n, e = reader.Read(data)
	}
}
