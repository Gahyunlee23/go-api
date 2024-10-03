package models

type DetailResponse[T any] struct {
	Category T   `json:"category"`
	Items    []T `json:"items"`
}

func NewCategoryResponse[T any](category T, items []T) DetailResponse[T] {
	return DetailResponse[T]{
		Category: category,
		Items:    items,
	}
}
