package src

import (
	"fmt"
)

func GenerateStationTable(fleets []Fleet) string {
	type Row struct {
		Name        string
		Version     string
		PlayerCount int
	}

	var rows []Row

	// Collect all stations in original order
	for _, fleet := range fleets {
		for _, station := range fleet.Stations {
			rows = append(rows, Row{
				Name:        station.StationName,
				Version:     station.Version,
				PlayerCount: station.PlayerCount,
			})
		}
	}

	// Limit to 16 stations total
	if len(rows) > 16 {
		rows = rows[:16]
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
