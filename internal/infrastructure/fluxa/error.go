package fluxa

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Status  int
	Code    string
	Message string
	Details json.RawMessage
}

func (e *Error) Error() string {
	if e.Code == "" {
		return fmt.Sprintf("Fluxa API returned HTTP %d", e.Status)
	}
	return fmt.Sprintf("Fluxa API %s (HTTP %d): %s", e.Code, e.Status, e.Message)
}
