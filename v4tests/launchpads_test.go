package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestLaunchpadsReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewLaunchpadsAPI(spacexLiveAPIBase)

	printMany[v4.LaunchPad](client.GetLaunchpadsAll)
	printOne[v4.LaunchPad](client.GetLaunchpadByID, "5e9e4501f5090910d4566f83")
}

func TestGetLaunchpadsAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchpadsAll)))
	defer server.Close()

	client := v4clients.NewLaunchpadsAPI(server.URL)

	launchpads, err := client.GetLaunchpadsAll()
	if err != nil {
		t.Error("[X] TestGetLaunchpadsAll:", err)
		return
	}

	if len(launchpads) != 2 {
		t.Error("[X] TestGetLaunchpadsAll:", ErrCollectionLenMismatch)
		return
	}

	// TODO: implement more checks
}

func TestGetLaunchpadByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchpadsOne)))
	defer server.Close()

	client := v4clients.NewLaunchpadsAPI(server.URL)

	id := "5e9e4501f5090910d4566f83"

	pad, err := client.GetLaunchpadByID(id)
	if err != nil {
		t.Error("[X] TestGetLaunchpadByID:", err)
		return
	}

	if pad.ID != id {
		t.Error("[X] TestGetLaunchpadByID:", ErrIDMismatch)
		return
	}

	// TODO: implement more checks
}

func TestPostLaunchpadsQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
