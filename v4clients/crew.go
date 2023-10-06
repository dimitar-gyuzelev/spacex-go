package v4clients

import (
	"fmt"
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// CrewAPI is a client for the SpaceX Capsules V4 API.
// Implements v4.CrewAPI interface.
type CrewAPI struct {
	baseURL string
	r       requester.Requester
}

func NewCrewAPI(baseURL string) CrewAPI {
	return CrewAPI{
		baseURL: fmt.Sprintf("%s%s", baseURL, v4.PathCrew),
		r:       requester.NewHTTPRequester(Timeout),
	}
}

func (c CrewAPI) GetCrewAll() (members []v4.CrewMember, err error) {
	return members, c.r.Get(requester.Config{URL: c.baseURL}, &members)
}

func (c CrewAPI) GetCrewByID(id string) (member v4.CrewMember, err error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, id)
	return member, c.r.Get(requester.Config{URL: url}, &member)
}

func (c CrewAPI) PostCrewQuery(query v4.Query) (v4.RespQuery[v4.CrewMember], error) {
	// TODO: implement
	return v4.RespQuery[v4.CrewMember]{}, fmt.Errorf("not implemented")
}
