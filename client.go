package spacexgo

import (
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

type Client struct {
	v4.CapsulesAPI
}

func NewClient(baseURL string) Client {
	return Client{
		CapsulesAPI: v4clients.NewCapsulesAPI(baseURL),
	}
}
