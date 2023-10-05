package v4clients

import (
	"fmt"
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// CapsulesAPI is a client for the SpaceX Capsules V4 API.
// Implements v4.CapsulesAPI interface
type CapsulesAPI struct {
	baseURL string
	r       requester.Requester
}

func NewCapsulesAPI(baseURL string) CapsulesAPI {
	return CapsulesAPI{
		baseURL: fmt.Sprintf("%s%s", baseURL, v4.PathCapsules),
		r:       requester.NewHTTPRequester(Timeout),
	}
}

func (c CapsulesAPI) GetCapsulesAll() ([]v4.Capsule, error) {
	var capsules []v4.Capsule
	return capsules, c.r.Get(requester.Config{URL: c.baseURL}, &capsules)
}

func (c CapsulesAPI) GetCapsuleByID(id string) (v4.Capsule, error) {
	var capsule v4.Capsule
	url := fmt.Sprintf("%s/%s", c.baseURL, id)
	return capsule, c.r.Get(requester.Config{URL: url}, &capsule)
}

func (c CapsulesAPI) PostCapsulesQuery(query v4.Query) (v4.RespQuery[v4.Capsule], error) {
	// TODO: implement
	return v4.RespQuery[v4.Capsule]{}, fmt.Errorf("not implemented")
}
