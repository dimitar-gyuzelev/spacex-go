package v4

const PathRockets = "/rockets"

// RocketsAPI defines functions for working with the SpaceX Rockets V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/rockets/v4
type RocketsAPI interface {
	GetRocketsAll() ([]Rocket, error)

	GetRocketByID(id string) (Rocket, error)

	PostRocketsQuery(query Query) (RespQuery[Rocket], error)
}

// Rocket represents the data returned by the SpaceX Rockets V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/rockets/v4/schema.md#rocket-schema
type Rocket struct {
	LandingLegs    LandingLegs   `json:"landing_legs"`
	ID             string        `json:"id"`
	Type           string        `json:"type"`
	Name           string        `json:"name"`
	Company        string        `json:"company"`
	Country        string        `json:"country"`
	Description    string        `json:"description"`
	Wikipedia      string        `json:"wikipedia"`
	FirstFlight    string        `json:"first_flight"`
	PayloadWeights []interface{} `json:"payload_weights"`
	FlickrImages   []string      `json:"flickr_images"`
	SecondStage    Stage2        `json:"second_stage"`
	FirstStage     Stage1        `json:"first_stage"`
	Engines        RocketEngines `json:"engines"`
	Mass           Mass          `json:"mass"`
	Stages         int           `json:"stages"`
	CostPerLaunch  int           `json:"cost_per_launch"`
	SuccessRatePCT int           `json:"success_rate_pct"`
	Booster        int           `json:"boosters"`
	Diameter       Distance      `json:"diameter"`
	Height         Distance      `json:"height"`
	Active         bool          `json:"active"`
}

func (r Rocket) String() string {
	return stringFn[Rocket](r)
}

type StageCommon struct {
	Reusable       string  `json:"reusable"`
	FuelAmountTons float32 `json:"fuel_amount_tons"`
	Engines        int     `json:"engines"`
	BurnTimeSec    int     `json:"burn_time_sec"`
}

type Stage1 struct {
	StageCommon
	ThrustSeaLevel ForceUnits `json:"thrust_sea_level"`
	ThrustVacuum   ForceUnits `json:"thrust_vacuum"`
}

type Stage2 struct {
	Payloads Stage2Payload `json:"payloads"`
	StageCommon
	Thrust ForceUnits `json:"thrust"`
}

type Stage2Payload struct {
	Option1          string           `json:"option_1"`
	CompositeFairing CompositeFairing `json:"composite_fairing"`
}

type CompositeFairing struct {
	Height   Distance `json:"height"`
	Diameter Distance `json:"diameter"`
}

type RocketEngines struct {
	Type           string     `json:"type"`
	Version        string     `json:"version"`
	Layout         string     `json:"layout"`
	Propellant1    string     `json:"propellant_1"`
	Propellant2    string     `json:"propellant_2"`
	ISP            ISP        `json:"isp"`
	Number         int        `json:"number"`
	EngineLossMax  int        `json:"engine_loss_max"`
	ThrustSeaLevel ForceUnits `json:"thrust_sea_level"`
	ThrustVacuum   ForceUnits `json:"thrust_vacuum"`
	ThrustToWeight float32    `json:"thrust_to_weight"`
}

type ISP struct {
	SeaLevel int `json:"sea_level"`
	Vacuum   int `json:"vacuum"`
}

type LandingLegs struct {
	Material interface{} `json:"material"`
	Number   int         `json:"number"`
}
