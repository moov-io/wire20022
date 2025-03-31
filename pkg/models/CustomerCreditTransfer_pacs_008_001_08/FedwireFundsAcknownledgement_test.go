package CustomerCreditTransfer_pacs_008_001_08

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	fedwire "github.com/moov-io/wire20022/pkg/internal"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknownledgementCreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000713"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01Step3EndToEndId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f099"
	mesage.model.InstrumentPropCode = InstrumentCTRD
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 60000.00,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 60000.00,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19037",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation A"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "5647772655"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "INV12345",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	jsonData, err := mesage.GetJson()
	require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	jsonFileName := filepath.Join("generated", "FedwireFundAcknowledgement.json")
	xmlFileName := filepath.Join("generated", "FedwireFundAcknowledgement.xml")
	err = os.WriteFile(jsonFileName, jsonData, 0644)
	require.NoError(t, err)
	err = os.WriteFile(xmlFileName, xmlData, 0644)
	require.NoError(t, err)

	sample_xml_file := "../../../test/swiftSample/pacs_008/FedwireFundsAcknowledgement_Scenario1_Step3_pacs.008"
	generated_xml_file := xmlFileName

	xmlData1, err := ioutil.ReadFile(sample_xml_file)
	if err != nil {
		log.Fatal("Error reading sample.xml:", err)
	}

	xmlData2, err := ioutil.ReadFile(generated_xml_file)
	if err != nil {
		log.Fatal("Error reading generate.xml:", err)
	}

	// Normalize XML structure
	var node1, node2 Document
	xml.Unmarshal(bytes.TrimSpace(xmlData1), &node1)
	xml.Unmarshal(bytes.TrimSpace(xmlData2), &node2)

	// Compare XML contents
	diff := cmp.Diff(node1, node2,
		cmpopts.IgnoreUnexported(fedwire.ISODateTime{}, fedwire.ISODate{}),
		cmp.Comparer(equateApproxTolerance),
		cmp.FilterPath(func(path cmp.Path) bool {
			ignoredPaths := map[string]bool{
				"FIToFICstmrCdtTrf.GrpHdr.CreDtTm":                 true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmDt":      true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf.Strd.":       true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.AccptncDtTm":        true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.SttlmTmReq":         true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.InstrPrty": true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf.Ustrd":       true,
				"FIToFICstmrCdtTrf.CdtTrfTxInf.SttlmTmIndctn":      true,
			}

			// Check if path matches any of the ignored fields
			return ignoredPaths[path.String()]
		}, cmp.Ignore()))
	if diff == "" {
		fmt.Println("XML files are identical")
	} else {
		fmt.Println("Differences found:")
		fmt.Println(diff)
	}
}
func equateApproxTolerance(x, y float64) bool {
	// Define the tolerance value
	const tolerance = 0.00001
	return math.Abs(x-y) < tolerance
}
