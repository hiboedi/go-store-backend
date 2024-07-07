package models

type Pagination struct {
	Total       uint64
	CurrentPage uint32
	LastPage    uint32
	PerPage     uint32
}

type PaginationResponse struct {
	Total       uint64
	CurrentPage uint32
	LastPage    uint32
	PerPage     uint32
}

func ToPaginationResponse(pagination Pagination) PaginationResponse {
	return PaginationResponse{
		Total:       pagination.Total,
		CurrentPage: pagination.CurrentPage,
		LastPage:    pagination.LastPage,
		PerPage:     pagination.PerPage,
	}
}
