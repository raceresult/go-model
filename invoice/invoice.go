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

// GetTax returns the amount of tax calculated for this item, this will not necessarily fit with the tax sum of the invoice because of different rounding
func (i Item) GetTax() decimal.Decimal {
	v := i.UnitPrice.Mult(i.Count)
	if i.TaxRate < 0 {
		return v - v.DivDecimal(decimal.FromInt(1)-i.TaxRate).Round(2)
	}
	return i.TaxRate.Mult(v).Round(2)
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
	// sum up amounts by tax rate
	taxRateSum := make(map[decimal.Decimal]decimal.Decimal)
	for _, invoiceItem := range q.Items {
		taxRateSum[invoiceItem.TaxRate] += invoiceItem.UnitPrice.Mult(invoiceItem.Count)
	}

	// calculate taxes per tax rate
	taxes := make(map[decimal.Decimal]decimal.Decimal)
	for k, v := range taxRateSum {
		if k > 0 {
			taxes[k] = v.Mult(decimal.FromInt(1)+k).Round(2) - v
		} else if k < 0 {
			taxes[k] = v - v.DivDecimal(decimal.FromInt(1)-k).Round(2)
		}
	}
	return taxes
}

// GetTaxSum returns the sum of the VAT of all items
func (q WithDetails) GetTaxSum() decimal.Decimal {
	taxes := q.GetTaxes()

	// sum up all taxes
	var taxSum decimal.Decimal
	for _, v := range taxes {
		taxSum += v
	}
	return taxSum
}

// GetGrossSum returns the gross amount of the invoice
func (q WithDetails) GetGrossSum() decimal.Decimal {
	gross := q.GetGrossByTaxRate()

	// sum up gross amounts
	var grossSum decimal.Decimal
	for _, v := range gross {
		grossSum += v
	}

	return grossSum
}

// GetGrossByTaxRate returns the gross amount of the invoice split up by tax rate
func (q WithDetails) GetGrossByTaxRate() map[decimal.Decimal]decimal.Decimal {
	// sum up amounts by tax rate
	gross := make(map[decimal.Decimal]decimal.Decimal)
	for _, invoiceItem := range q.Items {
		gross[invoiceItem.TaxRate] += invoiceItem.UnitPrice.Mult(invoiceItem.Count)
	}

	// add additional taxes not included in item unit price
	for k, v := range gross {
		if k > 0 {
			gross[k] += v.Mult(k).Round(2)
		}
	}

	return gross
}

// GetNetSum returns the net amount of the invoice
func (q WithDetails) GetNetSum() decimal.Decimal {
	net := q.GetNetByTaxRate()

	// sum up net amounts
	var netSum decimal.Decimal
	for _, v := range net {
		netSum += v
	}
	return netSum
}

// GetNetByTaxRate returns the net amount of the invoice split up by tax rate
func (q WithDetails) GetNetByTaxRate() map[decimal.Decimal]decimal.Decimal {
	// sum up amounts by tax rate
	net := make(map[decimal.Decimal]decimal.Decimal)
	for _, invoiceItem := range q.Items {
		net[invoiceItem.TaxRate] += invoiceItem.UnitPrice.Mult(invoiceItem.Count)
	}

	// subtract taxes included in item price
	for k, v := range net {
		if k < 0 {
			net[k] -= v - v.DivDecimal(decimal.FromInt(1)-k).Round(2)
		}
	}
	return net
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
