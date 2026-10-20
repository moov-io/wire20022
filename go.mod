module github.com/moov-io/wire20022

go 1.25.8

toolchain go1.27.3

require (
	cloud.google.com/go v0.123.0
	// Tagged fedwire20022 releases dropped gen/<MessageType>/<version> packages this module still imports.
	github.com/moov-io/fedwire20022 v0.0.0-20250827223334-b9613060d2a2
	github.com/stretchr/testify v1.12.1
)

require (
	github.com/moov-io/base v0.62.1 // indirect
	github.com/rickar/cal/v2 v2.1.28 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)
