package v4

const PathLandpads = "/landpads"

// LandpadsAPI defines functions to work with SpaceX Landpads v4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/landpads/v4
type LandpadsAPI interface {
	GetLandpadsAll() ([]LandingPad, error)

	GetLandpadByID(id string) (LandingPad, error)

	PostLandpadsQuery(query Query) (RespQuery[LandingPad], error)
}

// LandingPad represents the data returned by the SpaceX Landpads V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/landpads/v4/schema.md#landing-pad-schema
type LandingPad struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	FullName         string   `json:"full_name"`
	Status           string   `json:"status"`
	Type             string   `json:"type"`
	Locality         string   `json:"locality"`
	Region           string   `json:"region"`
	Wikipedia        string   `json:"wikipedia"`
	Details          string   `json:"details"`
	Launches         []string `json:"launches"`
	Latitude         float32  `json:"latitude"`
	Longitude        float32  `json:"longitude"`
	LandingAttempts  int      `json:"landing_attempts"`
	LandingSuccesses int      `json:"landing_successes"`
}

func (l LandingPad) String() string {
	return stringFn[LandingPad](l)
}
