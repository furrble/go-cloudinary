package cloudinary

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
)

// Returns SHA1 file checksum
func fileChecksum(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	io.WriteString(hash, string(data))
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
