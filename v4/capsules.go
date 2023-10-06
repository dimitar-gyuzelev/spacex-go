package v4

const PathCapsules = "/capsules"

// CapsulesAPI defines functions to work with SpaceX Capsules V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/capsules/v4
type CapsulesAPI interface {
	GetCapsulesAll() ([]Capsule, error)

	GetCapsuleByID(id string) (Capsule, error)

	PostCapsulesQuery(query Query) (RespQuery[Capsule], error)
}

// Capsule represents the data returned by the SpaceX Capsules V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/capsules/v4/schema.md#capsule-schema
type Capsule struct {
	ID            string   `json:"id"`
	ReuseCount    int      `json:"reuse_count,omitempty"`
	WaterLandings int      `json:"water_landings,omitempty"`
	LandLandings  int      `json:"land_landings,omitempty"`
	LastUpdate    string   `json:"last_update,omitempty"`
	Launches      []string `json:"launches,omitempty"`
	Serial        string   `json:"serial"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Dragon        string   `json:"dragon,omitempty"`
}

func (c Capsule) String() string {
	return stringFn[Capsule](c)
}
