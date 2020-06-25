package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

func choiceHashFunc(d *string) hash.Hash {
	switch *d {
	case "md5":
		return md5.New()

	case "sha1":
		return sha1.New()

	case "sha256":
		return sha256.New()

	case "sha512":
		return sha512.New()

	default:
		return sha256.New()

	}

}

func digest(bytes []byte, h hash.Hash) []byte {
	h.Write(bytes)
	return h.Sum(nil)
}
