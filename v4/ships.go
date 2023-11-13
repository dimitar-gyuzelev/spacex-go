package v4

const PathShips = "/ships"

// ShipsAPI defines functions for working with SpaceX Ships V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/ships/v4
type ShipsAPI interface {
	GetAll() ([]Ship, error)

	GetByID(string) (Ship, error)

	PostQuery(Query) (RespQuery[Ship], error)
}

// Ship represents the data returned by the SpaceX Ships V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/ships/v4/schema.md#ship-schema
type Ship struct {
	HomePort      string   `json:"home_port"`
	Name          string   `json:"name"`
	LegacyID      string   `json:"legacy_id"`
	Model         string   `json:"model"`
	Type          string   `json:"type"`
	ID            string   `json:"id"`
	LastAISUpdate string   `json:"last_ais_update"`
	Link          string   `json:"link"`
	Image         string   `json:"image"`
	Status        string   `json:"status"`
	Roles         []string `json:"roles"`
	Launches      []string `json:"launches"`
	MMSI          int      `json:"mmsi"`
	YearBuilt     int      `json:"year_built"`
	MassLBs       int      `json:"mass_lbs"`
	MassKG        int      `json:"mass_kg"`
	Class         int      `json:"class"`
	ABS           int      `json:"abs"`
	IMO           int      `json:"imo"`
	Latitude      float32  `json:"latitude"`
	Longitude     float32  `json:"longitude"`
	CourseDeg     float32  `json:"course_deg"`
	SpeedKN       float32  `json:"speed_kn"`
	Active        bool     `json:"active"`
}

func (s Ship) String() string {
	return stringFn[Ship](s)
}
