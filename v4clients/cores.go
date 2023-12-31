package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// CoresAPI is a client for the SpaceX Cores V4 API.
// Implements v4.CoresAPI interface.
type CoresAPI struct {
	r       requester.Requester
	baseURL string
}

func NewCoresAPI(baseURL string) CoresAPI {
	return CoresAPI{
		baseURL: baseURL + v4.PathCores,
		r:       requester.NewHTTPRequester(Timeout),
	}
}

func (c CoresAPI) GetCoresAll() (cores []v4.Core, err error) {
	return cores, c.r.Get(requester.Config{URL: c.baseURL}, &cores)
}

func (c CoresAPI) GetCoreByID(id string) (core v4.Core, err error) {
	url := c.baseURL + "/" + id
	return core, c.r.Get(requester.Config{URL: url}, &core)
}

func (c CoresAPI) PostCoresQuery(query v4.Query) (v4.RespQuery[v4.Core], error) {
	return v4.RespQuery[v4.Core]{}, ErrNotImplemented
}
