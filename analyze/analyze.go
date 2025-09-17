package analyze

import "strings"

func AnalyzeLogs(lines []string, levels, patterns []string, concurrency int) (map[string]int, map[string]int, int, error) {
	lvlCounts := make(map[string]int, len(levels))
	pattCounts := make(map[string]int, len(patterns))

	records := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		records++
		for _, level := range levels {
			if strings.Contains(line, level) {
				lvlCounts[level]++
			}
		}
		for _, pattern := range patterns {
			if strings.Contains(line, pattern) {
				pattCounts[pattern]++
			}
		}
	}
	return lvlCounts, pattCounts, records, nil
}
