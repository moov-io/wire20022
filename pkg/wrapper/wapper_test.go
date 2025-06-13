package wrapper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	AccountReportingRequest "github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	ActivityReport "github.com/moov-io/wire20022/pkg/models/ActivityReport"
	ConnectionCheck "github.com/moov-io/wire20022/pkg/models/ConnectionCheck"
	CustomerCreditTransfer "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
	DrawdownRequest "github.com/moov-io/wire20022/pkg/models/DrawdownRequest"
	DrawdownResponse "github.com/moov-io/wire20022/pkg/models/DrawdownResponse"
	EndpointDetailsReport "github.com/moov-io/wire20022/pkg/models/EndpointDetailsReport"
	EndpointGapReport "github.com/moov-io/wire20022/pkg/models/EndpointGapReport"
	EndpointTotalsReport "github.com/moov-io/wire20022/pkg/models/EndpointTotalsReport"
	FedwireFundsAcknowledgement "github.com/moov-io/wire20022/pkg/models/FedwireFundsAcknowledgement"
	FedwireFundsPaymentStatus "github.com/moov-io/wire20022/pkg/models/FedwireFundsPaymentStatus"
	FedwireFundsSystemResponse "github.com/moov-io/wire20022/pkg/models/FedwireFundsSystemResponse"
	Master "github.com/moov-io/wire20022/pkg/models/Master"
	PaymentReturn "github.com/moov-io/wire20022/pkg/models/PaymentReturn"
	PaymentStatusRequest "github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest"
	ReturnRequestResponse "github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse"
)

func getSwiftSampleFiles(model string) map[string][]string {
    // Dynamically determine the base path relative to the current file
    basePath, err := filepath.Abs(filepath.Join("..", "models", model, "swiftSample"))
    if err != nil {
        panic(err)
    }

    files := make(map[string][]string)

    err = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            // Remove the base path from the file path
            relativePath := strings.TrimPrefix(path, basePath+string(os.PathSeparator))
            files[model] = append(files[model], relativePath)
        }
        return nil
    })

    if err != nil {
        panic(err)
    }

    return files
}

func TestAccountReportingRequest(t *testing.T) {
    wrapper := &AccountReportingRequestWrapper{}
    files := getSwiftSampleFiles("AccountReportingRequest")
    for _, file := range files["AccountReportingRequest"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "AccountReportingRequest", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []AccountReportingRequest.CAMT_060_001_VESION{
                AccountReportingRequest.CAMT_060_001_02,
                AccountReportingRequest.CAMT_060_001_02,
                AccountReportingRequest.CAMT_060_001_02,
                AccountReportingRequest.CAMT_060_001_02,
                AccountReportingRequest.CAMT_060_001_02,
                AccountReportingRequest.CAMT_060_001_02,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestActivityReport(t *testing.T) {
    wrapper := &ActivityReportWrapper{}
    files := getSwiftSampleFiles("ActivityReport")
    for _, file := range files["ActivityReport"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "ActivityReport", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []ActivityReport.CAMT_052_001_VESION{
                ActivityReport.CAMT_052_001_01,
                ActivityReport.CAMT_052_001_02,
                ActivityReport.CAMT_052_001_03,
                ActivityReport.CAMT_052_001_04,
                ActivityReport.CAMT_052_001_05,
                ActivityReport.CAMT_052_001_06,
                ActivityReport.CAMT_052_001_07,
                ActivityReport.CAMT_052_001_08,
                ActivityReport.CAMT_052_001_09,
                ActivityReport.CAMT_052_001_10,
                ActivityReport.CAMT_052_001_11,
                ActivityReport.CAMT_052_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestConnectionCheck(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}
    files := getSwiftSampleFiles("ConnectionCheck")
    for _, file := range files["ConnectionCheck"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "ConnectionCheck", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []ConnectionCheck.ADMI_004_001_VESION{
                ConnectionCheck.ADMI_004_001_01,
                ConnectionCheck.ADMI_004_001_01,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestCustomerCreditTransfer(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}
    files := getSwiftSampleFiles("CustomerCreditTransfer")
    for _, file := range files["CustomerCreditTransfer"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "CustomerCreditTransfer", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []CustomerCreditTransfer.PACS_008_001_VESION{
                CustomerCreditTransfer.PACS_008_001_02,
                CustomerCreditTransfer.PACS_008_001_03,
				CustomerCreditTransfer.PACS_008_001_04,
				CustomerCreditTransfer.PACS_008_001_05,
				CustomerCreditTransfer.PACS_008_001_06,
				CustomerCreditTransfer.PACS_008_001_07,
				CustomerCreditTransfer.PACS_008_001_08,
				CustomerCreditTransfer.PACS_008_001_09,
				CustomerCreditTransfer.PACS_008_001_10,
				CustomerCreditTransfer.PACS_008_001_11,
				CustomerCreditTransfer.PACS_008_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}

func TestDrawdownRequest(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}
    files := getSwiftSampleFiles("DrawdownRequest")
    for _, file := range files["DrawdownRequest"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "DrawdownRequest", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []DrawdownRequest.PAIN_013_001_VESION{
                DrawdownRequest.PAIN_013_001_01,
                DrawdownRequest.PAIN_013_001_02,
                DrawdownRequest.PAIN_013_001_03,
                DrawdownRequest.PAIN_013_001_04,
                DrawdownRequest.PAIN_013_001_05,
                DrawdownRequest.PAIN_013_001_06,
                DrawdownRequest.PAIN_013_001_07,
                DrawdownRequest.PAIN_013_001_08,
                DrawdownRequest.PAIN_013_001_09,
                DrawdownRequest.PAIN_013_001_10,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestDrawdownResponse(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}
    files := getSwiftSampleFiles("DrawdownResponse")
    for _, file := range files["DrawdownResponse"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "DrawdownResponse", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []DrawdownResponse.PAIN_014_001_VESION{
                DrawdownResponse.PAIN_014_001_01,
                DrawdownResponse.PAIN_014_001_02,
                DrawdownResponse.PAIN_014_001_03,
                DrawdownResponse.PAIN_014_001_04,
                DrawdownResponse.PAIN_014_001_05,
                DrawdownResponse.PAIN_014_001_06,
                DrawdownResponse.PAIN_014_001_07,
                DrawdownResponse.PAIN_014_001_08,
                DrawdownResponse.PAIN_014_001_09,
                DrawdownResponse.PAIN_014_001_10,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestEndpointDetailsReport(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}
    files := getSwiftSampleFiles("EndpointDetailsReport")
    for _, file := range files["EndpointDetailsReport"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "EndpointDetailsReport", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []EndpointDetailsReport.CAMT_052_001_VESION{
                EndpointDetailsReport.CAMT_052_001_02,
                EndpointDetailsReport.CAMT_052_001_03,
                EndpointDetailsReport.CAMT_052_001_04,
                EndpointDetailsReport.CAMT_052_001_05,
                EndpointDetailsReport.CAMT_052_001_06,
                EndpointDetailsReport.CAMT_052_001_07,
                EndpointDetailsReport.CAMT_052_001_08,
                EndpointDetailsReport.CAMT_052_001_09,
                EndpointDetailsReport.CAMT_052_001_10,
                EndpointDetailsReport.CAMT_052_001_11,
                EndpointDetailsReport.CAMT_052_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestEndpointGapReport(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}
    files := getSwiftSampleFiles("EndpointGapReport")
    for _, file := range files["EndpointGapReport"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "EndpointGapReport", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []EndpointGapReport.CAMT_052_001_VESION{
                EndpointGapReport.CAMT_052_001_02,
                EndpointGapReport.CAMT_052_001_03,
                EndpointGapReport.CAMT_052_001_04,
                EndpointGapReport.CAMT_052_001_05,
                EndpointGapReport.CAMT_052_001_06,
                EndpointGapReport.CAMT_052_001_07,
                EndpointGapReport.CAMT_052_001_08,
                EndpointGapReport.CAMT_052_001_09,
                EndpointGapReport.CAMT_052_001_10,
                EndpointGapReport.CAMT_052_001_11,
                EndpointGapReport.CAMT_052_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestEndpointTotalsReport(t *testing.T) {
	wrapper := &EndpointTotalsReportWrapper{}
    files := getSwiftSampleFiles("EndpointTotalsReport")
    for _, file := range files["EndpointTotalsReport"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "EndpointTotalsReport", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []EndpointTotalsReport.CAMT_052_001_VESION{
                EndpointTotalsReport.CAMT_052_001_02,
                EndpointTotalsReport.CAMT_052_001_03,
                EndpointTotalsReport.CAMT_052_001_04,
                EndpointTotalsReport.CAMT_052_001_05,
                EndpointTotalsReport.CAMT_052_001_06,
                EndpointTotalsReport.CAMT_052_001_07,
                EndpointTotalsReport.CAMT_052_001_08,
                EndpointTotalsReport.CAMT_052_001_09,
                EndpointTotalsReport.CAMT_052_001_10,
                EndpointTotalsReport.CAMT_052_001_11,
                EndpointTotalsReport.CAMT_052_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestFedwireFundsAcknowledgement(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}
    files := getSwiftSampleFiles("FedwireFundsAcknowledgement")
    for _, file := range files["FedwireFundsAcknowledgement"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "FedwireFundsAcknowledgement", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []FedwireFundsAcknowledgement.ADMI_007_001_VESION{
                FedwireFundsAcknowledgement.ADMI_007_001_01,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestFedwireFundsPaymentStatus(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}
    files := getSwiftSampleFiles("FedwireFundsPaymentStatus")
    for _, file := range files["FedwireFundsPaymentStatus"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "FedwireFundsPaymentStatus", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []FedwireFundsPaymentStatus.PACS_002_001_VESION{
                FedwireFundsPaymentStatus.PACS_002_001_03,
                FedwireFundsPaymentStatus.PACS_002_001_04,
                FedwireFundsPaymentStatus.PACS_002_001_05,
                FedwireFundsPaymentStatus.PACS_002_001_06,
                FedwireFundsPaymentStatus.PACS_002_001_07,
                FedwireFundsPaymentStatus.PACS_002_001_08,
                FedwireFundsPaymentStatus.PACS_002_001_09,
                FedwireFundsPaymentStatus.PACS_002_001_10,
                FedwireFundsPaymentStatus.PACS_002_001_11,
                FedwireFundsPaymentStatus.PACS_002_001_12,
                FedwireFundsPaymentStatus.PACS_002_001_13,
                FedwireFundsPaymentStatus.PACS_002_001_14,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestFedwireFundsSystemResponse(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}
    files := getSwiftSampleFiles("FedwireFundsSystemResponse")
    for _, file := range files["FedwireFundsSystemResponse"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "FedwireFundsSystemResponse", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []FedwireFundsSystemResponse.ADMI_011_001_VESION{
                FedwireFundsSystemResponse.ADMI_011_001_01,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestMaster(t *testing.T) {
	wrapper := &MasterWrapper{}
    files := getSwiftSampleFiles("Master")
    for _, file := range files["Master"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "Master", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []Master.CAMT_052_001_VESION{
                Master.CAMT_052_001_02,
                Master.CAMT_052_001_03,
                Master.CAMT_052_001_04,
                Master.CAMT_052_001_05,
                Master.CAMT_052_001_06,
                Master.CAMT_052_001_07,
                Master.CAMT_052_001_08,
                Master.CAMT_052_001_09,
                Master.CAMT_052_001_10,
                Master.CAMT_052_001_11,
                Master.CAMT_052_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestPaymentReturn(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}
    files := getSwiftSampleFiles("PaymentReturn")
    for _, file := range files["PaymentReturn"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "PaymentReturn", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []PaymentReturn.PACS_004_001_VESION{
                PaymentReturn.PACS_004_001_02,
                PaymentReturn.PACS_004_001_03,
                PaymentReturn.PACS_004_001_04,
                PaymentReturn.PACS_004_001_05,
                PaymentReturn.PACS_004_001_06,
                PaymentReturn.PACS_004_001_07,
                PaymentReturn.PACS_004_001_08,
                PaymentReturn.PACS_004_001_09,
                PaymentReturn.PACS_004_001_10,
                PaymentReturn.PACS_004_001_11,
                PaymentReturn.PACS_004_001_12,
                PaymentReturn.PACS_004_001_13,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestPaymentStatusRequest(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}
    files := getSwiftSampleFiles("PaymentStatusRequest")
    for _, file := range files["PaymentStatusRequest"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "PaymentStatusRequest", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []PaymentStatusRequest.PACS_028_001_VESION{
                PaymentStatusRequest.PACS_028_001_01,
                PaymentStatusRequest.PACS_028_001_02,
                PaymentStatusRequest.PACS_028_001_03,
                PaymentStatusRequest.PACS_028_001_04,
                PaymentStatusRequest.PACS_028_001_05,
                PaymentStatusRequest.PACS_028_001_06,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}
func TestReturnRequestResponse(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}
    files := getSwiftSampleFiles("ReturnRequestResponse")
    for _, file := range files["ReturnRequestResponse"] {
        t.Run(file, func(t *testing.T) { // Subtest for each file
			xmlData, err := os.ReadFile(filepath.Join("..", "models", "ReturnRequestResponse", "swiftSample", file))
            require.NoError(t, err, "Failed to read file: %s", file)

            /* XML to model */
            model, err := wrapper.ConvertXMLToModel(xmlData)
            require.NoError(t, err, "Failed to convert XML to model for file: %s", file)
            require.NotNil(t, model, "Model is nil for file: %s", file)

            /* Check required fields */
            err = wrapper.CheckRequireField(model)
            require.NoError(t, err, "Required field check failed for file: %s", file)

            /* Get JSON data */
            jsonData, err := json.Marshal(model)
            require.NoError(t, err, "Failed to convert model to JSON for file: %s", file)
            require.NotEmpty(t, jsonData, "JSON data is empty for file: %s", file)

            /* Create Document */
            var versions = []ReturnRequestResponse.CAMT_029_001_VESION{
                ReturnRequestResponse.CAMT_029_001_03,
                ReturnRequestResponse.CAMT_029_001_04,
                ReturnRequestResponse.CAMT_029_001_05,
                ReturnRequestResponse.CAMT_029_001_06,
                ReturnRequestResponse.CAMT_029_001_07,
                ReturnRequestResponse.CAMT_029_001_08,
                ReturnRequestResponse.CAMT_029_001_09,
                ReturnRequestResponse.CAMT_029_001_10,
                ReturnRequestResponse.CAMT_029_001_11,
                ReturnRequestResponse.CAMT_029_001_12,
            }
            for _, version := range versions {
                t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) { // Subtest for each version
                    docData, _ := wrapper.CreateDocument(jsonData, version)
                    require.NotNil(t, docData, "Document is nil for file: %s with version: %s", file, version)

                    /* Validate Document */
                    err = wrapper.ValidateDocument(jsonData, version)
                    require.NoError(t, err, "Validation failed for file: %s with version: %s", file, version)
                })
            }
        })
    }
}