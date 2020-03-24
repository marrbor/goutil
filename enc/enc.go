package enc

import (
	"crypto/sha256"
	"fmt"
	"hash/fnv"
)

// Hash32 converts string to hash value.
func Hash32(s string) (uint32, error) {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

// Encrypt256Password returns SHA256 encrypted string.
func Encrypt256Password(src string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(src)))
}
