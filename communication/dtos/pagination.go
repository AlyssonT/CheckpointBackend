package communication

type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
	Data       any `json:"data"`
}

type PaginationRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
