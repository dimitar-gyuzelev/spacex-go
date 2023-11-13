package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestLaunchesReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewLaunchesAPI(spacexLiveAPIBase)

	printMany[v4.Launch](client.GetLaunchesAll)
	printMany[v4.Launch](client.GetLaunchesUpcoming)
	printMany[v4.Launch](client.GetLaunchesPast)
	printOneNoArgs[v4.Launch](client.GetLaunchLatest)
	printOneNoArgs[v4.Launch](client.GetLaunchNext)
	printOne[v4.Launch](client.GetLaunchByID, "5eb87cd9ffd86e000604b32a")
}

func TestGetLaunchesAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchesAll)))
	defer server.Close()

	client := v4clients.NewLaunchesAPI(server.URL)

	launches, err := client.GetLaunchesAll()
	if err != nil {
		t.Error("[X] TestGetLaunchesAll:", err)
		return
	}

	if len(launches) != 2 {
		t.Error("[X] TestGetLaunchesAll:", ErrCollectionLenMismatch)
	}

	// TODO: implement more checks
}

func TestGetLaunchesUpcoming(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchesUpcoming)))
	defer server.Close()

	client := v4clients.NewLaunchesAPI(server.URL)

	upcoming, err := client.GetLaunchesUpcoming()
	if err != nil {
		t.Error("[X] TestGetLaunchesUpcoming:", err)
		return
	}

	if len(upcoming) != 2 {
		t.Error("[X] TestGetLaunchesUpcoming", ErrCollectionLenMismatch)
		return
	}

	// TODO: implement more checks
}

func TestGetLaunchesLatest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchesLatest)))
	defer server.Close()

	client := v4clients.NewLaunchesAPI(server.URL)

	_, err := client.GetLaunchLatest()
	if err != nil {
		t.Error("[X] TestGetLaunchesLatest:", err)
		return
	}

	// TODO: implement more checks
}

func TestGetLaunchesNext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchesNext)))
	defer server.Close()

	client := v4clients.NewLaunchesAPI(server.URL)

	_, err := client.GetLaunchNext()
	if err != nil {
		t.Error("[X] TestGetLaunchesNext:", err)
		return
	}

	// TODO: implement more checks
}

func TestGetLaunchesPast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchesPast)))
	defer server.Close()

	client := v4clients.NewLaunchesAPI(server.URL)

	past, err := client.GetLaunchesPast()
	if err != nil {
		t.Error("[X] TestGetLaunchesNext:", err)
		return
	}

	if len(past) != 2 {
		t.Error("[X] TestGetLaunchesPast:", ErrCollectionLenMismatch)
		return
	}

	// TODO: implement more checks
}

func TestGetLaunchByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(launchesOne)))
	defer server.Close()

	client := v4clients.NewLaunchesAPI(server.URL)

	id := ""

	launch, err := client.GetLaunchByID(id)
	if err != nil {
		t.Error("[X] TestGetLaunchByID:", err)
		return
	}

	if launch.ID != id {
		t.Error("[X] TestGetLaunchByID:", ErrIDMismatch)
		return
	}

	// TODO: implement more checks
}

func TestPostLaunchesQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
