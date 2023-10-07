package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// CapsulesAPI is a client for the SpaceX Capsules V4 API.
// Implements v4.CapsulesAPI interface.
type CapsulesAPI struct {
	r       requester.Requester
	baseURL string
}

func NewCapsulesAPI(baseURL string) CapsulesAPI {
	return CapsulesAPI{
		baseURL: baseURL + v4.PathCapsules,
		r:       requester.NewHTTPRequester(Timeout),
	}
}

func (c CapsulesAPI) GetCapsulesAll() (capsules []v4.Capsule, err error) {
	return capsules, c.r.Get(requester.Config{URL: c.baseURL}, &capsules)
}

func (c CapsulesAPI) GetCapsuleByID(id string) (capsule v4.Capsule, err error) {
	url := c.baseURL + "/" + id
	return capsule, c.r.Get(requester.Config{URL: url}, &capsule)
}

func (c CapsulesAPI) PostCapsulesQuery(query v4.Query) (v4.RespQuery[v4.Capsule], error) {
	// TODO: implement
	return v4.RespQuery[v4.Capsule]{}, ErrNotImplemented
}
