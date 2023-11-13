package v4

const (
	PathRoadster      = "/roadster"
	PathRoadsterQuery = "/query"
)

// RoadsterAPI defines funcitons for working with the SpaceX Roadster V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/roadster/v4
type RoadsterAPI interface {
	GetRoadster() (Roadster, error)

	PostRoadsterQuery(query Query) (RespQuery[Roadster], error)
}

// Roadster represents the data returned by the SpaceX Roadster V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/roadster/v4/schema.md#roadster-schema
type Roadster struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	LaunchDateUTC   string   `json:"launch_date_utc"`
	Details         string   `json:"details"`
	Video           string   `json:"video"`
	Wikipedia       string   `json:"wikipedia"`
	OrbitType       string   `json:"orbit_type"`
	FlickrImages    []string `json:"flickr_images"`
	LaunchDateUnix  int      `json:"launch_date_unix"`
	NoradID         int      `json:"norad_id"`
	Eccentricity    float32  `json:"eccentricity"`
	EarthDistanceKM float32  `json:"earth_distance_km"`
	PeriapsisAU     float32  `json:"periapsis_au"`
	Inclination     float32  `json:"inclination"`
	Longitude       float32  `json:"longitude"`
	PeriapsisArg    float32  `json:"periapsis_arg"`
	PeriodDays      float32  `json:"period_days"`
	SpeedKPH        float32  `json:"speed_kph"`
	SpeedMPH        float32  `json:"speed_mph"`
	SemiMajorAxisAU float32  `json:"semi_major_axis_au"`
	EarthDistanceMI float32  `json:"earth_distance_mi"`
	MarsDistanceKM  float32  `json:"mars_distance_km"`
	MarsDistanceMI  float32  `json:"mars_distance_mi"`
	ApoapsisAU      float32  `json:"apoapsis_au"`
	EpochJD         float32  `json:"epoch_jd"`
	LaunchMassLBs   float32  `json:"launch_mass_lbs"`
	LaunchMassKG    float32  `json:"launch_mass_kg"`
}

func (r Roadster) String() string {
	return stringFn[Roadster](r)
}
