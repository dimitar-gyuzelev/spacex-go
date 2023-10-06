package v4tests

import (
	"fmt"
	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCoresReal(t *testing.T) {
	t.Skip("testing against the real SpaceXAPI")
	const base = "https://api.spacexdata.com/v4"
	cores := v4clients.NewCoresAPI(base)

	printMany[v4.Core](cores.GetCoresAll)
	printOne[v4.Core](cores.GetCoreByID, "5e9e289df35918033d3b2623")
}

func TestGetCoresAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(coresAll)))
	defer server.Close()

	client := v4clients.NewCoresAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	cores, err := client.GetCoresAll()
	if err != nil {
		t.Error("❌ TestGetCoresAll:", err)
		return
	}

	if len(cores) == 0 {
		t.Error("❌ TestGetCoresAll: cores were not retrieved")
		return
	}

	// TODO: improve test
}

func TestGetCoreByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler(coresOne)))
	defer server.Close()

	client := v4clients.NewCoresAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	id := "5e9e289df35918033d3b2623"

	core, err := client.GetCoreByID(id)
	if err != nil {
		t.Error("❌ TestGetCoreByID:", err)
		return
	}

	if core.ID != id {
		t.Error("❌ TestGetCoreByID: wrong Core ID")
		return
	}

	// TODO: improve test
}

func TestPostCoresQuery(t *testing.T) {
	t.Skip("PostCoresQuery is not implemented")

	server := httptest.NewServer(http.HandlerFunc(handler("")))
	defer server.Close()

	client := v4clients.NewCoresAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	_, err := client.PostCoresQuery(v4.Query{})
	if err != nil {
		t.Error("❌ TestPostCoresQuery:", err)
		return
	}
}
