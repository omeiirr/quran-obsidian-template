package main

import (
	"fmt"
	"time"

	"github.com/omeiirr/quran-obsidian-template/functions"
)

func main() {

	// To measure execution time
	startTime := time.Now()

	functions.GenerateFolderStructure()
	functions.GenerateNamesOfAllah()
	functions.GenerateQuran()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation completed in %s", elapsedTime)
}
