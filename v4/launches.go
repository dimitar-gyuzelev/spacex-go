package v4

const (
	PathLaunches         = "/launches"
	PathLaunchesUpcoming = "/upcoming"
	PathLaunchesLatest   = "/latest"
	PathLaunchesNext     = "/next"
	PathLaunchesPast     = "/past"
	PathLaunchesQuery    = "/query"
)

// LaunchesAPI defines functions for working with the SpaceX Launches V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/launches/v4
type LaunchesAPI interface {
	GetLaunchesAll() ([]Launch, error)

	GetLaunchesUpcoming() ([]Launch, error)

	GetLaunchesPast() ([]Launch, error)

	GetLaunchLatest() (Launch, error)

	GetLaunchNext() (Launch, error)

	GetLaunchByID(id string) (Launch, error)

	PostLaunchesQuery(query Query) (RespQuery[Launch], error)
}

// Launch represents the data returned by the SpaceX Launches V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/launches/v4/schema.md#launch-schema
type Launch struct {
	Rocket             string       `json:"rocket,omitempty"`
	Name               string       `json:"name"`
	DateUTC            string       `json:"date_utc"`
	DateLocal          string       `json:"date_local"`
	DatePrecision      string       `json:"date_precision"`
	StaticFireDateUTC  string       `json:"static_fire_date_utc,omitempty"`
	Launchpad          string       `json:"launchpad,omitempty"`
	ID                 string       `json:"id"`
	Details            string       `json:"details"`
	Links              LaunchLinks  `json:"links"`
	Cores              []LaunchCore `json:"cores"`
	Capsules           []string     `json:"capsules"`
	Failures           []Failure    `json:"failures,omitempty"`
	Crew               []string     `json:"crew"`
	Ships              []string     `json:"ships,omitempty"`
	Payloads           []string     `json:"payloads,omitempty"`
	Fairings           Fairings     `json:"fairings"`
	Window             int          `json:"window,omitempty"`
	FlightNumber       int          `json:"flight_number"`
	DateUnix           int          `json:"date_unix"`
	StaticFireDateUnix int          `json:"static_fire_date_unix,omitempty"`
	Net                bool         `json:"net,omitempty"`
	TDB                bool         `json:"tdb"`
	Success            bool         `json:"success,omitempty"`
	AutoUpdate         bool         `json:"auto_update"`
	Upcoming           bool         `json:"upcoming"`
}

func (l Launch) String() string {
	return stringFn[Launch](l)
}

type Failure struct {
	Reason   string `json:"reason,omitempty"`
	Time     int    `json:"time,omitempty"`
	Altitude int    `json:"altitude,omitempty"`
}

type Fairings struct {
	Ships           []string `json:"ships,omitempty"`
	Reused          bool     `json:"reused,omitempty"`
	RecoveryAttempt bool     `json:"recovery_attempt,omitempty"`
	Recovered       bool     `json:"recovered,omitempty"`
}

type LaunchCore struct {
	Core           string `json:"core,omitempty"`
	LandingType    string `json:"landing_type,omitempty"`
	Landpad        string `json:"landpad,omitempty"`
	Flight         int    `json:"flight,omitempty"`
	Gridfins       bool   `json:"gridfins,omitempty"`
	Legs           bool   `json:"legs,omitempty"`
	Reused         bool   `json:"reused,omitempty"`
	LandingAttempt bool   `json:"landing_attempt,omitempty"`
	LandingSuccess bool   `json:"landing_success,omitempty"`
}

type LaunchLinks struct {
	Reddit    LaunchLinkReddit `json:"reddit,omitempty"`
	Patch     LaunchLinkPatch  `json:"patch,omitempty"`
	PressKit  string           `json:"presskit,omitempty"`
	WebCast   string           `json:"webcast,omitempty"`
	YouTubeID string           `json:"youtube_id,omitempty"`
	Article   string           `json:"article,omitempty"`
	Wikipedia string           `json:"wikipedia,omitempty"`
	Flickr    LaunchLinkFlickr `json:"flickr,omitempty"`
}

type LaunchLinkPatch struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type LaunchLinkReddit struct {
	Campaign string `json:"campaign,omitempty"`
	Launch   string `json:"launch,omitempty"`
	Media    string `json:"media,omitempty"`
	Recovery string `json:"recovery,omitempty"`
}

type LaunchLinkFlickr struct {
	Small    []string `json:"small,omitempty"`
	Original []string `json:"original,omitempty"`
}
