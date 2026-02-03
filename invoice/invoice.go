package invoice

import (
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/variant"
)

type Party struct {
	Company      string
	Name         string
	AddressLine1 string
	AddressLine2 string
	City         string
	ZIP          string
	State        string
	Country      int
	TaxID        string
	VATID        string
}

type Invoice struct {
	ID                int
	Number            string
	ReceiverReference string
	Date              date.Date
	PID               int
	Receiver          Party
	Issuer            Party
	Currency          string
	LegalNotes        string
	PerformanceDate   date.Date
	EInvoice          string
	PaymentTerms      PaymentTerms
	Sent              time.Time
}

type PaymentTerms struct {
	Method       int
	AccountOwner string
	IBAN         string
	BIC          string
	AccountNo    string
	BranchNo     string
	BankName     string
	MandateID    string
	DueDays      int
	Reference    string
	Terms        string
}

type WithSum struct {
	*Invoice
	Sum  decimal.Decimal
	PIDs []int
}

type SourceItem struct {
	ID             int
	InvoiceID      int
	PID            int
	EntryFeeID     int
	EntryFeeName   string
	Multiplication decimal.Decimal
	Amount         decimal.Decimal
	TaxRate        decimal.Decimal
	Credit         bool
}

type Item struct {
	Count        decimal.Decimal
	EntryFeeID   int
	EntryFeeName string
	UnitPrice    decimal.Decimal
	TaxRate      decimal.Decimal
}

type WithDetails struct {
	*Invoice
	Sum         decimal.Decimal
	Items       []*Item
	SourceItems []*SourceItem
	Fields      variant.VariantMap
}

// GetTaxes returns a map that holds the sum of the VAT of the items per tax rate
func (q WithDetails) GetTaxes() map[decimal.Decimal]decimal.Decimal {
	taxes := make(map[decimal.Decimal]decimal.Decimal)
	for _, invoiceItem := range q.Items {
		v := invoiceItem.UnitPrice.Mult(invoiceItem.Count)

		if invoiceItem.TaxRate < 0 {
			taxes[invoiceItem.TaxRate] = v - v.DivDecimal(decimal.FromInt(1)-invoiceItem.TaxRate).Round(2)
		} else {
			taxes[invoiceItem.TaxRate] = invoiceItem.UnitPrice.Mult(invoiceItem.Count).Mult(invoiceItem.TaxRate).Round(2)
		}
	}
	return taxes
}

// GetTaxSum returns the sum of the VAT of all items
func (q WithDetails) GetTaxSum() decimal.Decimal {
	var taxSum decimal.Decimal
	for _, invoiceItem := range q.Items {
		v := invoiceItem.UnitPrice.Mult(invoiceItem.Count)

		if invoiceItem.TaxRate < 0 {
			taxSum += v - v.DivDecimal(decimal.FromInt(1)-invoiceItem.TaxRate).Round(2)
		} else {
			taxSum += v.Mult(invoiceItem.TaxRate).Round(2)
		}
	}
	return taxSum
}

// GetGrossSum returns the gross amount of the invoice
func (q WithDetails) GetGrossSum() decimal.Decimal {
	var grossSum decimal.Decimal
	for _, invoiceItem := range q.Items {
		v := invoiceItem.UnitPrice.Mult(invoiceItem.Count)
		grossSum += v
		if invoiceItem.TaxRate > 0 {
			grossSum += v.Mult(invoiceItem.TaxRate).Round(2)
		}
	}
	return grossSum
}

// GetNetSum returns the net amount of the invoice
func (q WithDetails) GetNetSum() decimal.Decimal {
	var netSum decimal.Decimal
	for _, invoiceItem := range q.Items {
		v := invoiceItem.UnitPrice.Mult(invoiceItem.Count)
		if invoiceItem.TaxRate < 0 {
			netSum += v.DivDecimal(decimal.FromInt(1) - invoiceItem.TaxRate).Round(2)
		} else {
			netSum += v
		}
	}
	return netSum
}

type WithTaxDetails struct {
	*WithDetails
	Taxes       map[decimal.Decimal]decimal.Decimal
	TaxSum      decimal.Decimal
	GrossAmount decimal.Decimal
	NetAmount   decimal.Decimal
}

type Settings struct {
	NumberScheme              string
	LegalNotes                string
	PerformanceDate           date.Date
	EInvoice                  string
	ReceiverFilter            string
	ReceiverMergeField        string
	ReceiverFieldCompany      string
	ReceiverFieldName         string
	ReceiverFieldAddressLine1 string
	ReceiverFieldAddressLine2 string
	ReceiverFieldCity         string
	ReceiverFieldZIP          string
	ReceiverFieldState        string
	ReceiverFieldCountry      string
	ReceiverFieldReference    string
	ReceiverFieldTaxID        string
	ReceiverFieldVATID        string
	Issuer                    Party
	PaymentTerms              PaymentTerms
	ExcludedEntryFees         []int
}

type Filter struct {
	ID                    []int
	Number                []string
	PID                   []int
	Sent                  []bool
	PaymentTermsReference []string
}
