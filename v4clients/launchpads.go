package v4clients

import (
	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

type LaunchpadsAPI struct {
	r       requester.Requester
	baseURL string
}

func NewLaunchpadsAPI(baseURL string) LaunchpadsAPI {
	return LaunchpadsAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: baseURL + v4.VersionPrefix + v4.PathLaunchpads,
	}
}

func (l LaunchpadsAPI) GetLaunchpadsAll() (pads []v4.LaunchPad, err error) {
	return pads, l.r.Get(requester.Config{URL: l.baseURL}, &pads)
}

func (l LaunchpadsAPI) GetLaunchpadByID(id string) (pad v4.LaunchPad, err error) {
	url := l.baseURL + "/" + id
	return pad, l.r.Get(requester.Config{URL: url}, &pad)
}

func (l LaunchpadsAPI) PostLaunchpadsQuery(query v4.Query) (v4.RespQuery[v4.LaunchPad], error) {
	return v4.RespQuery[v4.LaunchPad]{}, ErrNotImplemented
}
