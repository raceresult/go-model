package pay

import (
	"github.com/raceresult/go-model/decimal"
	"time"
)

type RegistrationFilter struct {
	ID           []int
	Event        []int
	PID          []int
	Month        []int
	Year         []int
	PaymentID    []int
	Search       string
	MinTimestamp time.Time
	MaxTimestamp time.Time
}

type PaymentFilter struct {
	ID         []int
	RetryOf    []int
	Event      []int
	Month      []int
	Year       []int
	Method     []int
	Reference  []string
	Email      []string
	Search     []string
	MinCreated time.Time
	MaxCreated time.Time
	ToPay      decimal.Decimal
	RequestID  []int
	BillNo     []int
	PayState   []int
}
