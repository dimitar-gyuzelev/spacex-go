package v4

const PathCrew = "/crew"

// CrewAPI defines functions to work with SpaceX Crews V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/crew/v4
type CrewAPI interface {
	GetCrewAll() ([]CrewMember, error)

	GetCrewByID(id string) (CrewMember, error)

	PostCrewQuery(query Query) (RespQuery[CrewMember], error)
}

// CrewMember represents the data returned by the SpaceX Crew V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/crew/v4/schema.md#crew-schema
type CrewMember struct {
	ID        string   `json:"id"`
	Name      string   `json:"name,omitempty"`
	Status    string   `json:"status"`
	Agency    string   `json:"agency,omitempty"`
	Image     string   `json:"image,omitempty"`
	Wikipedia string   `json:"wikipedia,omitempty"`
	Launches  []string `json:"launches,omitempty"`
}

func (c CrewMember) String() string {
	return stringFn[CrewMember](c)
}
