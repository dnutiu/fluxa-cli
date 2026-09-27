package application

type PageMeta struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Pages   int `json:"pages"`
}

type Collection[T any] struct {
	Data []T       `json:"data"`
	Meta *PageMeta `json:"meta,omitempty"`
}

type Record[T any] struct {
	Data T `json:"data"`
}
