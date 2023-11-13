package spacexgo

import (
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

type ClientV4 struct {
	Capsules   v4.CapsulesAPI
	Cores      v4.CoresAPI
	Dragons    v4.DragonsAPI
	History    v4.HistoryAPI
	Landpads   v4.LandpadsAPI
	Launches   v4.LaunchesAPI
	Launchpads v4.LaunchpadsAPI
	Payloads   v4.PayloadsAPI
	Roadster   v4.RoadsterAPI
	Rockets    v4.RocketsAPI
	Ships      v4.ShipsAPI
}

func NewClient(baseURL string) ClientV4 {
	baseURL = baseURL + v4.VersionPrefix
	return ClientV4{
		Capsules:   v4clients.NewCapsulesAPI(baseURL),
		Cores:      v4clients.NewCoresAPI(baseURL),
		Dragons:    v4clients.NewDragonsAPI(baseURL),
		History:    v4clients.NewHistoryAPI(baseURL),
		Landpads:   v4clients.NewLandpadsAPI(baseURL),
		Launches:   v4clients.NewLaunchesAPI(baseURL),
		Launchpads: v4clients.NewLaunchpadsAPI(baseURL),
		Payloads:   v4clients.NewPayloadsAPI(baseURL),
		Roadster:   v4clients.NewRoadsterAPI(baseURL),
		Rockets:    v4clients.NewRocketsAPI(baseURL),
		Ships:      v4clients.NewShipsAPI(baseURL),
	}
}
