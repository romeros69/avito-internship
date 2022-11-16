package models

type Pagination struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Size
}
