package domain

// Field records whether a value was supplied, so edits can distinguish an
// omitted flag from a zero value or an explicit null.
type Field[T any] struct {
	Value T
	Set   bool
	Null  bool
}

func With[T any](value T) Field[T] { return Field[T]{Value: value, Set: true} }

func Cleared[T any]() Field[T] { return Field[T]{Set: true, Null: true} }
