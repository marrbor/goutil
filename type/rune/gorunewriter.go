package gorunewriter

import (
	"io"
	"unicode/utf8"
)

// RuneWriter is a writer that replaces characters that cannot be converted with '?'.
// https://teratail.com/questions/106106
type RuneWriter struct {
	Writer io.Writer
}

// Write writes string and return output length or error.
func (rw *RuneWriter) Write(b []byte) (int, error) {
	var err error
	l := 0

loop:
	for len(b) > 0 {
		_, n := utf8.DecodeRune(b)
		if n == 0 {
			break loop
		}
		_, err = rw.Writer.Write(b[:n])
		if err != nil {
			_, err = rw.Writer.Write([]byte{'?'})
			if err != nil {
				break loop
			}
		}
		l += n
		b = b[n:]
	}
	return l, err
}
