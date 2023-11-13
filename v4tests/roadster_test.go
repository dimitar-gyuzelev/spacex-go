package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestRoadsterReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)
	client := v4clients.NewRoadsterAPI(spacexLiveAPIBase)

	printOneNoArgs[v4.Roadster](client.GetRoadster)
}

func TestGetRoadster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(roadsterOne)))
	defer server.Close()

	client := v4clients.NewRoadsterAPI(server.URL)

	_, err := client.GetRoadster()
	if err != nil {
		t.Error("[X] TestGetRoadster:", err)
		return
	}

	// TODO: implement more checks
}

func TestPostRoadsterQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
