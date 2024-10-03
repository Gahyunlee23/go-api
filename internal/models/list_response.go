package models

// ListResponse is a generic struct for paginated list responses
type ListResponse[T any] struct {
	TotalCount int64 `json:"totalCount"`
	Items      []T   `json:"items"`
}

// NewListResponse creates a new ListResponse
func NewListResponse[T any](totalCount int64, items []T) ListResponse[T] {
	return ListResponse[T]{
		TotalCount: totalCount,
		Items:      items,
	}
}
