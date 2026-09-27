package domain

import (
	"fmt"
	"strconv"
)

type ID int64

func ParseID(value string) (ID, error) {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil || parsed <= 0 {
		return 0, fmt.Errorf("invalid ID %q: expected a positive integer", value)
	}
	return ID(parsed), nil
}

func (id ID) Valid() bool { return id > 0 }
