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

		// (facoltativo) per matching case-insensitive:
		// l := strings.ToLower(line)

		for _, level := range levels {
			// if strings.Contains(l, strings.ToLower(level)) { // versione insensitive
			if strings.Contains(line, level) { // versione attuale
				lvlCounts[level]++
			}
		}
		for _, pattern := range patterns {
			// if strings.Contains(l, strings.ToLower(pattern)) {
			if strings.Contains(line, pattern) {
				pattCounts[pattern]++
			}
		}
	}
	return lvlCounts, pattCounts, records, nil
}
