package v4

import "fmt"

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
	ReuseCount    int      `json:"reuse_count"`
	WaterLandings int      `json:"water_landings"`
	LandLandings  int      `json:"land_landings"`
	LastUpdate    string   `json:"last_update"`
	Launches      []string `json:"launches"`
	Serial        string   `json:"serial"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Dragon        string   `json:"dragon"`
}

func (c Capsule) String() string {
	return fmt.Sprintf(`
		{
			ID: %s,
			ReuseCount: %d,
			WaterLandings: %d,
			LandLandings: %d,
			LastUpdate: %s,
			Launches: %v,
			Serial: %s,
			Status: %s,
			Type: %s,
			Dragon: %s,
		}`, c.ID, c.ReuseCount, c.WaterLandings, c.LandLandings, c.LastUpdate, c.Launches, c.Serial, c.Status, c.Type, c.Dragon)
}