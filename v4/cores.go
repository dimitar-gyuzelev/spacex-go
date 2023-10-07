package v4

const PathCores = "/cores"

// CoresAPI defines the functions to work with SpaceX Cores V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/cores/v4
type CoresAPI interface {
	GetCoresAll() ([]Core, error)

	GetCoreByID(id string) (Core, error)

	PostCoresQuery(query Query) (RespQuery[Core], error)
}

// Core represents the data returned by SpaceX Cores V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/cores/v4/schema.md#core-schema
type Core struct {
	ID           string   `json:"id"`
	Status       string   `json:"status,omitempty"`
	Serial       string   `json:"serial"`
	LastUpdate   string   `json:"last_update,omitempty"`
	Launches     []string `json:"launches,omitempty"`
	ReuseCount   int      `json:"reuse_count,omitempty"`
	Block        int      `json:"block,omitempty"`
	RTLSAttempts int      `json:"rtls_attempts,omitempty"`
	RTLSLandings int      `json:"rtls_landings,omitempty"`
	ASDSAttempts int      `json:"asds_attempts,omitempty"`
	ASDSLandings int      `json:"asds_landings,omitempty"`
}

func (c Core) String() string {
	return stringFn[Core](c)
}
