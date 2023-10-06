package v4tests

import (
	"fmt"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrewReal(t *testing.T) {
	t.Skip("testing against live SpaceX API")
	client := v4clients.NewCrewAPI(spacexLiveAPIBase)

	printMany[v4.CrewMember](client.GetCrewAll)
	printOne[v4.CrewMember](client.GetCrewByID, "5ebf1a6e23a9a60006e03a7a")
}

func TestGetCrewAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(crewAll)))
	defer server.Close()

	client := v4clients.NewCrewAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	members, err := client.GetCrewAll()
	if err != nil {
		t.Error("❌ TestGetCrewAll:", err)
		return
	}

	if len(members) != 2 {
		t.Error("❌ TestGetCrewAll: retrieved wrong number of crew memebers")
		return
	}
}

func TestGetCrewByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(crewOne)))
	defer server.Close()

	client := v4clients.NewCrewAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	id := "5ebf1a6e23a9a60006e03a7a"

	memeber, err := client.GetCrewByID(id)
	if err != nil {
		t.Error("❌ TestGetCrewByID:", err)
		return
	}

	if memeber.ID != "5ebf1a6e23a9a60006e03a7a" {
		t.Errorf("❌ TestGetCrewByID: ID doesn't match: got(%s) expected(%s)", memeber.ID, id)
		return
	}

	// TODO: expand on the test
}

func TestPostCrewQuery(t *testing.T) {
	// TODO: implement
	t.Skip("PostCrewQuery not implemented")
}
