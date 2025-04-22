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
	"strings"

	"github.com/aknopov/xmlcomparator"
)

func WriteXMLTo(filePath string, data []byte) error {
    // Ensure directory exists with proper permissions
    if err := os.MkdirAll("generated", 0755); err != nil && !os.IsExist(err) {
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
    err := os.WriteFile(tempFile, data, 0644)
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
func removeAttributes(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)
	
	// Handle encoder closure (though not strictly required for bytes.Buffer)
	defer func() {
		if err := encoder.Close(); err != nil {
			// Log or handle residual errors if needed
			fmt.Printf("encoder close warning: %v", err)
		}
	}()

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decoder error: %w", err)
		}

		switch tok := t.(type) {
		case xml.StartElement:
			// Remove all attributes
			tok.Attr = nil
			if err := encoder.EncodeToken(tok); err != nil {
				return nil, fmt.Errorf("start element encode error: %w", err)
			}

		case xml.EndElement:
			if err := encoder.EncodeToken(tok); err != nil {
				return nil, fmt.Errorf("end element encode error: %w", err)
			}

		case xml.CharData, xml.Comment, xml.ProcInst, xml.Directive:
			if err := encoder.EncodeToken(tok); err != nil {
				return nil, fmt.Errorf("token encode error: %w", err)
			}

		default:
			return nil, fmt.Errorf("unhandled token type: %T", tok)
		}
	}

	// Final flush with error handling
	if err := encoder.Flush(); err != nil {
		return nil, fmt.Errorf("final flush error: %w", err)
	}

	return output.Bytes(), nil
}

// Matches ISO8601 date and datetime (very basic)
var dateRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var datetimeRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[T ]\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+-]\d{2}:\d{2})?$`)

func isDateOrDatetime(s string) bool {
	return dateRegex.MatchString(s) || datetimeRegex.MatchString(s)
}
func removeDateValues(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)

	// Ensure encoder closure (though Close() is not strictly needed for bytes.Buffer)
	defer func() {
		if closeErr := encoder.Close(); closeErr != nil {
			// Log if needed, but can't return error from defer
			fmt.Printf("encoder close warning: %v", closeErr)
		}
	}()

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decoder error: %w", err)
		}

		switch tok := t.(type) {
		case xml.StartElement:
			// Process attributes
			newAttrs := make([]xml.Attr, 0, len(tok.Attr))
			for _, attr := range tok.Attr {
				if isDateOrDatetime(attr.Value) {
					attr.Value = ""
				}
				newAttrs = append(newAttrs, attr)
			}
			tok.Attr = newAttrs
			
			if err := encoder.EncodeToken(tok); err != nil {
				return nil, fmt.Errorf("start element encode error: %w", err)
			}

		case xml.CharData:
			// Process text content
			content := strings.TrimSpace(string(tok))
			var encodedToken xml.Token = tok
			
			if isDateOrDatetime(content) {
				encodedToken = xml.CharData([]byte(""))
			}

			if err := encoder.EncodeToken(encodedToken); err != nil {
				return nil, fmt.Errorf("chardata encode error: %w", err)
			}

		case xml.EndElement:
			if err := encoder.EncodeToken(tok); err != nil {
				return nil, fmt.Errorf("end element encode error: %w", err)
			}

		default:
			if err := encoder.EncodeToken(tok); err != nil {
				return nil, fmt.Errorf("token encode error: %w", err)
			}
		}
	}

	// Final flush
	if err := encoder.Flush(); err != nil {
		return nil, fmt.Errorf("final flush error: %w", err)
	}

	return output.Bytes(), nil
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
