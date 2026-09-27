package application

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func requireIDs(ids ...domain.ID) error {
	for _, id := range ids {
		if !id.Valid() {
			return errors.New("IDs must be positive integers")
		}
	}
	return nil
}

func requireObject(payload []byte) error {
	if len(payload) == 0 || !json.Valid(payload) || !bytes.HasPrefix(bytes.TrimSpace(payload), []byte("{")) {
		return errors.New("request body must be a JSON object")
	}
	return nil
}
