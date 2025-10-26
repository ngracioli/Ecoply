package utils

type PaginationWrapper[T any] struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
	HasPrev  bool `json:"has_prev"`
	Data     []T  `json:"data"`
}

func NewPaginationWrapper[T any](page int, pageSize int, data []T) *PaginationWrapper[T] {
	var hasNext bool = len(data)-1 == pageSize
	var hasPrev bool = page > 1
	var paginatedData []T = make([]T, 0)

	if len(data) > pageSize {
		paginatedData = data[:pageSize]
	} else {
		paginatedData = data
	}

	return &PaginationWrapper[T]{
		Page:     page,
		PageSize: pageSize,
		HasNext:  hasNext,
		HasPrev:  hasPrev,
		Data:     paginatedData,
	}
}
