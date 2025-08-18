package src

import (
	"fmt"
	"sort"
)

func GenerateStationTable(fleets []Fleet) string {
	type Row struct {
		Name        string
		Version     string
		PlayerCount int
	}

	var rows []Row

	type FleetScore struct {
		Fleet       Fleet
		Score       int
		TopStations []Station
	}

	var fleetScores []FleetScore

	// Calculate fleet scores (sum of top 3 stations)
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
	}

	// Determine column widths
	maxNameLen := 0
	maxVersionLen := 0
	for _, row := range rows {
		if len(row.Name) > maxNameLen {
			maxNameLen = len(row.Name)
		}
		if len(row.Version) > maxVersionLen {
			maxVersionLen = len(row.Version)
		}
	}

	table := ""
	// Top border
	table += fmt.Sprintf("+-%s-+-%s-+------+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen))

	// Rows
	for _, row := range rows {
		table += fmt.Sprintf("| %s | %s | %s |\n",
			center(row.Name, maxNameLen),
			center(row.Version, maxVersionLen),
			center(fmt.Sprintf("%d", row.PlayerCount), 4))
	}

	// Bottom border
	table += fmt.Sprintf("+-%s-+-%s-+------+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen))

	return table
}

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
	}
	padding := width - len(text)
	left := padding / 2
	right := padding - left
	return repeat(" ", left) + text + repeat(" ", right)
}
