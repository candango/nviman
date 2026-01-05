package filehash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// SHA256 computes the SHA256 hash of the file at the given path and returns it
// as a hexadecimal string prefixed with "sha256:".
func SHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil)), nil
}
