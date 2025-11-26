package pay

import (
	model "github.com/raceresult/go-model"
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/invoice"
	"github.com/raceresult/go-model/variant"
	"time"
)

const (
	PmNoPayment        int = 0
	PmCCEUR            int = 2
	PmCCCHF            int = 3
	PmUEBD             int = 4
	PmBAR              int = 5
	PmSPF              int = 6
	PmPPalEUR          int = 7
	PmUEBCH            int = 8
	PmEINZCH           int = 10
	PmUEBSOF           int = 12
	PmPPalGBP          int = 14
	PmPPalUSD          int = 15
	PmSEPA             int = 16
	PmCCGBP            int = 17
	PmSEPADATA         int = 19
	PmOwnEPAY          int = 20
	PmOwnPPal          int = 21
	PmOwnWireT         int = 22
	PmOwnPaytrail      int = 25
	PmOwnOnePAY        int = 26
	PmTelr             int = 27
	PmOwnOnePAYDom     int = 28
	PmFatora           int = 29
	PmTwint            int = 30
	PmStripeCard       int = 31
	PmOwnPaytrailV2    int = 32
	PmTelrSale         int = 33
	PmRedSys           int = 34
	PmMollieBancontact int = 35
	PmPayTabs          int = 36
	PmAsiaPay          int = 37
	PmMercadoPago      int = 38
	PmCB               int = 99
)

const (
	PayStateUndefined int = 0
	PayStatePending   int = 1
	PayStateUnderpaid int = 2
	PayStatePaid      int = 3
	PayStateOverpaid  int = 4
	PayStateNoPayout  int = 5
)

// MethodOption is a payment method option the user will be offered
type MethodOption struct {
	ID            int
	NameShort     string
	Name          string
	EntryFee      decimal.Decimal
	PaymentFee    decimal.Decimal
	UserFee       decimal.Decimal
	Kickback      decimal.Decimal `json:"KB"`
	Currency      string
	ExchangeRate  float64
	SEPANotBefore string `json:",omitempty"`
	NoTestMode    bool
	Token         string
}

type Registration struct {
	ID        int
	Timestamp time.Time
	IP        string
	Event     int
	FirstName string
	LastName  string
	Sex       string
	Contest   int
	Mail      string
	Street    string
	ZIP       string
	City      string
	Country   int
	Lang      string
	PaymentID int
	PID       int
}

type Payment struct {
	ID            int
	CustNo        int
	Event         int
	Method        int
	Currency      string
	AmountNew     decimal.Decimal
	Fees          decimal.Decimal
	UserFees      decimal.Decimal
	Kickback      decimal.Decimal
	ExchangeRate  float64
	Created       time.Time
	PayState      int
	EventCurrency string
	Reference     string
	Email         string
	BillNo        int
	RetryOf       int
	Lang          string
	IgnorePayment bool
	IgnoreReason  string
	RequestID     int
	KickbackInvID int
}

type Contract struct {
	CustNo       int
	ContractDate time.Time
	VATRate      decimal.Decimal
	AccountOwner string
	IBAN         string
	SWIFT        string
	Type         string
	Language     string
	HasContract  bool `db:"length(Contract)>0"`
}

type EventReceiver struct {
	EventID       int
	BillCustNo    int
	AccountOwner  string
	IBAN          string
	SWIFT         string
	Email         string
	VATRate       decimal.Decimal
	SEPANotBefore time.Time
	ModJobID      int
}

type Method struct {
	ID                     int
	NameShort              string
	Name                   string
	Currency               string
	TransactionFee         decimal.Decimal
	Disagio                decimal.Decimal
	RegFee                 decimal.Decimal
	RefundFee              decimal.Decimal
	TransactionCosts       decimal.Decimal
	DisagioCosts           decimal.Decimal
	TransferDelay          int
	Activated              bool
	NoPayout               bool
	BankAccountID          int
	CaptureAmountAccountID int
	NoPayoutReceival       bool
	NoTestMode             bool
	Rounding               decimal.Decimal
	DontShowFee            bool
}

type BankCheckResult struct {
	Result bool
	Error  string
}

type CreateRequest struct {
	CustNo             int
	Event              int
	IP                 string
	RegName            string
	RegKey             string
	TestMode           bool
	ChangeID           int
	Amount             decimal.Decimal
	Fees               decimal.Decimal
	UserFees           decimal.Decimal
	Kickback           decimal.Decimal
	Currency           string
	ExchangeRate       float64
	EventCurrency      string
	PaymentData        variant.VariantMap
	Data               string
	EntryFeeDetails    [][]model.EntryFeeItem
	Method             int
	ReferenceID        int
	Lang               string
	DataToken          string
	PaymentMethodToken string
}

type CreateResult struct {
	PaymentToken   string            `json:"paymentToken"`
	Redirect       string            `json:"redirect"`
	RedirectParams map[string]string `json:"redirectParams"`
	ConfirmDetails map[string]string `json:"confirmDetails"`
}

type FinishResult struct {
	PaymentToken  string
	RefusalReason string
	PmID          int
	Data          string
	DataToken     string
}

type CaptureResponse struct {
	AddFields    variant.VariantMap
	OPJSONs      map[int]string
	PaymentTerms invoice.PaymentTerms
}
