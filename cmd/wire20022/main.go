// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/moov-io/wire20022/pkg/messages"
)

// ValidationResult represents the result of validating a single file
type ValidationResult struct {
	File           string                 `json:"file"`
	Success        bool                   `json:"success"`
	MessageType    messages.MessageType   `json:"messageType"`
	Version        string                 `json:"version"`
	DetectionInfo  messages.DetectionInfo `json:"detectionInfo"`
	ValidationTime time.Duration          `json:"validationTime"`
	Error          string                 `json:"error,omitempty"`
	ErrorDetails   map[string]string      `json:"errorDetails,omitempty"`
}

// BatchResult represents the results of validating multiple files
type BatchResult struct {
	TotalFiles        int                `json:"totalFiles"`
	SuccessCount      int                `json:"successCount"`
	FailureCount      int                `json:"failureCount"`
	TotalTime         time.Duration      `json:"totalTime"`
	Results           []ValidationResult `json:"results"`
	MessageTypeCounts map[string]int     `json:"messageTypeCounts"`
}

var (
	verbose    bool
	jsonOutput bool
	recursive  bool
	pattern    string
	maxErrors  int
	showHelp   bool
	version    bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose error reporting with debugging details")
	flag.BoolVar(&verbose, "v", false, "Enable verbose error reporting (shorthand)")
	flag.BoolVar(&jsonOutput, "json", false, "Output results in JSON format")
	flag.BoolVar(&recursive, "recursive", false, "Recursively process directories")
	flag.BoolVar(&recursive, "r", false, "Recursively process directories (shorthand)")
	flag.StringVar(&pattern, "pattern", "*.xml", "File pattern to match (e.g., '*.xml', 'pacs.008*')")
	flag.IntVar(&maxErrors, "max-errors", 0, "Maximum number of errors to display (0 = all)")
	flag.BoolVar(&showHelp, "help", false, "Show help information")
	flag.BoolVar(&showHelp, "h", false, "Show help information (shorthand)")
	flag.BoolVar(&version, "version", false, "Show version information")
}

func main() {
	flag.Parse()

	if version {
		printVersion()
		os.Exit(0)
	}

	if showHelp || flag.NArg() == 0 {
		printHelp()
		if showHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	reader := messages.NewUniversalReader()

	// Process all arguments
	var allResults BatchResult
	allResults.MessageTypeCounts = make(map[string]int)
	startTime := time.Now()

	for _, arg := range flag.Args() {
		info, err := os.Stat(arg)
		if err != nil {
			printError(fmt.Sprintf("Cannot access %s: %v", arg, err))
			continue
		}

		if info.IsDir() {
			// Process directory
			results := processDirectory(reader, arg)
			allResults.Results = append(allResults.Results, results.Results...)
			allResults.TotalFiles += results.TotalFiles
			allResults.SuccessCount += results.SuccessCount
			allResults.FailureCount += results.FailureCount

			// Merge message type counts
			for msgType, count := range results.MessageTypeCounts {
				allResults.MessageTypeCounts[msgType] += count
			}
		} else {
			// Process single file
			result := processFile(reader, arg)
			allResults.Results = append(allResults.Results, result)
			allResults.TotalFiles++
			if result.Success {
				allResults.SuccessCount++
				allResults.MessageTypeCounts[string(result.MessageType)]++
			} else {
				allResults.FailureCount++
			}
		}
	}

	allResults.TotalTime = time.Since(startTime)

	// Output results
	if jsonOutput {
		outputJSON(allResults)
	} else {
		outputHuman(allResults)
	}

	// Exit with error code if any failures
	if allResults.FailureCount > 0 {
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Println("wire20022 - Fedwire ISO20022 Message Processing Tool")
	fmt.Println("Version: dev")
	fmt.Println("Built with Go")
}

func printHelp() {
	fmt.Println("wire20022 - Fedwire ISO20022 Message Processing Tool")
	fmt.Println("\nUsage: wire20022 [options] <file|directory> [<file|directory>...]")
	fmt.Println("\nThis tool automatically detects and validates Fedwire ISO20022 message files.")
	fmt.Println("It provides detailed error reporting to help debug parsing and validation issues.")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  wire20022 payment.xml                    # Validate single file")
	fmt.Println("  wire20022 -v payment.xml                 # Validate with verbose errors")
	fmt.Println("  wire20022 -json *.xml                    # Validate multiple files, output JSON")
	fmt.Println("  wire20022 -r messages/                   # Recursively validate directory")
	fmt.Println("  wire20022 -pattern 'pacs.008*' samples/  # Validate only pacs.008 files")
	fmt.Println("\nSupported Message Types:")
	fmt.Println("  - CustomerCreditTransfer (pacs.008)")
	fmt.Println("  - PaymentReturn (pacs.004)")
	fmt.Println("  - PaymentStatusRequest (pacs.028)")
	fmt.Println("  - FedwireFundsPaymentStatus (pacs.002)")
	fmt.Println("  - DrawdownRequest (pain.013)")
	fmt.Println("  - DrawdownResponse (pain.014)")
	fmt.Println("  - AccountReportingRequest (camt.060)")
	fmt.Println("  - ActivityReport (camt.052/086)")
	fmt.Println("  - EndpointDetailsReport (camt.052/090)")
	fmt.Println("  - EndpointGapReport (camt.052/087)")
	fmt.Println("  - EndpointTotalsReport (camt.052/089)")
	fmt.Println("  - ReturnRequestResponse (camt.029)")
	fmt.Println("  - ConnectionCheck (admi.001/004)")
	fmt.Println("  - FedwireFundsAcknowledgement (admi.004/007)")
	fmt.Println("  - FedwireFundsSystemResponse (admi.011)")
	fmt.Println("  - Master (camt.052)")
	fmt.Println("\nFuture Features:")
	fmt.Println("  - Message format conversion between versions")
	fmt.Println("  - HTTP server mode for message processing")
	fmt.Println("  - Message generation from templates")
}

func processFile(reader *messages.UniversalReader, filepath string) ValidationResult {
	startTime := time.Now()
	result := ValidationResult{
		File: filepath,
	}

	// Read file
	file, err := os.Open(filepath)
	if err != nil {
		result.Error = fmt.Sprintf("Failed to open file: %v", err)
		result.ValidationTime = time.Since(startTime)
		return result
	}
	defer file.Close()

	// Parse message
	parsed, err := reader.Read(file)
	if err != nil {
		result.Error = fmt.Sprintf("Failed to parse: %v", err)
		result.ValidationTime = time.Since(startTime)

		// Extract detailed error information if available
		if verbose {
			result.ErrorDetails = extractErrorDetails(err)
		}

		return result
	}

	// Fill in detection info
	result.MessageType = parsed.Type
	result.Version = parsed.Version
	result.DetectionInfo = parsed.Detection

	// Validate message
	err = reader.ValidateMessage(parsed)
	if err != nil {
		result.Error = fmt.Sprintf("Validation failed: %v", err)

		// Extract detailed validation error information
		if verbose {
			result.ErrorDetails = extractErrorDetails(err)
		}
	} else {
		result.Success = true
	}

	result.ValidationTime = time.Since(startTime)
	return result
}

func processDirectory(reader *messages.UniversalReader, dirPath string) BatchResult {
	result := BatchResult{
		MessageTypeCounts: make(map[string]int),
	}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			printError(fmt.Sprintf("Error accessing %s: %v", path, err))
			return nil
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file matches pattern
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err != nil || !matched {
			return nil
		}

		// Process file
		fileResult := processFile(reader, path)
		result.Results = append(result.Results, fileResult)
		result.TotalFiles++

		if fileResult.Success {
			result.SuccessCount++
			result.MessageTypeCounts[string(fileResult.MessageType)]++
		} else {
			result.FailureCount++
		}

		// Stop if we've hit max errors
		if maxErrors > 0 && result.FailureCount >= maxErrors {
			return filepath.SkipDir
		}

		return nil
	}

	if recursive {
		filepath.Walk(dirPath, walkFn)
	} else {
		// Only process files in the immediate directory
		files, err := os.ReadDir(dirPath)
		if err != nil {
			printError(fmt.Sprintf("Error reading directory %s: %v", dirPath, err))
			return result
		}

		for _, file := range files {
			if !file.IsDir() {
				walkFn(filepath.Join(dirPath, file.Name()), nil, nil)
			}
		}
	}

	return result
}

func extractErrorDetails(err error) map[string]string {
	details := make(map[string]string)

	// Extract different parts of the error
	errStr := err.Error()
	lines := strings.Split(errStr, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Look for specific patterns
		if strings.Contains(line, "Original error:") {
			details["original_error"] = strings.TrimPrefix(line, "Original error:")
		} else if strings.Contains(line, "Root element:") {
			details["root_element"] = strings.TrimPrefix(line, "Root element:")
		} else if strings.Contains(line, "Namespace:") {
			details["namespace"] = strings.TrimPrefix(line, "Namespace:")
		} else if strings.Contains(line, "Version:") {
			details["version"] = strings.TrimPrefix(line, "Version:")
		} else if strings.Contains(line, "Validation field:") {
			details["validation_field"] = strings.TrimPrefix(line, "Validation field:")
		} else if strings.Contains(line, "Validation reason:") {
			details["validation_reason"] = strings.TrimPrefix(line, "Validation reason:")
		} else if strings.Contains(line, "Detection method:") {
			details["detection_method"] = strings.TrimPrefix(line, "Detection method:")
		}
	}

	return details
}

func outputJSON(result BatchResult) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(result)
}

func outputHuman(result BatchResult) {
	// Summary header
	fmt.Printf("\nValidation Summary\n")
	fmt.Printf("==================\n")
	fmt.Printf("Total files processed: %d\n", result.TotalFiles)
	fmt.Printf("Successful: %d\n", result.SuccessCount)
	fmt.Printf("Failed: %d\n", result.FailureCount)
	fmt.Printf("Total time: %s\n", result.TotalTime)

	// Message type breakdown
	if len(result.MessageTypeCounts) > 0 {
		fmt.Printf("\nMessage Types Found:\n")
		for msgType, count := range result.MessageTypeCounts {
			fmt.Printf("  %s: %d\n", msgType, count)
		}
	}

	// Failed validations
	failedCount := 0
	if result.FailureCount > 0 {
		fmt.Printf("\nFailed Validations:\n")
		fmt.Printf("-------------------\n")

		for _, r := range result.Results {
			if !r.Success {
				failedCount++
				fmt.Printf("\n[%d] File: %s\n", failedCount, r.File)
				fmt.Printf("    Error: %s\n", r.Error)

				if verbose && len(r.ErrorDetails) > 0 {
					fmt.Printf("    Details:\n")
					for key, value := range r.ErrorDetails {
						fmt.Printf("      %s: %s\n", key, strings.TrimSpace(value))
					}
				}

				if r.MessageType != "" {
					fmt.Printf("    Detected Type: %s\n", r.MessageType)
				}
				if r.Version != "" {
					fmt.Printf("    Version: %s\n", r.Version)
				}
				if r.DetectionInfo.DetectedBy != "" {
					fmt.Printf("    Detection Method: %s\n", r.DetectionInfo.DetectedBy)
				}

				// Stop if we've shown enough errors
				if maxErrors > 0 && failedCount >= maxErrors {
					remaining := result.FailureCount - failedCount
					if remaining > 0 {
						fmt.Printf("\n... and %d more errors (use -max-errors=0 to see all)\n", remaining)
					}
					break
				}
			}
		}
	}

	// Success indicator
	if result.FailureCount == 0 && result.TotalFiles > 0 {
		fmt.Printf("\n✓ All validations passed!\n")
	}
}

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
}
