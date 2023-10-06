package v4clients

import (
	"fmt"
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// CoresAPI is a client for the SpaceX Cores V4 API.
// Implements v4.CoresAPI interface.
type CoresAPI struct {
	baseURL string
	r       requester.Requester
}

func NewCoresAPI(baseURL string) CoresAPI {
	return CoresAPI{
		baseURL: fmt.Sprintf("%s%s", baseURL, v4.PathCores),
		r: requester.NewHTTPRequester(Timeout),
	}
}

func (c CoresAPI) GetCoresAll() (cores []v4.Core, err error) {
	return cores, c.r.Get(requester.Config{URL: c.baseURL}, &cores)
}

func (c CoresAPI) GetCoreByID(id string) (core v4.Core, err error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, id)
	return core, c.r.Get(requester.Config{URL: url}, &core)
}

func (c CoresAPI) PostCoresQuery(query v4.Query) (v4.RespQuery[v4.Core], error) {
	return v4.RespQuery[v4.Core]{}, fmt.Errorf("not implemented")
}