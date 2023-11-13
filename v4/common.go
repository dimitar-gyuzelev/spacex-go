package v4

const VersionPrefix = "/v4"

type Query struct {
	Query   interface{} `json:"query"`
	Options interface{} `json:"options"`
}

type RespQuery[T any] struct {
	Docs          []T  `json:"docs"`
	TotalDocs     int  `json:"totalDocs"`
	Offset        int  `json:"offset"`
	Limit         int  `json:"limit"`
	TotalPages    int  `json:"totalPages"`
	Page          int  `json:"page"`
	PagingCounter int  `json:"pagingCounter"`
	PrevPage      int  `json:"prevPage"`
	NextPage      int  `json:"nextPage"`
	HasPrevPage   bool `json:"hasPrevPage"`
	HasNextPage   bool `json:"hasNextPage"`
}

type Mass struct {
	KG int `json:"kg"`
	LB int `json:"lb"`
}

type Volume struct {
	CubicMeters float32 `json:"cubic_meters"`
	CubicFeet   float32 `json:"cubic_feet"`
}

type Thrust struct {
	KN  float32 `json:"kN"`
	LBF float32 `json:"lbf"`
}

type Distance struct {
	Meters float32 `json:"meters"`
	Feet   float32 `json:"feet"`
}

type ForceUnits struct {
	KN  float32 `json:"kN"`
	LBF float32 `json:"lbf"`
}
