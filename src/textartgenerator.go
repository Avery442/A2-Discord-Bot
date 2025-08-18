package src

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// FleetScore holds a fleet and its top station scores
type FleetScore struct {
	Fleet       Fleet
	Score       int
	TopStations []Station
}

func GenerateStationTable(fleets []Fleet) string {
	type Row struct {
		Name        string
		Version     string
		PlayerCount int
	}

	var rows []Row
	var fleetScores []FleetScore

	// Collect top stations and calculate fleet scores
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
	maxNameLen, maxVersionLen, maxPlayerLen := 0, 0, 0
	for _, row := range rows {
		if len(row.Name) > maxNameLen {
			maxNameLen = len(row.Name)
		}
		if len(row.Version) > maxVersionLen {
			maxVersionLen = len(row.Version)
		}
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

	// Rows
	for _, row := range rows {
		buf.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
			center(row.Name, maxNameLen),
			center(row.Version, maxVersionLen),
			center(strconv.Itoa(row.PlayerCount), maxPlayerLen)))
	}

	writeBorder() // bottom border

	return buf.String()
}

// center adds spaces to both sides to center the string in a cell
func center(s string, width int) string {
	padding := width - len(s)
	if padding <= 0 {
		return s
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
