package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// DragonsAPI is a client for the SpaceX Dragons V4 API.
// Implements the v4.DragonsAPI interface.
type DragonsAPI struct {
	r       requester.Requester
	baseURL string
}

func NewDragonsAPI(baseURL string) DragonsAPI {
	return DragonsAPI{
		baseURL: baseURL + v4.VersionPrefix + v4.PathDragons,
		r:       requester.NewHTTPRequester(Timeout),
	}
}

func (d DragonsAPI) GetDragonsAll() (dragons []v4.Dragon, err error) {
	return dragons, d.r.Get(requester.Config{URL: d.baseURL}, &dragons)
}

func (d DragonsAPI) GetDragonByID(id string) (dragon v4.Dragon, err error) {
	url := d.baseURL + "/" + id
	return dragon, d.r.Get(requester.Config{URL: url}, &dragon)
}

func (d DragonsAPI) PostDragonsQuery(query v4.Query) (v4.RespQuery[v4.Dragon], error) {
	return v4.RespQuery[v4.Dragon]{}, ErrNotImplemented
}
