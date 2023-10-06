package v4

import "fmt"

const PathCores = "/cores"

// CoresAPI defines the functions to work with SpaceX Cores V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/cores/v4
type CoresAPI interface{
	GetCoresAll() ([]Core, error)
	
	GetCoreByID(id string) (Core, error)
	
	PostCoresQuery(query Query) (RespQuery[Core], error)
}

// Core represents the data returned by SpaceX Cores V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/cores/v4/schema.md#core-schema
type Core struct {
	ID           string   `json:"id"`
	Serial       string   `json:"serial"`
	Block        int      `json:"block"`
	Status       string   `json:"status"`
	ReuseCount   int      `json:"reuse_count"`
	RTLSAttempts int      `json:"rtls_attempts"`
	RTLSLandings int      `json:"rtls_landings"`
	ASDSAttempts int      `json:"asds_attempts"`
	ASDSLandings int      `json:"asds_landings"`
	LastUpdate   string   `json:"last_update"`
	Launches     []string `json:"launches"`
}

func (c Core) String() string {
	return fmt.Sprintf(`
		{
			ID: %s,
			Serial: %s,
			Block: %d,
			Status: %s,
			ReuseCount: %d,
			RTLSAttempts: %d,
			RTLSLandings: %d,
			ASDSAttempts: %d,
			ASDSLandings: %d,
			LastUpdate: %s,
			Launches: %v
		}`, c.ID, c.Serial, c.Block, c.Status, c.ReuseCount, c.RTLSAttempts, c.RTLSLandings, c.ASDSAttempts, c.ASDSLandings, c.LastUpdate, c.Launches)
}
