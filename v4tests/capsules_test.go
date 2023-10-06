package v4tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestCapsulesReal(t *testing.T) {
	t.Skip("testing against the real SpaceX API")
	const base = "https://api.spacexdata.com/v4"
	capsules := v4clients.NewCapsulesAPI(base)

	printMany[v4.Capsule](capsules.GetCapsulesAll)
	printOne[v4.Capsule](capsules.GetCapsuleByID, "5e9e2c5bf35918ed873b2664")
}

func TestGetCapsulesAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(capsulesAll)))
	defer server.Close()

	client := v4clients.NewCapsulesAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	capsules, err := client.GetCapsulesAll()
	if err != nil {
		t.Error("❌ TestGetCapsulesAll:", err)
		return
	}

	if len(capsules) != 2 {
		t.Error("❌ TestGetCapsulesAll: retrieved wrong number of capsules")
		return
	}
}

func TestGetCapsuleByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(capsulesOne)))
	defer server.Close()

	client := v4clients.NewCapsulesAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	id := "5e9e2c5bf35918ed873b2664"

	capsule, err := client.GetCapsuleByID(id)
	if err != nil {
		t.Error("❌ TestGetCapsuleByID:", err)
		return
	}

	if capsule.ID != id {
		t.Errorf("❌ TestGetCapsuleByID: id doesn't match: got(%s) expected(%s)", capsule.ID, id)
		return
	}

	// TODO: complete field comparison
}

func TestPostCapsulesQuery(t *testing.T) {
	// TODO: implement
	t.Skip("❗️ PostCapsulesQuery not implemented")

	server := httptest.NewServer(http.HandlerFunc(handler("")))
	defer server.Close()

	client := v4clients.NewCapsulesAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	_, err := client.PostCapsulesQuery(v4.Query{})
	if err != nil {
		t.Error("❌ TestPostCapsulesQuery:", err)
		return
	}
}
