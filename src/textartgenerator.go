package src

import (
	"fmt"
	"strconv"
)

type Row struct {
	Name        string
	Version     string
	PlayerCount int
}

func GenerateStationTable(fleets []Fleet) string {
	// Maintain the original 16-station limit
	var rows []Row

	// Collect all stations in original order
	for _, fleet := range fleets {
		for _, station := range fleet.Stations {
			if len(rows) >= 16 {
				break // Limit to 16 stations total
			}
			rows = append(rows, Row{
				Name:        station.StationName,
				Version:     station.Version,
				PlayerCount: station.PlayerCount,
			})
		}
		if len(rows) >= 16 {
			break
		}
	}

	if len(rows) == 0 {
		return ""
	}

	return generateTableFromRows(rows)
}

func generateTableFromRows(rows []Row) string {
	if len(rows) == 0 {
		return ""
	}

	// Determine column widths
	maxNameLen := 0
	maxVersionLen := 0
	maxPlayerLen := 0

	for _, row := range rows {
		if len(row.Name) > maxNameLen {
			maxNameLen = len(row.Name)
		}
		if len(row.Version) > maxVersionLen {
			maxVersionLen = len(row.Version)
		}
		playerLen := len(strconv.Itoa(row.PlayerCount))
		if playerLen > maxPlayerLen {
			maxPlayerLen = playerLen
		}
	}

	table := ""

	// Top border
	table += fmt.Sprintf("+-%s-+-%s-+-%s-+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen),
		repeat("-", maxPlayerLen))

	// Data rows
	for _, row := range rows {
		table += fmt.Sprintf("| %s | %s | %s |\n",
			center(row.Name, maxNameLen),
			center(row.Version, maxVersionLen),
			center(strconv.Itoa(row.PlayerCount), maxPlayerLen))
	}

	// Bottom border
	table += fmt.Sprintf("+-%s-+-%s-+-%s-+\n",
		repeat("-", maxNameLen),
		repeat("-", maxVersionLen),
		repeat("-", maxPlayerLen))

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

// helper function to center strings properly
func center(s string, width int) string {
	padding := width - len(s)
	if padding <= 0 {
		return s
	}
	left := padding / 2
	right := padding - left
	return fmt.Sprintf("%s%s%s", repeat(" ", left), s, repeat(" ", right))
}
