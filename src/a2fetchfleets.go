package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	return getFleets(1, 16) // Get first page with 16 items (original behavior)
}

func GetAllFleetsAllPages() ([]Fleet, error) {
	return getFleets(1, 100) // Start with larger page size and get all pages
}

func getFleets(page int, pageSize int) ([]Fleet, error) {
	baseURL := "https://a2-station-api-prod-708695367983.us-central1.run.app/v2/fleets"
	url := fmt.Sprintf("%s?include_config=false&include_stations=true&include_offline_fleets=false&page_size=%d&page=%d", baseURL, pageSize, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request in a2fetchfleets.go:", err)
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("x-api-key", os.Getenv("A2_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, fmt.Errorf("There was an error sending the request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return nil, fmt.Errorf("There was an error reading the body: %v", err)

	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("decoding JSON failure: %v", err)
	}

	allFleets := response.Items

	// If this is the bulk fetch and there are more pages, get them all
	if pageSize > 16 && response.Page.Pages > 1 {
		for currentPage := 2; currentPage <= response.Page.Pages; currentPage++ {
			additionalFleets, err := getFleets(currentPage, pageSize)
			if err != nil {
				fmt.Printf("Error fetching page %d: %v\n", currentPage, err)
				continue // Continue with other pages even if one fails
			}
			allFleets = append(allFleets, additionalFleets...)
		}
	}

	return allFleets, nil
}
