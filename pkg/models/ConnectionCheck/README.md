# ConnectionCheck

The `ConnectionCheck` package is part of the [`moov-io/wire20022`](https://github.com/moov-io/wire20022) library. It provides functionality for handling ISO 20022 Connection Check messages (`admi.004`) across multiple versions. This package includes tools for creating, validating, and converting between XML documents and Go data models.


## Features

- **Message Model**: Defines the `MessageModel` struct for representing account reporting request data.
- **Namespace Mapping**: Supports multiple versions of `admi.004` messages using `NameSpaceModelMap`.
- **Validation**: Ensures required fields are present and valid.
- **XML Conversion**: Converts between XML documents and Go models.
- **Version Support**: Handles versions `admi.004.001.01` through `admi.004.001.02`.


## Installation

To use this package in your Go project:

```bash
go get github.com/moov-io/wire20022/pkg/ConnectionCheck
```


## Usage

### Create a Document from a Model

You can create an XML document from a `MessageModel` using the `DocumentWith` function.

```go
    // Define a sample MessageModel
    model := ConnectionCheck.MessageModel{
        EventType: "PING",
        EventTime:    time.Now(),
        EventParam:    "BMQFMI01",
    }

    // Create a document from the model
    doc, err := DocumentWith(model, ConnectionCheck.ADMI_004_001_02)
    if err != nil {
        log.Fatal(err)
    }
```

### Validate a Document

You can validate the structure and required fields of a document using the `Validate` method.

```go
if err := doc.Validate(); err != nil {
    log.Fatal("Validation failed:", err)
}
```


### Convert XML to a Model

You can convert a raw XML document back into a `MessageModel` using the `ParseXML` function.

```go
model, err := ParseXML(xmlBytes)
if err != nil {
    log.Fatal("Failed to parse XML:", err)
}
```

### Check Required Fields

You can use the `CheckRequiredFields` function to verify that all required fields are present in the model.

```go
if err := CheckRequiredFields(model); err != nil {
    log.Fatal("Missing required fields:", err)
}
```


## Supported Versions

The package supports the following versions of `admi.004`:

- `admi.004.001.01`
- `admi.004.001.02`


## Testing

The package includes comprehensive tests for all supported versions.

To run the tests:

```bash
go test ./...
```


### Example test cases include:

- Creating documents from models
- Validating documents
- Converting XML to models and back
- Checking required fields


## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Write tests for your changes.
4. Submit a pull request.


## License

This project is licensed under the [Apache 2.0 License](LICENSE).


## Contact

For questions or support, please [open an issue](https://github.com/moov-io/wire20022/issues) on the GitHub repository.
