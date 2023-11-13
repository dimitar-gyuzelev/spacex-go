package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestPayloadsReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	client := v4clients.NewPayloadsAPI(spacexLiveAPIBase)

	printMany[v4.Payload](client.GetPayloadsAll)
	printOne[v4.Payload](client.GetPayloadByID, "5eb0e4b5b6c3bb0006eeb1e1")
}

func TestGetPayloadsAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(payloadsAll)))
	defer server.Close()

	client := v4clients.NewPayloadsAPI(server.URL)

	payloads, err := client.GetPayloadsAll()
	if err != nil {
		t.Error("[X] TestGetPayloadsAll:", err)
		return
	}

	if len(payloads) != 2 {
		t.Error("[X] TestGetPayloadsAll:", ErrCollectionLenMismatch)
		return
	}

	// TODO: implement more checks
}

func TestGetPayloadByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(payloadsOne)))
	defer server.Close()

	client := v4clients.NewPayloadsAPI(server.URL)

	id := "5eb0e4b5b6c3bb0006eeb1e1"

	payload, err := client.GetPayloadByID(id)
	if err != nil {
		t.Error("[X] TestGetPayloadByID:", err)
		return
	}

	if payload.ID != id {
		t.Error("[X] TestGetPayloadByID:", ErrIDMismatch)
		return
	}

	// TODO: implement more checks
}

func TestPostPayloadsQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
