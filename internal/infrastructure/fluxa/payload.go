package fluxa

import "github.com/dnutiu/fluxa-cli/internal/domain"

func putField[T any](body map[string]any, key string, field domain.Field[T]) {
	if !field.Set {
		return
	}
	if field.Null {
		body[key] = nil
		return
	}
	body[key] = field.Value
}
