package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestCrewReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewCrewAPI(spacexLiveAPIBase)

	printMany[v4.CrewMember](client.GetCrewAll)
	printOne[v4.CrewMember](client.GetCrewByID, "5ebf1a6e23a9a60006e03a7a")
}

func TestGetCrewAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(crewAll)))
	defer server.Close()

	client := v4clients.NewCrewAPI(server.URL)

	members, err := client.GetCrewAll()
	if err != nil {
		t.Error("[X] TestGetCrewAll:", err)
		return
	}

	if len(members) != 2 {
		t.Error("[X] TestGetCrewAll:", ErrCollectionLenMismatch)
		return
	}
}

func TestGetCrewByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(crewOne)))
	defer server.Close()

	client := v4clients.NewCrewAPI(server.URL)

	id := "5ebf1a6e23a9a60006e03a7a"

	memeber, err := client.GetCrewByID(id)
	if err != nil {
		t.Error("[X] TestGetCrewByID:", err)
		return
	}

	if memeber.ID != "5ebf1a6e23a9a60006e03a7a" {
		t.Error("[X] TestGetCrewByID:", ErrIDMismatch)
		return
	}

	// TODO: expand on the test
}

func TestPostCrewQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
