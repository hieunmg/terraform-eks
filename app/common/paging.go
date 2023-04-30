package common

// Paging defined struct for pagination
type Paging struct {
	PagingQuery `json:",inline"`
	Total       int64 `json:"total"`
} //@name Paging

type PagingQuery struct {
	// Default: 1
	Page int `json:"page" form:"page" binding:"omitempty,min=1"`
	// Default: 50
	Limit int `json:"limit" form:"limit" binding:"omitempty,min=1"`
} //@name PagingQuery

// Fulfill check page and limit when over range
func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}
