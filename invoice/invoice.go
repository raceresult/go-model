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
	var tax decimal.Decimal
	if i.TaxRate < 0 {
		tax = i.UnitPrice - i.UnitPrice.DivDecimal(decimal.FromInt(1)-i.TaxRate).Round(2)
	} else {
		tax = i.TaxRate.Mult(i.UnitPrice).Round(2)
	}
	return tax.Mult(i.Count)
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
	// sum up taxes by tax rate
	taxes := make(map[decimal.Decimal]decimal.Decimal)
	for _, invoiceItem := range q.Items {
		taxes[invoiceItem.TaxRate] += invoiceItem.GetTax()
	}
	return taxes
}

// GetTaxSum returns the sum of the VAT of all items
func (q WithDetails) GetTaxSum() decimal.Decimal {
	// sum up all taxes
	var taxSum decimal.Decimal
	for _, invoiceItem := range q.Items {
		taxSum += invoiceItem.GetTax()
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

		// add additional taxes
		if invoiceItem.TaxRate > 0 {
			gross[invoiceItem.TaxRate] += invoiceItem.GetTax()
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

		// subtract included taxes
		if invoiceItem.TaxRate < 0 {
			net[invoiceItem.TaxRate] -= invoiceItem.GetTax()
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
