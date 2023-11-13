package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

type StarlinkAPI struct {
	r       requester.Requester
	baseURL string
}

func NewStarlinkAPI(baseURL string) StarlinkAPI {
	return StarlinkAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: baseURL + v4.VersionPrefix + v4.PathStarlink,
	}
}

func (s StarlinkAPI) GetAll() (starlink []v4.Starlink, err error) {
	return starlink, s.r.Get(requester.Config{URL: s.baseURL}, &starlink)
}

func (s StarlinkAPI) GetByID(id string) (starlink v4.Starlink, err error) {
	url := s.baseURL + "/" + id
	return starlink, s.r.Get(requester.Config{URL: url}, &starlink)
}

func (s StarlinkAPI) PostQuery(v4.Query) (v4.RespQuery[v4.Starlink], error) {
	// TODO: implement
	return v4.RespQuery[v4.Starlink]{}, ErrNotImplemented
}
