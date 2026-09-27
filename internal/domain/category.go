package domain

type Category struct {
	ID             ID     `json:"id"`
	Name           string `json:"name"`
	Code           *int   `json:"code"`
	Kind           string `json:"kind"`
	CategoryPackID ID     `json:"category_pack_id"`
}
