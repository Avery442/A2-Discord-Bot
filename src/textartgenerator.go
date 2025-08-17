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

	// Collect top 3 stations by PlayerCount for each fleet
	for _, fleet := range fleets {
		// Sort stations by PlayerCount descending
		sort.Slice(fleet.Stations, func(i, j int) bool {
			return fleet.Stations[i].PlayerCount > fleet.Stations[j].PlayerCount
		})

		// Take up to top 3 stations
		limit := 3
		if len(fleet.Stations) < 3 {
			limit = len(fleet.Stations)
		}

		for i := 0; i < limit; i++ {
			station := fleet.Stations[i]
			rows = append(rows, Row{
				Name:        station.StationName,
				Version:     station.Version,
				PlayerCount: station.PlayerCount,
			})
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
	table += fmt.Sprintf("+-%s-+-%s-+----+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen))

	// Rows
	for _, row := range rows {
		table += fmt.Sprintf("| %-*s | %-*s | %-3d |\n",
			maxNameLen, row.Name,
			maxVersionLen, row.Version,
			row.PlayerCount)
	}

	// Bottom border
	table += fmt.Sprintf("+-%s-+-%s-+----+\n",
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
