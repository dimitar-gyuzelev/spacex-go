package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

type RocketsAPI struct {
	r       requester.Requester
	baseURL string
}

func NewRocketsAPI(baseURL string) RocketsAPI {
	return RocketsAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: baseURL + v4.VersionPrefix + v4.PathRockets,
	}
}

func (r RocketsAPI) GetRocketsAll() (rockets []v4.Rocket, err error) {
	return rockets, r.r.Get(requester.Config{URL: r.baseURL}, &rockets)
}

func (r RocketsAPI) GetRocketByID(id string) (rocket v4.Rocket, err error) {
	url := r.baseURL + "/" + id
	return rocket, r.r.Get(requester.Config{URL: url}, &rocket)
}

func (r RocketsAPI) PostRocketsQuery(query v4.Query) (v4.RespQuery[v4.Rocket], error) {
	// TODO: implement
	return v4.RespQuery[v4.Rocket]{}, ErrNotImplemented
}
