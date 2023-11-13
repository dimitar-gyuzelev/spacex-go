package v4clients

import (
	"fmt"

	"github.com/dimitar-gyuzelev/spacex-go/requester"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
)

type LaunchesAPI struct {
	r       requester.Requester
	baseURL string
}

func NewLaunchesAPI(baseURL string) LaunchesAPI {
	return LaunchesAPI{
		r:       requester.NewHTTPRequester(Timeout),
		baseURL: fmt.Sprintf("%s%s%s", baseURL, v4.VersionPrefix, v4.PathLaunches),
	}
}

func (l LaunchesAPI) GetLaunchesAll() (launches []v4.Launch, err error) {
	return launches, l.r.Get(requester.Config{URL: l.baseURL}, &launches)
}

func (l LaunchesAPI) GetLaunchesUpcoming() (launches []v4.Launch, err error) {
	url := l.baseURL + v4.PathLaunchesUpcoming
	return launches, l.r.Get(requester.Config{URL: url}, &launches)
}

func (l LaunchesAPI) GetLaunchLatest() (launch v4.Launch, err error) {
	url := l.baseURL + v4.PathLaunchesLatest
	return launch, l.r.Get(requester.Config{URL: url}, &launch)
}

func (l LaunchesAPI) GetLaunchNext() (launch v4.Launch, err error) {
	url := l.baseURL + v4.PathLaunchesNext
	return launch, l.r.Get(requester.Config{URL: url}, &launch)
}

func (l LaunchesAPI) GetLaunchByID(id string) (launch v4.Launch, err error) {
	url := l.baseURL + "/" + id
	return launch, l.r.Get(requester.Config{URL: url}, &launch)
}

func (l LaunchesAPI) GetLaunchesPast() (launches []v4.Launch, err error) {
	url := l.baseURL + v4.PathLaunchesPast
	return launches, l.r.Get(requester.Config{URL: url}, &launches)
}

func (l LaunchesAPI) PostLaunchesQuery(query v4.Query) (v4.RespQuery[v4.Launch], error) {
	return v4.RespQuery[v4.Launch]{}, ErrNotImplemented
}
