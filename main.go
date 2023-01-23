package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/omeiirr/quran-obsidian-template/functions"
)

// Embedded file being passed from main.go
// Embedding inside modules not suported as of go v1.18.1

//go:embed data/ontology-concepts.json
var embeddedFile_concepts []byte

//go:embed data/asma-ul-husna.json
var embeddedFile_AllahNames []byte

func main() {
	// To measure execution time
	startTime := time.Now()

	// set path to working directory
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("Could not determine path of executable %s", err)
	}
	exPath := filepath.Dir(executable)
	workingDir := fmt.Sprintf("%s/Quranic Guidance/", exPath)
	fmt.Printf("workingDir: %s \n", workingDir)

	functions.GenerateFolderStructure(workingDir)
	functions.GenerateNamesOfAllah(workingDir, embeddedFile_AllahNames)
	functions.GenerateQuran(workingDir, 1, 114)
	functions.GenerateOntology(workingDir, embeddedFile_concepts)

	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\nAlhamdulillah! Operation completed in %s \nHappy learning.\n", elapsedTime)
}
