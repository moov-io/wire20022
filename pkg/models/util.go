package model

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/aknopov/xmlcomparator"
)

func WriteXMLTo(filePath string, data []byte) error {
	// Ensure directory exists with proper permissions
	if err := os.MkdirAll("generated", 0750); err != nil && !os.IsExist(err) {
		return fmt.Errorf("directory creation failed: %w", err)
	}

	// Construct full file path
	xmlFileName := filepath.Join("generated", filePath)

	// Validate file extension
	if ext := filepath.Ext(xmlFileName); ext != ".xml" {
		return fmt.Errorf("invalid file extension %q, must be .xml", ext)
	}

	// Write file with atomic replacement
	tempFile := xmlFileName + ".tmp"
	err := os.WriteFile(tempFile, data, 0600)
	if err != nil {
		return fmt.Errorf("temporary file write failed: %w", err)
	}

	// Atomic rename for crash safety
	if err := os.Rename(tempFile, xmlFileName); err != nil {
		// Clean up temp file if rename fails
		if err := os.Remove(tempFile); err != nil && !os.IsNotExist(err) {
			log.Printf("failed to remove temp file %q: %v", tempFile, err)
		}
		return fmt.Errorf("file rename failed: %w", err)
	}

	return nil
}
func ReadXMLFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return data, nil
}

//nolint:gocyclo // date parser is verbose by nature and acceptable here
func removeAttributes(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)
	//nolint:errcheck
	defer encoder.Close()
	for {
		tok, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode token: %w", err)
		}

		switch t := tok.(type) {
		case xml.StartElement:
			t.Attr = nil // Strip all attributes
			if err := encoder.EncodeToken(t); err != nil {
				return nil, fmt.Errorf("encode start element: %w", err)
			}
		case xml.EndElement, xml.CharData:
			if err := encoder.EncodeToken(t); err != nil {
				return nil, fmt.Errorf("encode token: %w", err)
			}
		}
	}

	if err := encoder.Flush(); err != nil {
		return nil, fmt.Errorf("flush encoder: %w", err)
	}

	return buf.Bytes(), nil
}

// Matches ISO8601 date and datetime (very basic)
var dateOrDateTimeRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}([T\s]\d{2}:\d{2}:\d{2}([+-]\d{2}:\d{2}|Z)?)?$`)

//nolint:gocyclo // date parser is verbose by nature and acceptable here
func removeDateValues(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)
	//nolint:errcheck
	defer encoder.Close()
	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decode token: %w", err)
		}

		switch t := tok.(type) {
		case xml.StartElement, xml.EndElement:
			if err := encoder.EncodeToken(t); err != nil {
				return nil, fmt.Errorf("encode element: %w", err)
			}
		case xml.CharData:
			text := string(t)
			if dateOrDateTimeRegex.MatchString(text) {
				t = []byte("") // Clear value if it's a date or datetime
			}
			if err := encoder.EncodeToken(t); err != nil {
				return nil, fmt.Errorf("encode chardata: %w", err)
			}
		}
	}

	if err := encoder.Flush(); err != nil {
		return nil, fmt.Errorf("flush encoder: %w", err)
	}

	return buf.Bytes(), nil
}
func CompareXMLs(filePath1 string, filePath2 string) bool {
	xml1, err := ReadXMLFile(filePath1)
	if err != nil {
		return false
	}
	xml2, err := ReadXMLFile(filePath2)
	if err != nil {
		return false
	}
	clean1, err := removeAttributes(xml1)
	if err != nil {
		log.Fatal(err)
	}
	clean2, err := removeAttributes(xml2)
	if err != nil {
		log.Fatal(err)
	}
	clean1, err = removeDateValues(clean1)
	if err != nil {
		log.Fatal(err)
	}
	clean2, err = removeDateValues(clean2)
	if err != nil {
		log.Fatal(err)
	}
	diffs := xmlcomparator.CompareXmlStrings(string(clean1), string(clean2), true)
	if len(diffs) == 0 {
		log.Println("XML files are equal (ignoring date/datetime fields).")
		return true
	} else {
		log.Println("Differences found (ignoring date/datetime fields):")
		for _, diff := range diffs {
			log.Println(diff)
		}
		return false
	}
}
