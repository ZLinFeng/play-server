package response

type PageResponse struct {
	Total    uint64      `json:"total"`
	Page     uint        `json:"page"`
	PageSize uint        `json:"pageSize"`
	Data     interface{} `json:"data"`
}
