package v4

const (
	PathLaunchpads      = "/launchpads"
	PathLaunchpadsQuery = "/query"
)

// LaunchpadsAPI defines functions for working with SpaceX Launchpads V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/launchpads/v4
type LaunchpadsAPI interface {
	GetLaunchpadsAll() ([]LaunchPad, error)

	GetLaunchpadByID(id string) (LaunchPad, error)

	PostLaunchpadsQuery(query Query) (RespQuery[LaunchPad], error)
}

// LaunchPad represents the data returned by the SpaceX Launchpads V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/launchpads/v4/schema.md#launchpad-schema
type LaunchPad struct {
	Timezone        string   `json:"timezone"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Status          string   `json:"status"`
	Locality        string   `json:"locality"`
	Region          string   `json:"region"`
	ID              string   `json:"id"`
	Rockets         []string `json:"rockets"`
	Launches        []string `json:"launches"`
	LaunchAttempts  int      `json:"launch_attempts"`
	LaunchSuccesses int      `json:"launch_successes"`
	Latitude        float32  `json:"latitude"`
	Longitude       float32  `json:"longitude"`
}

func (l LaunchPad) String() string {
	return stringFn[LaunchPad](l)
}
