package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/omeiirr/quran-obsidian-template/functions"
)

// Embedded file being passed from main.go
// Embedding inside modules not suported as of go v1.18.1

//go:embed data/asma-ul-husna.json
var embedded_file []byte

func main() {
	// To measure execution time
	startTime := time.Now()

	functions.GenerateFolderStructure()
	functions.GenerateQuran(1, 114)

	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\nAlhamdulillah! Operation completed in %s \nHappy learning.\n", elapsedTime)
}
