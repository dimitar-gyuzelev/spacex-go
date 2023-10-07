package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestDragonsReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewDragonsAPI(spacexLiveAPIBase)

	printMany[v4.Dragon](client.GetDragonsAll)
	printOne[v4.Dragon](client.GetDragonByID, "5e9d058759b1ff74a7ad5f8f")
}

func TestGetDragonsAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(dragonsAll)))
	defer server.Close()

	client := v4clients.NewDragonsAPI(server.URL)

	dragons, err := client.GetDragonsAll()
	if err != nil {
		t.Error("[X] TestGetDragonsAll:", err)
		return
	}

	if len(dragons) != 2 {
		t.Error("[X] TestGetDragonsAll:", ErrCollectionLenMismatch)
		return
	}

	// TODO: perform more checks
}

func TestDragonGetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(dragonsOne)))
	defer server.Close()

	client := v4clients.NewDragonsAPI(server.URL)

	id := "5e9d058759b1ff74a7ad5f8f"

	dragon, err := client.GetDragonByID(id)
	if err != nil {
		t.Error("[X] TestDragonByID: ", err)
		return
	}

	if dragon.ID != id {
		t.Error("[X] TestDragonGetByID:", ErrIDMismatch)
		return
	}

	// TODO: perform more checks
}

func TestPostDragonsQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
