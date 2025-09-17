package main

import (
	"bufio"
	"fmt"
	"logs-exercise/analyze"
	"os"
	"time"
)

func main() {

	start := time.Now()
	fmt.Printf("started!")

	const inputPath = "logs.txt"

	lines, err := readLines(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "errore lettura %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	// Parametri base per l'esercizio (puoi modificarli a piacere).
	levels := []string{"ERR", "WARN"}
	patterns := []string{"CALL GET /orders", "CALL GET /customers"}

	// TODO: implementa tu AnalyzeLogs in un altro file dello stesso package.
	lvlCounts, patCounts, total, err := analyze.AnalyzeLogs(lines, levels, patterns /*concurrency*/, 4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "errore analisi: %v\n", err)
		os.Exit(1)
	}

	// Output semplice
	fmt.Printf("righe_totali=%d\n", total)
	fmt.Println("Conteggio livelli:")
	for _, lv := range levels {
		fmt.Printf("  %s: %d\n", lv, lvlCounts[lv])
	}
	fmt.Println("Conteggio pattern:")
	for _, p := range patterns {
		fmt.Printf("  %s: %d\n", p, patCounts[p])
	}

	fmt.Printf("Tempo totale: %v\n", time.Since(start))
}

// readLines legge il file e introduce 0.5s di latenza per ogni riga letta.
// Nota: richiede `import "time"`.
func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		out = append(out, sc.Text())
		time.Sleep(500 * time.Millisecond) // latenza per riga
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
