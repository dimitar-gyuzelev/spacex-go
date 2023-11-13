package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

type RoadsterAPI struct {
	r       requester.Requester
	baseURL string
}

func NewRoadsterAPI(baseURL string) RoadsterAPI {
	return RoadsterAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: baseURL + v4.VersionPrefix + v4.PathRoadster,
	}
}

func (r RoadsterAPI) GetRoadster() (roadster v4.Roadster, err error) {
	return roadster, r.r.Get(requester.Config{URL: r.baseURL}, &roadster)
}

func (r RoadsterAPI) PostRoadsterQuery(query v4.Query) (v4.RespQuery[v4.Roadster], error) {
	// TODO: implement
	return v4.RespQuery[v4.Roadster]{}, ErrNotImplemented
}
