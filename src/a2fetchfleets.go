package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Page  Page    `json:"page"`
	Items []Fleet `json:"items"`
}

type Page struct {
	TotalItems int `json:"total_items"`
	ItemCount  int `json:"item_count"`
	PageSize   int `json:"page_size"`
	Page       int `json:"page"`
	Pages      int `json:"pages"`
}

type Fleet struct {
	FleetID   string    `json:"fleet_id"`
	FleetName string    `json:"fleet_name"`
	Created   time.Time `json:"created"`
	Stations  []Station `json:"stations"`
	Config    *string   `json:"config"` // nullable field
}

type Station struct {
	StationID           string    `json:"station_id"`
	FleetID             string    `json:"fleet_id"`
	SessionID           string    `json:"session_id"`
	StationName         string    `json:"station_name"`
	Region              string    `json:"region"`
	IP                  string    `json:"ip"`
	Version             string    `json:"version"`
	DeploymentCL        string    `json:"deployment_cl"`
	Created             time.Time `json:"created"`
	Online              bool      `json:"online"`
	LastEvent           time.Time `json:"last_event"`
	PlayerCount         int       `json:"player_count"`
	Disabled            bool      `json:"disabled"`
	Config              *string   `json:"config"`               // nullable field
	DistrictPopulations *string   `json:"district_populations"` // nullable field
}

func GetAllFleets() ([]Fleet, error) {
	apiKey := os.Getenv("A2_API_KEY")

	req, err := http.NewRequest("GET",
		"https://a2-station-api-prod-708695367983.us-central1.run.app/v2/fleets?include_config=false&include_stations=true&include_offline_fleets=false&page_size=16&page=1",
		nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var response Response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, fmt.Errorf("decoding JSON failure: %v", err)
	}

	return response.Items, nil
}
