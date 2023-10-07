package v4

const (
	PathDragons      = "/dragons"
	PathDragonsQuery = "/query"
)

// DragonsAPI defines functions to work with SpaceX Dragons V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/dragons/v4
type DragonsAPI interface {
	GetDragonsAll() ([]Dragon, error)

	GetDragonByID(id string) (Dragon, error)

	PostDragonsQuery(query Query) (RespQuery[Dragon], error)
}

// Dragon represents the data returned by the SpaceX Dragons V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/dragons/v4/schema.md#dragon-schema
type Dragon struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	Type               string      `json:"type"`
	FirstFlight        string      `json:"first_flight,omitempty"`
	Wikipedia          string      `json:"wikipedia,omitempty"`
	Description        string      `json:"description,omitempty"`
	FlickImages        []string    `json:"flickr_images,omitempty"`
	HeatShield         DHeatShield `json:"heat_shield"`
	Thrusters          []DThruster `json:"thrusters,omitempty"`
	LaunchPayloadMass  Mass        `json:"launch_payload_mass,omitempty"`
	LaunchPayloadVol   Volume      `json:"launch_payload_vol,omitempty"`
	ReturnPayloadMass  Mass        `json:"return_payload_mass,omitempty"`
	ReturnPayloadVol   Volume      `json:"return_payload_vol,omitempty"`
	PressurizedCapsule DCapsule    `json:"pressurized_capsule,omitempty"`
	Trunk              DTrunk      `json:"trunk,omitempty"`
	HeightWTrunk       Distance    `json:"height_w_trunk,omitempty"`
	Diameter           Distance    `json:"diameter,omitempty"`
	CrewCapacity       int         `json:"crew_capacity"`
	SidewallAngleDeg   int         `json:"sidewall_angle_deg"`
	OrbitDurationYr    int         `json:"orbit_duration_yr"`
	DryMassKG          int         `json:"dry_mass_kg"`
	DryMassLB          int         `json:"dry_mass_lb"`
	Active             bool        `json:"active"`
}

func (d Dragon) String() string {
	return stringFn[Dragon](d)
}

type DThruster struct {
	Type   string `json:"type,omitempty"`
	Fuel1  string `json:"fuel_1,omitempty"`
	Fuel2  string `json:"fuel_2,omitempty"`
	Thrust Thrust `json:"thrust,omitempty"`
	Amount int    `json:"amount,omitempty"`
	Pods   int    `json:"pods,omitempty"`
	ISP    int    `json:"isp,omitempty"`
}

type DHeatShield struct {
	Material    string  `json:"material"`
	DevPartner  string  `json:"dev_partner,omitempty"`
	SizeMeters  float32 `json:"size_meters"`
	TempDegrees float32 `json:"temp_degrees,omitempty"`
}

type DCapsule struct {
	PayloadVolume Volume `json:"payload_volume"`
}

type DTrunk struct {
	TrunkVolume Volume `json:"trunk_volume,omitempty"`
	Cargo       DCargo `json:"cargo,omitempty"`
}

type DCargo struct {
	SolarArray         int  `json:"solar_array"`
	UnpressurizedCargo bool `json:"unpressurized_cargo"`
}
