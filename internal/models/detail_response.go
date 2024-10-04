package models

type DetailResponse[T any, U any] struct {
	Category T   `json:"category"`
	Items    []U `json:"items"`
}

func NewDetailResponse[T any, U any](category T, items []U) DetailResponse[T, U] {
	return DetailResponse[T, U]{
		Category: category,
		Items:    items,
	}
}
