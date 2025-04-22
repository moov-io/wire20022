package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moov-io/wire20022/pkg/document"
)

func main() {
	xsdFileName := "Fedwire_Funds_Service_Release_2025_AccountReportingRequest_camt_060_001_05_20241122_1718_iso15"
	xsdPath := filepath.Join("xsd", xsdFileName+".xsd")

	// Read XSD file
	xsdContent, err := os.ReadFile(xsdPath)
	if err != nil {
		fmt.Printf("error reading XSD file: %v\n", err)
		os.Exit(1)
	}

	// Parse document
	doc, err := document.ParseIso20022Document(xsdContent)
	if err != nil {
		fmt.Printf("error parsing XSD file: %v\n", err)
		os.Exit(1)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Printf("error marshaling to JSON: %v\n", err)
		os.Exit(1)
	}

	// Write JSON file
	jsonFileName := filepath.Join("json", xsdFileName+".json")
	err = os.WriteFile(jsonFileName, jsonData, 0644)
	if err != nil {
		fmt.Printf("error writing JSON file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully converted XSD to JSON at:", jsonFileName)
}
