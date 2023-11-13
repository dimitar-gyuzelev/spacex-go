package v4

const (
	PathPayloads      = "/payloads"
	PathPayloadsQuery = "/query"
)

// PayloadsAPI defines functions for working with the SpaceX Payloads V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/payloads/v4
type PayloadsAPI interface {
	GetPayloadsAll() ([]Payload, error)

	GetPayloadByID(id string) (Payload, error)

	PostPayloadsQuery(query Query) (RespQuery[Payload], error)
}

// Payload represents the data returned by the SpaceX Payloads V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/payloads/v4/schema.md#payload-schema
type Payload struct {
	Epoch           string        `json:"epoch"`
	Name            string        `json:"name"`
	Type            string        `json:"type"`
	ReferenceSystem string        `json:"reference_system"`
	Launch          string        `json:"launch"`
	Oribit          string        `json:"orbit"`
	Regime          string        `json:"egime"`
	ID              string        `json:"id"`
	NoradIDs        []int         `json:"norad_ids"`
	Manufacturers   []string      `json:"manufacturers"`
	Nationalities   []string      `json:"nationalities"`
	Customers       []string      `json:"customers"`
	Dragon          PayloadDragon `json:"dragon"`
	Longitude       float32       `json:"longitude"`
	PeriodMin       float32       `json:"period_min"`
	SemiMajorAxisKM float32       `json:"semi_major_axis_km"`
	Eccentricity    float32       `json:"eccentricity"`
	PeriapsisKM     float32       `json:"periapsis_km"`
	ApoapsisKM      float32       `json:"apoapsis_km"`
	InclinationDeg  float32       `json:"inclination_deg"`
	MassLBs         float32       `json:"mass_lbs"`
	LifespanYears   float32       `json:"lifespan_years"`
	MassKG          float32       `json:"mass_kg"`
	MeanMotion      float32       `json:"mean_motion"`
	Raan            float32       `json:"raan"`
	ArgOfPericenter float32       `json:"arg_of_pericenter"`
	MeanAnomaly     float32       `json:"mean_anomaly"`
	Reused          bool          `json:"reused"`
}

func (p Payload) String() string {
	return stringFn[Payload](p)
}

type PayloadDragon struct {
	Capsule         string  `json:"capsule"`
	Manifest        string  `json:"manifest"`
	FlightTimeSec   int     `json:"flight_time_sec"`
	MassReturnedKG  float32 `json:"mass_returned_kg"`
	MassReturnedLBs float32 `json:"mass_returned_lbs"`
	WaterLanding    bool    `json:"water_landing"`
	LandLanding     bool    `json:"land_landing"`
}
