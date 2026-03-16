package portal

import (
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/decimal"
)

type RegistrationContestEntryFee struct {
	DateStart date.Date
	DateEnd   date.Date
	RegStart  time.Time
	RegEnd    time.Time
	Fee       decimal.Decimal
	Currency  string
}

type RegistrationContest struct {
	ID          int
	EnabledFrom time.Time
	EnabledTo   time.Time
	Start       time.Time
	Name        string
	Sex         string
	AgeStart    date.Date
	AgeEnd      date.Date
	EntryFees   []RegistrationContestEntryFee
	SlotsLeft   int
}

type Registration struct {
	Name          string
	Title         string
	Key           string
	Type          string
	TestModeKey   string
	EnabledFrom   time.Time
	EnabledTo     time.Time
	InfoText      string
	ShowContests  bool
	HideEntryFees bool
	ButtonText    string
	SlotsLeft     int
	Contests      []RegistrationContest
	IfLoggedIn    int
}

type RegistrationConfig struct {
	Registrations []Registration
	Login         RegistrationLogin
}

type RegLoginField struct {
	Name  string
	Label string
}

type RegistrationLogin struct {
	Enabled        bool
	InfoText       string
	LoginNameField string
	Fields         []RegLoginField
}

type RegistrationLoginResponse struct {
	ID         int
	ChangeKeys map[string]string
	LoginName  string
}
