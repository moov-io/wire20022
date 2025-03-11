# wire20022
A go reader &amp; writer for supporting fedwire iso20022 

# Fedwire ISO20022 Go Wrapper

An implementation of the [Fedwire iso20022](https://www.frbservices.org/financial-services/fednow/what-is-iso-20022-why-does-it-matter)

## Project Overview

This project aims to create a Go wrapper for reading, writing, and validating Fedwire ISO20022 messages. The core functionality will build upon the existing structs generated from XSD schemas in the [moov-io/fewire20022](https://github.com/moov-io/fedwire20022)) project.

## Project Goals

1. Create a comprehensive wrapper in Go to handle ISO20022 Fedwire message formats
2. Implement robust validation to ensure all required fields are properly populated
3. Provide simple interfaces for reading and writing ISO20022 messages from/to files and byte streams
4. Ensure compliance with Fedwire ISO20022 specifications

## Technical Requirements

- Parse and generate valid ISO20022 XML for Fedwire messages
- Validate message structures against required field specifications
- Support error reporting with detailed field validation information
- Maintain type safety while working with complex nested structures
- Provide helper utilities to simplify message creation and inspection
- Flatten the complex XML structure to make it easier to interact with the underlying data
- Implement human-readable field naming that maps abbreviated XML names (e.g., "InsgrAgt") to their full descriptive names (e.g., "Instructed Agent")

## Implementation Approach

The project will use a modular architecture with separate packages for:

- Parsing (XML reading/writing)
- Validation (required field checking)
- Service layer (business logic)
- Helper utilities (structure introspection and documentation)

This separation will ensure maintainability and testability of the codebase.

## Target Use Cases

- Financial institutions needing to generate Fedwire messages in ISO20022 format
- Applications that need to validate and process incoming ISO20022 Fedwire messages
- Development tools for working with Fedwire ISO20022 messages
- Integration with existing financial systems that need to migrate to ISO20022 formats
