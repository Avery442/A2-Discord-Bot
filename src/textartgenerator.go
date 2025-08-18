package src

import (
	"bytes"
	"fmt"
<<<<<<< Updated upstream
	"sort"
=======
	"strconv"
	"strings"
>>>>>>> Stashed changes
)

func GenerateStationTable(fleets []Fleet) string {
	type Row struct {
		Name        string
		Version     string
		PlayerCount int
	}

	var rows []Row

<<<<<<< Updated upstream
	type FleetScore struct {
		Fleet       Fleet
		Score       int
		TopStations []Station
	}

	var fleetScores []FleetScore

	// Calculate fleet scores (sum of top 3 stations)
=======
	// Collect all stations
>>>>>>> Stashed changes
	for _, fleet := range fleets {
		// Sort stations by PlayerCount descending
		sort.Slice(fleet.Stations, func(i, j int) bool {
			return fleet.Stations[i].PlayerCount > fleet.Stations[j].PlayerCount
		})

		limit := 3
		if len(fleet.Stations) < 3 {
			limit = len(fleet.Stations)
		}

		topStations := fleet.Stations[:limit]

		score := 0
		for _, s := range topStations {
			score += s.PlayerCount
		}

		fleetScores = append(fleetScores, FleetScore{
			Fleet:       fleet,
			Score:       score,
			TopStations: topStations,
		})
	}

	// Sort fleets by total top 3 player counts descending
	sort.Slice(fleetScores, func(i, j int) bool {
		return fleetScores[i].Score > fleetScores[j].Score
	})

	// Fill rows with top stations, respecting the 16-station limit
	totalStations := 0
	for _, fs := range fleetScores {
		for _, station := range fs.TopStations {
			if totalStations >= 16 {
				break
			}
			rows = append(rows, Row{
				Name:        station.StationName,
				Version:     station.Version,
				PlayerCount: station.PlayerCount,
			})
			totalStations++
		}
		if totalStations >= 16 {
			break
		}
<<<<<<< Updated upstream
	}

	// Determine column widths
	maxNameLen := 0
	maxVersionLen := 0
=======
	}

	if len(rows) > 16 {
		rows = rows[:16]
	}

	// Determine column widths
	maxNameLen, maxVersionLen, maxPlayerLen := 0, 0, 0
>>>>>>> Stashed changes
	for _, row := range rows {
		if len(row.Name) > maxNameLen {
			maxNameLen = len(row.Name)
		}
		if len(row.Version) > maxVersionLen {
			maxVersionLen = len(row.Version)
		}
<<<<<<< Updated upstream
	}

	table := ""
	// Top border
	table += fmt.Sprintf("+-%s-+-%s-+------+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen))
=======
		if l := len(strconv.Itoa(row.PlayerCount)); l > maxPlayerLen {
			maxPlayerLen = l
		}
	}

	var buf bytes.Buffer

	// Function to write a border line
	writeBorder := func() {
		buf.WriteString(fmt.Sprintf("+-%s-+-%s-+-%s-+\n",
			strings.Repeat("-", maxNameLen),
			strings.Repeat("-", maxVersionLen),
			strings.Repeat("-", maxPlayerLen)))
	}

	writeBorder() // top border
>>>>>>> Stashed changes

	// Rows
	for _, row := range rows {
		buf.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
			center(row.Name, maxNameLen),
			center(row.Version, maxVersionLen),
<<<<<<< Updated upstream
			center(fmt.Sprintf("%d", row.PlayerCount), 4))
	}

	// Bottom border
	table += fmt.Sprintf("+-%s-+-%s-+------+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen))
=======
			center(strconv.Itoa(row.PlayerCount), maxPlayerLen)))
	}

	writeBorder() // bottom border
>>>>>>> Stashed changes

	return buf.String()
}

<<<<<<< Updated upstream
// helper function to repeat strings
func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// helper function to center text
func center(text string, width int) string {
	if len(text) >= width {
		return text
=======
// center adds spaces to both sides to center the string in a cell
func center(s string, width int) string {
	padding := width - len(s)
	if padding <= 0 {
		return s
>>>>>>> Stashed changes
	}
	padding := width - len(text)
	left := padding / 2
	right := padding - left
<<<<<<< Updated upstream
	return repeat(" ", left) + text + repeat(" ", right)
=======
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
>>>>>>> Stashed changes
}
