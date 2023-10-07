package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

// LandpadsAPI is a client for the SpaceX Landpads V4 API.
// Implements v4.LandpadsAPI interface.
type LandpadsAPI struct {
	r       requester.Requester
	baseURL string
}

func NewLandpadsAPI(baseURL string) LandpadsAPI {
	return LandpadsAPI{
		baseURL: baseURL + v4.VersionPrefix + v4.PathLandpads,
		r:       requester.NewHTTPRequester(Timeout),
	}
}

func (l LandpadsAPI) GetLandpadsAll() (landpads []v4.LandingPad, err error) {
	return landpads, l.r.Get(requester.Config{URL: l.baseURL}, &landpads)
}

func (l LandpadsAPI) GetLandpadByID(id string) (landpad v4.LandingPad, err error) {
	url := l.baseURL + "/" + id
	return landpad, l.r.Get(requester.Config{URL: url}, &landpad)
}

func (l LandpadsAPI) PostLandpadsQuery(query v4.Query) (v4.RespQuery[v4.LandingPad], error) {
	// TODO: implement
	return v4.RespQuery[v4.LandingPad]{}, ErrNotImplemented
}
