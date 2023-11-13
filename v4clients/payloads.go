package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// PayloadsAPI is a client for the SpaceX Payloads V4 API.
// Implements the v4.PayloadsAPI interface.
type PayloadsAPI struct {
	r       requester.Requester
	baseURL string
}

func NewPayloadsAPI(baseURL string) PayloadsAPI {
	return PayloadsAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: baseURL + v4.VersionPrefix + v4.PathPayloads,
	}
}

func (p PayloadsAPI) GetPayloadsAll() (payloads []v4.Payload, err error) {
	return payloads, p.r.Get(requester.Config{URL: p.baseURL}, &payloads)
}

func (p PayloadsAPI) GetPayloadByID(id string) (payload v4.Payload, err error) {
	url := p.baseURL + "/" + id
	return payload, p.r.Get(requester.Config{URL: url}, &payload)
}

func (p PayloadsAPI) PostPayloadsQuery(query v4.Query) (v4.RespQuery[v4.Payload], error) {
	return v4.RespQuery[v4.Payload]{}, ErrNotImplemented
}
