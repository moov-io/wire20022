// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

// import (
// 	"bytes"
// 	"encoding/json"
// 	"encoding/xml"
// 	"errors"
// 	"fmt"
// 	"log"

// 	"io"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/wadearnold/wire20022/pkg/document"
// 	"github.com/wadearnold/wire20022/pkg/utils"
// )

// func outputError(w http.ResponseWriter, code int, err error) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(code)

// 	// Handle encoding error
// 	if encodeErr := json.NewEncoder(w).Encode(map[string]interface{}{
// 		"error": err.Error(),
// 	}); encodeErr != nil {
// 		// Critical failure handling
// 		http.Error(w, `{"error":"failed to encode error response"}`, http.StatusInternalServerError)
// 		log.Printf("JSON encoding failure: %v (Original error: %v)", encodeErr, err)
// 	}
// }

// func outputSuccess(w http.ResponseWriter, output string) {
// 	// Set headers first
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(http.StatusOK) // Must be called after headers, before body

// 	// Handle encoding error
// 	if err := json.NewEncoder(w).Encode(map[string]interface{}{
// 		"status": output,
// 	}); err != nil {
// 		// Critical fallback handling
// 		http.Error(w, `{"error":"failed to encode success response"}`, http.StatusInternalServerError)
// 		log.Printf("JSON encoding failure: %v (Output: %q)", err, output)
// 	}
// }

// func parseInputFromRequest(r *http.Request) (document.Iso20022Document, error) {
// 	inputFile, _, err := r.FormFile("input")
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Defer closure with error handling
// 	defer func() {
// 		closeErr := inputFile.Close()
// 		if closeErr != nil {
// 			// Combine original error (if any) with close error
// 			err = errors.Join(err, fmt.Errorf("close error: %w", closeErr))
// 		}
// 	}()

// 	var input bytes.Buffer
// 	if _, err = io.Copy(&input, inputFile); err != nil {
// 		return nil, err
// 	}

// 	return document.ParseIso20022Document(input.Bytes())
// }

// func messageToBuf(format utils.DocumentType, doc document.Iso20022Document) ([]byte, error) {
// 	var output []byte
// 	var err error
// 	switch format {
// 	case utils.DocumentTypeJson:
// 		output, err = json.MarshalIndent(doc, "", "\t")
// 	case utils.DocumentTypeXml:
// 		output, err = xml.MarshalIndent(doc, "", "\t")
// 	case utils.DocumentTypeUnknown:
// 		err = errors.New("unknown document type")
// 	}
// 	return output, err
// }

// func outputBufferToWriter(w http.ResponseWriter, doc document.Iso20022Document, format utils.DocumentType) {
// 	w.WriteHeader(http.StatusOK) // Should be set AFTER headers

// 	switch format {
// 	case utils.DocumentTypeJson:
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		if err := json.NewEncoder(w).Encode(doc); err != nil {
// 			log.Printf("JSON encoding error: %v", err)
// 			http.Error(w, "Internal server error", http.StatusInternalServerError)
// 			return
// 		}

// 	case utils.DocumentTypeXml:
// 		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
// 		encoder := xml.NewEncoder(w)
// 		if err := encoder.Encode(doc); err != nil {
// 			log.Printf("XML encoding error: %v", err)
// 			http.Error(w, "Internal server error", http.StatusInternalServerError)
// 			return
// 		}
// 		// Handle any buffered data
// 		if err := encoder.Close(); err != nil {
// 			log.Printf("XML encoder close error: %v", err)
// 			// Continue - best effort to complete the response
// 		}

// 	case utils.DocumentTypeUnknown:
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		if _, err := w.Write([]byte(`{"error": "invalid format"}`)); err != nil {
// 			log.Printf("Error writing response: %v", err)
// 		}
// 		return // No further processing needed

// 	default:
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		http.Error(w, `{"error": "unsupported format"}`, http.StatusBadRequest)
// 		return
// 	}
// }
// func getFormat(r *http.Request) (utils.DocumentType, error) {
// 	var format utils.DocumentType
// 	ff := r.FormValue("format")
// 	if ff == "" {
// 		format = utils.DocumentTypeXml
// 	} else {
// 		format = utils.DocumentType(ff)
// 	}
// 	if format != utils.DocumentTypeXml && format != utils.DocumentTypeJson {
// 		return format, fmt.Errorf("%s is an invalid format: %v", ff, format)
// 	}
// 	return format, nil
// }

// // validator - validate the file based on publication 1220
// func validator(w http.ResponseWriter, r *http.Request) {
// 	doc, err := parseInputFromRequest(r)
// 	if err != nil {
// 		outputError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	err = doc.Validate()
// 	if err != nil {
// 		outputError(w, http.StatusNotImplemented, err)
// 		return
// 	}

// 	outputSuccess(w, "valid file")
// }

// // validator - print file with ascii or json format
// func print(w http.ResponseWriter, r *http.Request) {
// 	doc, err := parseInputFromRequest(r)
// 	if err != nil {
// 		outputError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	format, err := getFormat(r)
// 	if err != nil {
// 		outputError(w, http.StatusNotImplemented, err)
// 		return
// 	}
// 	_, err = messageToBuf(format, doc)
// 	if err != nil {
// 		outputError(w, http.StatusNotImplemented, err)
// 		return
// 	}

// 	outputBufferToWriter(w, doc, format)
// }

// // convert - convert file with ascii or json format
// func convert(w http.ResponseWriter, r *http.Request) {
// 	message, err := parseInputFromRequest(r)
// 	if err != nil {
// 		outputError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	format, err := getFormat(r)
// 	if err != nil {
// 		outputError(w, http.StatusNotImplemented, err)
// 		return
// 	}

// 	output, err := messageToBuf(format, message)
// 	if err != nil {
// 		outputError(w, http.StatusNotImplemented, err)
// 		return
// 	}

// 	filename := "converted_file"
// 	w.Header().Set("Content-Type", "application/octet-stream")
// 	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
// 	w.Header().Set("Content-Transfer-Encoding", "binary")
// 	w.Header().Set("Expires", "0")
// 	w.WriteHeader(http.StatusOK)
// 	if _, err := w.Write(output); err != nil {
// 		log.Printf("Failed to write response: %v", err)
// 	}
// }

// // health - health check
// func health(w http.ResponseWriter, r *http.Request) {
// 	outputSuccess(w, "alive")
// }

// // configure handlers
// func ConfigureHandlers(r *mux.Router) error {
// 	r.HandleFunc("/health", health).Methods("GET")
// 	r.HandleFunc("/print", print).Methods("POST")
// 	r.HandleFunc("/validator", validator).Methods("POST")
// 	r.HandleFunc("/convert", convert).Methods("POST")
// 	return nil
// }
