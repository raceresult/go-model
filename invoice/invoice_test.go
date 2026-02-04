package invoice

import (
	"testing"

	"github.com/raceresult/go-model/decimal"
	"github.com/stretchr/testify/assert"
)

func TestWithDetails_GetTaxSum(t *testing.T) {
	invoiceWithDetails := &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.01), invoiceWithDetails.GetTaxSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.01), invoiceWithDetails.GetTaxSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(-0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.01), invoiceWithDetails.GetTaxSum())

}

func TestWithDetails_GetGrossSum(t *testing.T) {
	invoiceWithDetails := &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.02), invoiceWithDetails.GetGrossSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.03), invoiceWithDetails.GetGrossSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(-0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.03), invoiceWithDetails.GetGrossSum())
}

func TestWithDetails_GetNetSum(t *testing.T) {
	invoiceWithDetails := &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.01), invoiceWithDetails.GetNetSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.02), invoiceWithDetails.GetNetSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(-0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.02), invoiceWithDetails.GetNetSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(10.0),
				TaxRate:   decimal.FromFloat(-0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(6.67), invoiceWithDetails.GetNetSum())
}

func TestWithDetails_GetTaxes(t *testing.T) {
	invoiceWithDetails := &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	taxes := invoiceWithDetails.GetTaxes()
	assert.Equal(t, decimal.FromFloat(0.005), taxes[decimal.FromFloat(0.5)])

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	taxes = invoiceWithDetails.GetTaxes()
	assert.Equal(t, decimal.FromFloat(0.01), taxes[decimal.FromFloat(0.5)])

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(-0.5),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	taxes = invoiceWithDetails.GetTaxes()
	assert.Equal(t, decimal.FromFloat(0.005), taxes[decimal.FromFloat(0.5)])
	assert.Equal(t, decimal.FromFloat(0.0033), taxes[decimal.FromFloat(-0.5)])
}
