package v4tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestCoresReal(t *testing.T) {
	t.Skip(MsgTestingRealAPI)

	cores := v4clients.NewCoresAPI(spacexLiveAPIBase)

	printMany[v4.Core](cores.GetCoresAll)
	printOne[v4.Core](cores.GetCoreByID, "5e9e289df35918033d3b2623")
}

func TestGetCoresAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(coresAll)))
	defer server.Close()

	client := v4clients.NewCoresAPI(server.URL)

	cores, err := client.GetCoresAll()
	if err != nil {
		t.Error("[X] TestGetCoresAll:", err)
		return
	}

	if len(cores) != 2 {
		t.Error("[X] TestGetCoresAll:", ErrCollectionLenMismatch)
		return
	}

	// TODO: improve test
}

func TestGetCoreByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(coresOne)))
	defer server.Close()

	client := v4clients.NewCoresAPI(server.URL)

	id := "5e9e289df35918033d3b2623"

	core, err := client.GetCoreByID(id)
	if err != nil {
		t.Error("[X] TestGetCoreByID:", err)
		return
	}

	if core.ID != id {
		t.Error("[X] TestGetCoreByID:", ErrIDMismatch)
		return
	}

	// TODO: improve test
}

func TestPostCoresQuery(t *testing.T) {
	// TODO: implement
	t.Skip(MsgNotImplemented)
}
