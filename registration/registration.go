package registration

import (
	"github.com/raceresult/go-model/datetime"
)

type Registration struct {
	Name                                       string
	Key                                        string
	ChangeKeySalt                              string
	Title                                      string
	Enabled                                    bool
	EnabledFrom                                datetime.DateTime
	EnabledTo                                  datetime.DateTime
	TestModeKey                                string
	Type                                       string
	GroupMin, GroupMax, GroupDefault, GroupInc int
	Contest                                    int
	Limit                                      int
	ChangeIdentityField                        string
	ChangeIdentityFilter                       string
	Steps                                      []Step
	AdditionalValues                           []AdditionalValue
	CheckSex                                   bool
	CheckDuplicate                             bool
	OnlinePayment                              bool
	OnlinePaymentButtonText                    string
	PaymentMethods                             []PaymentMethod
	OnlineRefund                               bool
	RefundMethods                              []PaymentMethod
	Confirmation                               Confirmation
	AfterSave                                  []AfterSave
	CSS                                        string
	ErrorMessages                              ErrorMessages
}

type Step struct {
	ID          int
	Title       string
	Enabled     bool
	EnabledFrom datetime.DateTime
	EnabledTo   datetime.DateTime
	Elements    []Element
	ButtonText  string
	Filter      string
}

type Element struct {
	Type            string // text, box, field, entryfee list, ...
	Label           string
	Enabled         bool
	EnabledFrom     datetime.DateTime
	EnabledTo       datetime.DateTime
	Field           *Field
	ShowIf          string
	ShowIfMode      int
	ShowIfCurr      string
	ShowIfCurrMode  int
	ShowIfInitial   bool // legacy, can be removed at some point
	Styles          []Style
	ClassName       string
	ID              int
	Common          int // common in group reg
	ValidationRules []ValidationRule
	Children        []Element
}

type Field struct {
	Name              string // field name
	ControlType       string // currently unused
	Mandatory         int
	DefaultValue      string
	DefaultValueType  int
	Placeholder       string
	Unique            string
	Special           string
	SpecialDetails    string
	ForceUpdate       bool
	Values            []Value  // advanced drop down value settings
	AdditionalOptions []string // additional options for PROPOSE/SELECT
	Flags             []string
}

type Style struct {
	Attribute string
	Value     string
}

type Value struct {
	Value       interface{}
	Label       string
	Enabled     bool
	EnabledFrom datetime.DateTime
	EnabledTo   datetime.DateTime
	MaxCapacity int
	ShowIf      string
}

type AdditionalValue struct {
	FieldName     string
	Source        string
	Value         string
	Filter        string
	FilterInitial string
}

type Confirmation struct {
	Title      string
	Expression string
}

type AfterSave struct {
	Type        string
	Value       string
	Destination string
	Filter      string
	Flags       []string
}

type PaymentMethod struct {
	ID          int
	Label       string
	Enabled     bool
	EnabledFrom datetime.DateTime
	EnabledTo   datetime.DateTime
	Filter      string
}

type ValidationRule struct {
	Rule string
	Msg  string
}

type ErrorMessages struct {
	BeforRegStart string
	AfterRegEnd   string
}
