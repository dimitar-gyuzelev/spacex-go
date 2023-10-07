package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// HistoryAPI is a client for the SpaceX History V4 API.
// Implements v4.HistoryAPI interface.
type HistoryAPI struct {
	r       requester.Requester
	baseURL string
}

func NewHistoryAPI(baseURL string) HistoryAPI {
	return HistoryAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: baseURL + v4.VersionPrefix + v4.PathHistory,
	}
}

func (h HistoryAPI) GetHistoryAll() (histEvents []v4.HistEvent, err error) {
	return histEvents, h.r.Get(requester.Config{URL: h.baseURL}, &histEvents)
}

func (h HistoryAPI) GetHistoryByID(id string) (histEvent v4.HistEvent, err error) {
	url := h.baseURL + "/" + id
	return histEvent, h.r.Get(requester.Config{URL: url}, &histEvent)
}

func (h HistoryAPI) PostHistoryQuery(query v4.Query) (v4.RespQuery[v4.HistEvent], error) {
	// TODO: implement
	return v4.RespQuery[v4.HistEvent]{}, ErrNotImplemented
}
