package request

type PageParams struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
