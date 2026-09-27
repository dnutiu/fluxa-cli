package fluxa

import (
	"crypto/rand"
	"encoding/hex"
)

func NewIdempotencyKey() (string, error) {
	var random [16]byte
	if _, err := rand.Read(random[:]); err != nil {
		return "", err
	}
	return "fluxa-cli-" + hex.EncodeToString(random[:]), nil
}
