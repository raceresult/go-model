package presenter

import (
	"github.com/raceresult/go-model/datetime"
)

// Presenter represents all settings of a Presenter screen
type Presenter struct {
	Name         string
	Key          string
	Title        string
	Enabled      bool
	EnabledFrom  datetime.DateTime
	EnabledTo    datetime.DateTime
	SwitchMode   string
	AutoHideTabs bool
	Screens      []Screen
	CSS          string
}

// Screen represents a screen as part of a Presenter
type Screen struct {
	Disabled        bool
	Title           string
	BackgroundColor string
	BackgroundImage string
	CSS             string
	Windows         []Window
}

// Window represents a window as part of a Presenter Screen
type Window struct {
	ItemType             string
	ItemName             string
	Left, Top            int
	Width, Height        int
	Contests             []int
	Results              []int
	TimingPoints         []string
	Splits               []string
	ScrollMode           string
	PageTime             int
	FontSize             int
	ScrollBar            int
	Filter               string
	ShowFilter           bool
	IgnoreManualPassings bool
}
