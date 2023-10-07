package spacexgo

import (
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

type ClientV4 struct {
	v4.CapsulesAPI
	v4.CoresAPI
	v4.DragonsAPI
	v4.HistoryAPI
	v4.LandpadsAPI
}

func NewClient(baseURL string) ClientV4 {
	baseURL = baseURL + v4.VersionPrefix
	return ClientV4{
		CapsulesAPI: v4clients.NewCapsulesAPI(baseURL),
		CoresAPI:    v4clients.NewCoresAPI(baseURL),
		DragonsAPI:  v4clients.NewDragonsAPI(baseURL),
		HistoryAPI:  v4clients.NewHistoryAPI(baseURL),
		LandpadsAPI: v4clients.NewLandpadsAPI(baseURL),
	}
}
