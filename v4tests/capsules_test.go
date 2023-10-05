package v4tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	v4 "github.com/dimitar-gyuzelev/spacex-go/v4"
	"github.com/dimitar-gyuzelev/spacex-go/v4clients"
)

func TestGetCapsulesAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeGetCapsules))
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
	server := httptest.NewServer(http.HandlerFunc(fakeGetCapsuleByID))
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

	server := httptest.NewServer(http.HandlerFunc(fakePostCapsulesQuery))
	defer server.Close()

	client := v4clients.NewCapsulesAPI(fmt.Sprintf("%s/%s", server.URL, "v4"))

	_, err := client.PostCapsulesQuery(v4.Query{})
	if err != nil {
		t.Error("❌ TestPostCapsulesQuery:", err)
		return
	}
}

func fakeGetCapsules(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := []v4.Capsule{
		{
			ReuseCount:    0,
			WaterLandings: 1,
			LandLandings:  0,
			LastUpdate:    "Hanging in atrium at SpaceX HQ in Hawthorne ",
			Launches:      []string{"5eb87cdeffd86e000604b330"},
			Serial:        "C101",
			Status:        "retired",
			Type:          "Dragon 1.0",
			ID:            "5e9e2c5bf35918ed873b2664",
		},
		{
			ReuseCount:    0,
			WaterLandings: 1,
			LandLandings:  0,
			LastUpdate:    "On display at KSC Visitor's Center ",
			Launches:      []string{"5eb87cdfffd86e000604b331"},
			Serial:        "C102",
			Status:        "retired",
			Type:          "Dragon 1.0",
			ID:            "5e9e2c5bf3591882af3b2665",
		},
	}

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func fakeGetCapsuleByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := v4.Capsule{
		ReuseCount:    0,
		WaterLandings: 1,
		LandLandings:  0,
		LastUpdate:    "Hanging in atrium at SpaceX HQ in Hawthorne ",
		Launches:      []string{"5eb87cdeffd86e000604b330"},
		Serial:        "C101",
		Status:        "retired",
		Type:          "Dragon 1.0",
		ID:            "5e9e2c5bf35918ed873b2664",
	}

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func fakePostCapsulesQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO: implement
	w.WriteHeader(http.StatusServiceUnavailable)
}
