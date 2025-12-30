package models

type ListDriversResponse struct {
	Drivers []Driver  `json:"drivers"`
	Meta    *Metadata `json:"meta,omitempty"`
}

type Metadata struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	TotalPages int `json:"totalPages"`
}
