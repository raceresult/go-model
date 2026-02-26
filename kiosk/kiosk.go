package kiosk

import (
	"encoding/json"

	"github.com/raceresult/go-model/datetime"
)

// Kiosk represents all settings of a kiosk
type Kiosk struct {
	Name                 string
	Key                  string
	Enabled              bool
	EnabledFrom          datetime.DateTime
	EnabledTo            datetime.DateTime
	TransponderMode      int
	AcceptedTransponders int
	IgnoreBibRanges      bool
	AutoFinish           bool
	CSS                  string
	Title                string
	Steps                []Step
	AfterSave            []AfterSave
}

type Step struct {
	Type  string
	Label string

	Title         string
	Text          string
	OnlyShowIf    string
	SearchFields  []SearchField
	DisplayFields []DisplayField
	EditFields    []EditField
	Settings      map[string]interface{}
}

type AfterSave struct {
	Type            string
	Value           string
	Destination     string
	Filter          string
	Printer         string
	PrintAutorotate bool
}

type DisplayField struct {
	Type  string
	Value string
	Label string
}

type EditField struct {
	Label          string
	Field          string
	Special        string
	Mandatory      bool
	ValidationRule string
	ValidationMsg  string
	EventTools     string
}

type SearchField struct {
	Field    string
	Hide     bool // legacy
	Function string
}

func (a *AfterSave) UnmarshalJSON(data []byte) error {
	type Alias AfterSave
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	// default true, if value does not exist
	a.PrintAutorotate = true
	return json.Unmarshal(data, &aux)
}
