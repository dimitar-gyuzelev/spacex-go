package v4

type Query struct {
	Query   interface{} `json:"query"`
	Options interface{} `json:"options"`
}

type RespQuery[T interface{}] struct {
	Docs          []T  `json:"docs"`
	TotalDocs     int  `json:"totalDocs"`
	Offset        int  `json:"offset"`
	Limit         int  `json:"limit"`
	TotalPages    int  `json:"totalPages"`
	Page          int  `json:"page"`
	PagingCounter int  `json:"pagingCounter"`
	HasPrevPage   bool `json:"hasPrevPage"`
	HasNextPage   bool `json:"hasNextPage"`
	PrevPage      int  `json:"prevPage"`
	NextPage      int  `json:"nextPage"`
}