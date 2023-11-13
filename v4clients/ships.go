package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

type ShipsAPI struct {
	r       requester.Requester
	baseURL string
}

func NewShipsAPI(baseURL string) ShipsAPI {
	return ShipsAPI{
		requester.NewHTTPRequester(Timeout),
		baseURL + v4.VersionPrefix + v4.PathShips,
	}
}

func (s ShipsAPI) GetAll() (ships []v4.Ship, err error) {
	return ships, s.r.Get(requester.Config{URL: s.baseURL}, &ships)
}

func (s ShipsAPI) GetByID(id string) (ship v4.Ship, err error) {
	url := s.baseURL + "/" + id
	return ship, s.r.Get(requester.Config{URL: url}, &ship)
}

func (s ShipsAPI) PostQuery(v4.Query) (v4.RespQuery[v4.Ship], error) {
	// TODO: implement
	return v4.RespQuery[v4.Ship]{}, ErrNotImplemented
}
