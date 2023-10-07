package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestCapsulesReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	capsules := v4clients.NewCapsulesAPI(spacexLiveAPIBase)

	printMany[v4.Capsule](capsules.GetCapsulesAll)
	printOne[v4.Capsule](capsules.GetCapsuleByID, "5e9e2c5bf35918ed873b2664")
}

func TestGetCapsulesAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(capsulesAll)))
	defer server.Close()

	client := v4clients.NewCapsulesAPI(server.URL)

	capsules, err := client.GetCapsulesAll()
	if err != nil {
		t.Error("[X] TestGetCapsulesAll:", err)
		return
	}

	if len(capsules) != 2 {
		t.Error("[X] TestGetCapsulesAll:", ErrCollectionLenMismatch)
		return
	}
}

func TestGetCapsuleByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(capsulesOne)))
	defer server.Close()

	client := v4clients.NewCapsulesAPI(server.URL)

	id := "5e9e2c5bf35918ed873b2664"

	capsule, err := client.GetCapsuleByID(id)
	if err != nil {
		t.Error("[X] TestGetCapsuleByID:", err)
		return
	}

	if capsule.ID != id {
		t.Error("[X] TestGetCapsuleByID:", ErrIDMismatch)
		return
	}

	// TODO: complete field comparison
}

func TestPostCapsulesQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
