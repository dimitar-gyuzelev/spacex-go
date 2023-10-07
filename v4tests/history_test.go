package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestHistoryReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewHistoryAPI(spacexLiveAPIBase)

	printMany[v4.HistEvent](client.GetHistoryAll)
	printOne[v4.HistEvent](client.GetHistoryByID, "5f6fb2cfdcfdf403df37971e")
}

func TestGetHistoryAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(historyAll)))
	defer server.Close()

	client := v4clients.NewHistoryAPI(server.URL)

	events, err := client.GetHistoryAll()
	if err != nil {
		t.Error("[X] TestGetHistoryAll:", err)
		return
	}

	if len(events) != 2 {
		t.Error("[X] TestGetHistoryAll:", ErrCollectionLenMismatch)
		return
	}

	// TODO: more checks are required
}

func TestGetHistoryByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(historyOne)))
	defer server.Close()

	client := v4clients.NewHistoryAPI(server.URL)

	id := ""

	event, err := client.GetHistoryByID(id)
	if err != nil {
		t.Error("[X] TestGetHistoryByID:", err)
		return
	}

	if event.ID != id {
		t.Error("[X] TestGetHistoryByID:", ErrIDMismatch)
		return
	}

	// TODO: more checks are required
}

func TestPostHistoryQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
