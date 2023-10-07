package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestLandpadsReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewLandpadsAPI(spacexLiveAPIBase)

	printMany[v4.LandingPad](client.GetLandpadsAll)
	printOne[v4.LandingPad](client.GetLandpadByID, "5e9e3032383ecb267a34e7c7")
}

func TestGetLandpadsAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(landpadsAll)))
	defer server.Close()

	client := v4clients.NewLandpadsAPI(server.URL)

	landpads, err := client.GetLandpadsAll()
	if err != nil {
		t.Error("[X] TestGetLandpadsAll:", err)
		return
	}

	if len(landpads) != 2 {
		t.Error("[X] TestGetLandpadsAll:", ErrCollectionLenMismatch)
		return
	}

	// TODO: implement more checks
}

func TestGetLandpadsByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(landpadsOne)))
	defer server.Close()

	client := v4clients.NewLandpadsAPI(server.URL)

	id := "5e9e3032383ecb90a834e7c8"

	landingPad, err := client.GetLandpadByID(id)
	if err != nil {
		t.Error("[X] TestGetLandpadByID:", err)
		return
	}

	if landingPad.ID != id {
		t.Error("[X] TestGetLandpadByID:", ErrIDMismatch)
		return
	}

	// TODO: perform further checks
}

func TestPostLandpadsQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
