package domain

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	TotalItems  int64 `json:"totalItems"`
	TotalPages  int   `json:"totalPages"`
	HasNext     bool  `json:"hasNext"`
	HasPrev     bool  `json:"hasPrev"`
}

type PaginationParams struct {
	Page  int
	Limit int
}

func NewPaginationParams(page, limit int) PaginationParams {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 { // máximo 100 por página
		limit = 10
	}
	return PaginationParams{Page: page, Limit: limit}
}

// Método para calcular offset (este es el que falta)
func (p PaginationParams) Offset() int {
	return (p.Page - 1) * p.Limit
}

// Método para verificar si hay paginación válida
func (p PaginationParams) IsValid() bool {
	return p.Page > 0 && p.Limit > 0
}

// Constructor para respuesta paginada (este es el que falta)
func NewPaginatedResponse(data interface{}, params PaginationParams, totalItems int) PaginatedResponse {
	totalPages := (totalItems + params.Limit - 1) / params.Limit
	if totalPages == 0 {
		totalPages = 1
	}

	return PaginatedResponse{
		Data: data,
		Pagination: Pagination{
			CurrentPage: params.Page,
			PageSize:    params.Limit,
			TotalItems:  int64(totalItems),
			TotalPages:  totalPages,
			HasNext:     params.Page < totalPages,
			HasPrev:     params.Page > 1,
		},
	}
}
