package invoice

import (
	"testing"

	"github.com/raceresult/go-model/decimal"
	"github.com/stretchr/testify/assert"
)

func TestItem_GetTax(t *testing.T) {
	item := &Item{
		Count:     decimal.FromInt(1),
		UnitPrice: decimal.FromFloat(10.01),
		TaxRate:   decimal.FromFloat(0.5),
	}
	assert.Equal(t, decimal.FromFloat(5.01), item.GetTax())

	item = &Item{
		Count:     decimal.FromInt(1),
		UnitPrice: decimal.FromFloat(10.01),
		TaxRate:   decimal.FromFloat(-0.5),
	}
	assert.Equal(t, decimal.FromFloat(3.34), item.GetTax())
}

func TestWithDetails_GetTaxSum(t *testing.T) {
	invoiceWithDetails := &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(2),
				UnitPrice: decimal.FromFloat(0.01),
				TaxRate:   decimal.FromFloat(0.5),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(0.02), invoiceWithDetails.GetTaxSum())

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
	assert.Equal(t, decimal.FromFloat(0.02), invoiceWithDetails.GetTaxSum())

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

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(10.01),
				TaxRate:   decimal.FromFloat(-1),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(5.0), invoiceWithDetails.GetTaxSum())
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
	assert.Equal(t, decimal.FromFloat(0.04), invoiceWithDetails.GetGrossSum())

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

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(105),
				UnitPrice: decimal.FromFloat(27.5),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(105),
				UnitPrice: decimal.FromFloat(10),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(600),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(9.5),
				TaxRate:   decimal.FromFloat(0.19),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(5411.46), invoiceWithDetails.GetGrossSum())
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

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(10.01),
				TaxRate:   decimal.FromFloat(-1),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(5.01), invoiceWithDetails.GetNetSum())

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(105),
				UnitPrice: decimal.FromFloat(27.5),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(105),
				UnitPrice: decimal.FromFloat(10),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(600),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(9.5),
				TaxRate:   decimal.FromFloat(0.19),
			},
		},
	}
	assert.Equal(t, decimal.FromFloat(4547), invoiceWithDetails.GetNetSum())
}

func TestWithDetails_GetTaxes(t *testing.T) {
	// additional taxes are rounded to 2 digits per tax rate
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
	assert.Equal(t, decimal.FromFloat(0.01), taxes[decimal.FromFloat(0.5)])

	// additional taxes are rounded to 2 digits per tax rate
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
	assert.Equal(t, decimal.FromFloat(0.02), taxes[decimal.FromFloat(0.5)])

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
	assert.Equal(t, decimal.FromFloat(0.01), taxes[decimal.FromFloat(0.5)])
	assert.Equal(t, decimal.FromFloat(0.0), taxes[decimal.FromFloat(-0.5)])

	invoiceWithDetails = &WithDetails{
		Items: []*Item{
			{
				Count:     decimal.FromInt(105),
				UnitPrice: decimal.FromFloat(27.5),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(105),
				UnitPrice: decimal.FromFloat(10),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(600),
				TaxRate:   decimal.FromFloat(0.19),
			},
			{
				Count:     decimal.FromInt(1),
				UnitPrice: decimal.FromFloat(9.5),
				TaxRate:   decimal.FromFloat(0.19),
			},
		},
	}
	taxes = invoiceWithDetails.GetTaxes()
	assert.Equal(t, decimal.FromFloat(864.46), taxes[decimal.FromFloat(0.19)])
}
