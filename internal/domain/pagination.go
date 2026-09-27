package domain

import "errors"

type Pagination struct {
	Page    int
	PerPage int
}

func (p Pagination) Validate() error {
	if p.Page < 1 {
		return errors.New("page must be at least 1")
	}
	if p.PerPage < 1 || p.PerPage > 100 {
		return errors.New("per-page must be between 1 and 100")
	}
	return nil
}
