package v4

const PathStarlink = "/starlink"

// StarlinkAPI defines functions for working with the SpaceX v4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/starlink/v4
type StarlinkAPI interface {
	GetAll() ([]Starlink, error)

	GetByID(string) (Starlink, error)

	PostQuery(Query) (RespQuery[Starlink], error)
}

// SpaceTrack represents the data returned by the SpaceX Starlink V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/starlink/v4/schema.md#starlink-schema
type Starlink struct {
	ID          string     `json:"id"`
	Version     string     `json:"version"`
	Launch      string     `json:"launch"`
	SpaceTrack  SpaceTrack `json:"spaceTrack"`
	Longitude   float32    `json:"longitude"`
	Latitude    float32    `json:"latitude"`
	HeightKM    float32    `json:"height_km"`
	VelocityKMS float32    `json:"velocity_kms"`
}

func (s Starlink) String() string {
	return stringFn[Starlink](s)
}

type SpaceTrack struct {
	RCSSize            string  `json:"RCS_SIZE"`
	CountryCode        string  `json:"COUNTRY_CODE"`
	CreationDate       string  `json:"CREATION_DATE"`
	TLELine2           string  `json:"TLE_LINE2"`
	ObjectName         string  `json:"OBJECT_NAME"`
	TLELine1           string  `json:"TLE_LINE1"`
	CenterName         string  `json:"CENTER_NAME"`
	TLELine0           string  `json:"TLE_LINE0"`
	TimeSystem         string  `json:"TIME_SYSTEM"`
	MeanElementTheory  string  `json:"MEAN_ELEMENT_THEORY"`
	Epoch              string  `json:"EPOCH"`
	DecayDate          string  `json:"DECAY_DATE"`
	Site               string  `json:"SITE"`
	LaunchDate         string  `json:"LAUNCH_DATE"`
	Comment            string  `json:"COMMENT"`
	CCSDS_OMM_VERS     string  `json:"CCSDS_OMM_VERS"`
	ObjectType         string  `json:"OBJECT_TYPE"`
	ObjectID           string  `json:"OBJECT_ID"`
	ClassificationType string  `json:"CLASSIFICATION_TYPE"`
	RefFrame           string  `json:"REF_FRAME"`
	Originator         string  `json:"ORIGINATOR"`
	EphemerisType      int     `json:"EPHEMERIS_TYPE"`
	ElementSetNo       int     `json:"ELEMENT_SET_NO"`
	RevAtEpoch         int     `json:"REV_AT_EPOCH"`
	GPID               int     `json:"GP_ID"`
	File               int     `json:"FILE"`
	NoradCatID         int     `json:"NORAD_CAT_ID"`
	Decayed            int     `json:"DECAYED"`
	Period             float32 `json:"PERIOD"`
	MeanAnomaly        float32 `json:"MEAN_ANOMALY"`
	Periapsis          float32 `json:"PERIAPSIS"`
	MeanMotionDDOT     float32 `json:"MEAN_MOTION_DDOT"`
	MeanMotionDOT      float32 `json:"MEAN_MOTION_DOT"`
	Eccentricity       float32 `json:"ECCENTRICITY"`
	MeanMotion         float32 `json:"MEAN_MOTION"`
	Apoapsis           float32 `json:"APOAPSIS"`
	SemiMajorAxis      float32 `json:"SEMIMAJOR_AXIS"`
	RaOfAscNode        float32 `json:"RA_OF_ASC_NODE"`
	Inclination        float32 `json:"INCLINATION"`
	BStar              float32 `json:"BSTAR"`
	ArgOfPericenter    float32 `json:"ARG_OF_PERICENTER"`
}
