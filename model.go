package sesbase

import (
	"github.com/raceresult/go-model/invoice"
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/variant"
)

// AgeGroup describes the internal go model
type AgeGroup struct {
	ID        int
	Name      string
	NameShort string
	DateStart date.Date
	DateEnd   date.Date
	AgeFrom   int
	AgeTo     int
	Contest   int
	AGSet     int
	OrderPos  int
	Sex       string
}

// BibRange describes the internal go model
type BibRange struct {
	ID              int
	BibStart        int
	BibEnd          int
	Contest         int
	TimeDifference  decimal.Decimal
	FinishTimeLimit decimal.Decimal
	Comment         string
	Filter          string
}

// Contest describes the internal go model
type Contest struct {
	ID               int
	Name             string
	NameShort        string
	AgeStart         date.Date
	AgeEnd           date.Date
	Sex              string
	Day              int
	StartTime        decimal.Decimal
	Length           decimal.Decimal
	LengthUnit       string
	TimeFormat       string
	TimeRounding     int
	StartTransponder int
	StartResult      int
	TimeDifference   decimal.Decimal
	FinishResult     int
	FinishTimeLimit  decimal.Decimal
	Laps             int
	MinResultID      int
	MinLapTime       decimal.Decimal
	TimingMode       int
	TimingModeFilter string
	Attributes       string
	OrderPos         float64
	Sort1            string
	Sort2            string
	Sort3            string
	Sort4            string
	SortDesc1        bool
	SortDesc2        bool
	SortDesc3        bool
	SortDesc4        bool
	Inactive         bool
}

// CustomFieldType defines the type of a customfield
type CustomFieldType int

// CustomFieldType constants
const (
	CustomFieldTypeText        CustomFieldType = 0
	CustomFieldTypeDropDown    CustomFieldType = 1
	CustomFieldTypeYesNo       CustomFieldType = 2
	CustomFieldTypeInteger     CustomFieldType = 3
	CustomFieldTypeDecimal     CustomFieldType = 4
	CustomFieldTypeDate        CustomFieldType = 5
	CustomFieldTypeCurrency    CustomFieldType = 6
	CustomFieldTypeCountry     CustomFieldType = 7
	CustomFieldTypeEmail       CustomFieldType = 8
	CustomFieldTypeCellPhone   CustomFieldType = 9
	CustomFieldTypeTransponder CustomFieldType = 10
)

// CustomField describes the internal go model
type CustomField struct {
	ID          int
	Name        string
	AltName     string
	Group       string
	FieldType   CustomFieldType `json:"Type"`
	Enabled     bool
	Mandatory   bool
	Config      string
	Default     string
	Placeholder string
	Label       string
	OrderPos    int
	MinLen      int
	MaxLen      int
}

// Participant describes the internal go model
type Participant struct {
	ID                int
	Bib               int
	Transponder1      string
	Transponder2      string
	RegNo             string
	Title             string
	Lastname          string
	Firstname         string
	Sex               string
	DateOfBirth       date.Date
	Street            string
	ZIP               string
	City              string
	State2            string
	Country           string
	Nation            string
	AgeGroup1         int
	AgeGroup2         int
	AgeGroup3         int
	Club              string
	Contest           int
	Status            int
	Booleans          int
	PaidEntryFee      decimal.Decimal
	Phone             string
	CellPhone         string
	SendSMS           int
	Email             string
	AccountNo         string
	BranchNo          string
	Bankname          string
	AccountOwner      string
	IBAN              string
	BIC               string
	SEPAMandate       string
	Comment           string
	Created           datetime.DateTime
	Modified          datetime.DateTime
	Uploaded          datetime.DateTime
	CreatedBy         string
	ForeignID         int
	RecordPayGUID     string
	ActivationEventID string
	OPJSON            string
	OPID              int
	OPMethod          int
	OPStatus          int
	OPEntryFee        decimal.Decimal
	OPUserFee         decimal.Decimal
	OPPaymentFee      decimal.Decimal
	OPCurrency        string
	OPToPay           decimal.Decimal
	OPBalance         decimal.Decimal
	OPBalanceDate     datetime.DateTime
	OPReference       string
	License           string
	ShowUnderscores   bool
	GroupRegPos       int
	GroupID           int
	Password          string
	Voucher           string
	Language          string
}

// ParticipantNewResponse is the response to part/new?v2=true
type ParticipantNewResponse struct {
	ID  int
	Bib int
}

// EntryFee describes the internal go model
type EntryFee struct {
	ID              int
	Name            string
	Contest         int
	DateStart       date.Date
	DateEnd         date.Date
	RegStart        date.Date
	RegEnd          date.Date
	Field           string
	Operator        string
	Value           string
	Fee             decimal.Decimal
	ShowAsBasicFee  bool
	IsMultiplicator bool
	Multiplication  string
	Category        string
	Tax             decimal.Decimal
	OrderPos        int
}

// Exporter describes the internal go model
type Exporter struct {
	ID                 int
	Name               string
	Filter             string
	TriggerTimingPoint string
	TriggerSplit       string
	TriggerResultID    int
	DestinationType    string
	Destination        string
	Data               string
	MTB                int
	MQL                int
	LineEnding         string
	StartPaused        bool
	IgnoreBefore       decimal.Decimal
	IgnoreAfter        decimal.Decimal
	Encoding           string
	ConnectMsg         string
	OrderPos           int
}

// HistoryCount is the result of a history result query
type HistoryCount struct {
	PID   int
	Count int
}

// HistoryEntry is external model api model for history entries
type HistoryEntry struct {
	ID          int
	Bib         int
	PartID      int
	DateTime    time.Time
	FieldName   string
	OldValue    variant.Variant
	NewValue    variant.Variant
	User        string
	Application string
}

// Overwrite describes the internal go model
type Overwrite struct {
	ID       int
	PID      int
	ResultID int
	Value    decimal.Decimal
}

// Ranking describes the internal go model
type Ranking struct {
	ID          int
	Name        string
	Group       []string
	Sort        []string
	SortDesc    []bool
	UseTies     bool
	ContestSort bool
	Filter      string
}

// RawData describes the internal go model
type RawData struct {
	ID          int
	PID         int
	TimingPoint string
	Result      int
	Time        decimal.Decimal
	Invalid     bool
	Passing     Passing
}

// RawDataReduced describes the internal go model
type RawDataReduced struct {
	TimingPoint string
	PID         int
	Time        decimal.Decimal
	Invalid     bool
	OrderID     int
	Result      int
	IsMarker    bool
	RSSI        int
}

// RawDataFilter describes the internal go model
type RawDataFilter struct {
	ID          []int             `json:",omitempty"`
	MinID       int               `json:",omitempty"`
	MaxID       int               `json:",omitempty"`
	TimingPoint []string          `json:",omitempty"`
	MinTime     decimal.Decimal   `json:",omitempty"`
	MaxTime     decimal.Decimal   `json:",omitempty"`
	Result      []int             `json:",omitempty"`
	DeviceID    []string          `json:",omitempty"`
	DeviceName  []string          `json:",omitempty"`
	Transponder []string          `json:",omitempty"`
	OrderID     []int             `json:",omitempty"`
	Hits        []int             `json:",omitempty"`
	RSSI        []int             `json:",omitempty"`
	LoopID      []byte            `json:",omitempty"`
	Channel     []byte            `json:",omitempty"`
	Battery     []decimal.Decimal `json:",omitempty"`
	Port        []int             `json:",omitempty"`
	StatusFlags []int             `json:",omitempty"`
	FileNo      []int             `json:",omitempty"`
	PassingNo   []int             `json:",omitempty"`
	IsMarker    []bool            `json:",omitempty"`
}

// RawDataWithAdditionalFields describes the internal go model
type RawDataWithAdditionalFields struct {
	RawData
	Bib    int
	Fields variant.VariantMap `json:",omitempty"`
}

// RawDataRule describes the internal go model
type RawDataRule struct {
	ID        int
	ResultID  int
	ContestID int
	Mode      int
	N         int
	Min       int
	MinOffset decimal.Decimal
	Max       int
	MaxOffset decimal.Decimal
	Ref       int
	RefOffset decimal.Decimal
}

type RawDataDistinctValues struct {
	DecoderID      []string
	OrderID        []int
	BatteryVoltage []decimal.Decimal
	Hits           []int
	RSSI           []int
}

// Result describes the internal go model
type Result struct {
	ID           int
	Name         string
	Formula      string
	TimeFormat   string
	Location     string
	TimeRounding int
}

// Split describes the internal go model
type Split struct {
	ID           int
	Contest      int
	Name         string
	TimingPoint  string
	Backup       string
	BackupOffset decimal.Decimal
	TypeOfSport  int
	Distance     decimal.Decimal
	DistanceUnit string
	DistanceFrom int
	TimeMin      decimal.Decimal
	TimeMax      decimal.Decimal
	Color        string
	OrderPos     int
	SplitType    int
	SectorFrom   int
	SectorTo     int
	SpeedOrPace  string
	TimeMode     int
	Label        string
	SectorFrom2  int
	SectorTo2    int
}

// SplitType constants
const (
	SplitTypeSplit    = 0
	SplitTypeInternal = 2
	SplitTypeLeg      = 9
)

const (
	SplitTimeModeRefSplit = 1
	SplitTimeModeRaceTime = 0
	SplitTimeModeTOD      = -1
	SplitTimeModeDelta    = -2
	SplitTimeModeMinKm    = -3
	SplitTimeModeMinMile  = -4
	SplitTimeModeMin100m  = -5
	SplitTimeModeKmh      = -6
	SplitTimeModeMph      = -7
	SplitTimeModeMps      = -8
)

const (
	SplitSpeedNone       = 0
	SplitSpeedMinPerKM   = 1
	SplitSpeedMinPerMile = 2
	SplitSpeedMinPer100m = 3
	SplitSpeedKmh        = 4
	SplitSpeedMph        = 5
	SplitSpeedMps        = 6
)

// TableValue describes the internal go model
type TableValue struct {
	ID     int
	Index1 int
	Index2 int
	Index3 int
	Index4 int
	Value  decimal.Decimal
}

// TeamScore describes the internal go model
type TeamScore struct {
	ID                        int
	ResultID1                 int
	ResultID2                 int
	ResultID3                 int
	ResultID4                 int
	ResultMode1               int
	ResultMode2               int
	ResultMode3               int
	ResultMode4               int
	SortDesc1                 bool
	SortDesc2                 bool
	SortDesc3                 bool
	RealTime                  bool
	MinTotal                  int
	MaxTotal                  int
	MinFemale                 int
	MaxFemale                 int
	MaxTeams                  int
	Filter                    string
	TimeFormat                string
	LapTimes                  int
	LapTimesLemans            bool
	LapTimesZeroStart         bool
	Name                      string
	LapModeLocation           string
	TeamSort                  string
	Assigning1                string
	Grouping1                 string
	Assigning2                string
	Grouping2                 string
	Assigning3                string
	Grouping3                 string
	Assigning4                string
	Grouping4                 string
	UseTies                   bool
	LapTimesSubtractT0        bool
	LapTimesCountLemansAsLap  bool
	LapTimesPenaltyTimeResult int
	LapTimesPenaltyLapsResult int
	LapTimesMinLapTime        decimal.Decimal
	LapTimesIgnoreBefore      decimal.Decimal
	LapTimesIgnoreAfter       decimal.Decimal
}

// Time describes the internal go model
type Time struct {
	PID         int
	Result      int
	DecimalTime decimal.Decimal
	TimeText    string
	InfoText    string
}

// TimingPointRule describes the internal go model
type TimingPointRule struct {
	ID          int
	DecoderID   string
	DecoderName string
	LoopID      byte
	ChannelID   byte
	OrderID     int
	MinTime     decimal.Decimal
	MaxTime     decimal.Decimal
	OrderPos    int
	TimingPoint string
}

// TimingPoint describes the internal go model
type TimingPoint struct {
	Name           string
	Type           int
	DDT            int
	IgnoreIfTimeIn int
	IgnoreBefore   decimal.Decimal
	IgnoreAfter    decimal.Decimal
	SubtractT0     int
	IgnorePS       int
	Position       string
	OrderPos       int
	Color          string
}

type UserDefinedField struct {
	Name       string
	Expression string
	Note       string
}

type VoucherType byte

const (
	VoucherTypeAmount   VoucherType = 0
	VoucherTypePercent  VoucherType = 1
	VoucherTypeFirstReg VoucherType = 2
	VoucherTypePrevReg  VoucherType = 3
)

// Voucher describes the internal go model
type Voucher struct {
	ID         int
	Code       string
	Type       VoucherType
	Amount     decimal.Decimal
	Tax        decimal.Decimal
	Contest    int
	Category   string
	ValidUntil datetime.DateTime
	ValidFrom  datetime.DateTime
	Reusable   int
	UseCounter int
	Remark     string
	OrderPos   float64
}

type PassingToProcess struct {
	Bib         int
	TimingPoint string
	ResultID    int
	Time        decimal.Decimal
	InfoText    string
	Passing     Passing
}

// Passing is a passing for the times/add function
type Passing struct {
	Transponder   string
	Position      PassingPosition
	Hits          int
	RSSI          int
	Battery       decimal.Decimal
	Temperature   int
	WakeupCounter int `json:"WUC"`
	LoopID        byte
	Channel       byte
	InternalData  string
	StatusFlags   int
	DeviceID      string
	DeviceName    string
	OrderID       int
	Port          int
	IsMarker      bool
	FileNo        int
	PassingNo     int
	Customer      int
	Received      time.Time
	UTCTime       time.Time
}

type PassingPosition struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Flag      string
}

type TimesAddResponseItem struct {
	Status      int
	Time        decimal.Decimal
	ResultID    int
	ResultName  string
	RawDataID   int
	TimingPoint string
	Fields      variant.VariantMap
}

// ForwardingInfo contains statistics about the backup/forwarding
type ForwardingInfo struct {
	BytesSent     int
	BytesReceived int
}

type ChatMessage struct {
	ID       int    `json:"i" xml:"i"`
	UserName string `json:"u" xml:"u"`
	Date     string `json:"d" xml:"d"`
	Message  string `json:"m" xml:"m"`
}

type Version struct {
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
	Revision int    `json:"revision"`
	Tag      string `json:"tag"`
	Hash     string `json:"hash"`
}

type ChipFileEntry struct {
	Transponder    string
	Identification string
}

// TriggerType defines the type of a trigger
type TriggerType int

// TriggerType constants
const (
	TriggerNewTime          TriggerType = 0
	TriggerNewParticipant   TriggerType = 1
	TriggerNewChatMessage   TriggerType = 3
	TriggerNewRawData       TriggerType = 4
	TriggerExporter         TriggerType = 5
	TriggerNewRawDataSimple TriggerType = 6
	TriggerNewSplit         TriggerType = 7
	TriggerModJobIDSimple   TriggerType = 8
	TriggerSettingValue     TriggerType = 9
	TriggerNewRawDataV2     TriggerType = 10
)

type SaveValueArrayItem struct {
	Bib       int
	PID       int
	FieldName string
	Value     variant.Variant
}

// EntryFeeItem is an entry fee item which is charged to a participant
type EntryFeeItem struct {
	ID             int
	Name           string
	Fee            decimal.Decimal
	Field          string
	Tax            decimal.Decimal
	Multiplication decimal.Decimal
}

type ImportResult struct {
	Added   int
	Updated int
	PIDs    []int
}

type UserInfo struct {
	CustNo   int
	UserName string
	UserPic  string
}

type UserRights map[string][]string

type UserRole struct {
	ID       int
	UserID   int
	Name     string
	Rights   UserRights
	OrderPos int
}

type UserRight struct {
	UserID   int
	UserName string
	UserPic  string
	Rights   string
	RoleID   int
}

// SimpleAPIItem represents one entry in the SimpleAPI table
type SimpleAPIItem struct {
	Disabled bool
	Key      string
	URL      string
	Label    string
}

type WebHookType int

const (
	WebHookTypeParticipantNew = iota
	WebHookTypeParticipantUpdated
	WebHookTypeRawDataNew
	WebHookTypeModJobID
	WebHookTypeModJobIDSettings
)

type WebHook struct {
	ID       int
	Disabled bool
	Name     string
	Type     WebHookType
	URL      string
	Fields   []string
	Filter   string
	OrderPos int
}

type WebHookMessage struct {
	EventID   string
	WebHookID int
	TimeStamp time.Time
	Values    map[string]variant.Variant
}

type ContestStatisticsResult struct {
	ID        int
	Male      int
	Female    int
	IsFormula bool `json:"isFormula"`
}
type ContestStatistics struct {
	ID       int
	Name     string
	Male     int
	Female   int
	Adults   int
	Children int
	MeanAge  float64
	Finished int
	Results  []ContestStatisticsResult
}

type GroupTimes struct {
	Mode      string
	WaveField string
	Items     []GroupTime
}

type GroupTime struct {
	ID    interface{}
	Time  decimal.Decimal
	Item  interface{} `json:",omitempty"`
	Count int
}

type RegistrationRequest struct {
	DataToken    string
	PaymentToken string
	RegName      string
	Records      []RegistrationRequestRecord
}

type RegistrationRequestRecord struct {
	PID         int
	Record      variant.VariantMap
	Expressions []string
}

type RegistrationRequestResponse struct {
	Records      []RegistrationRequestResponseRecord
	PaymentTerms *invoice.PaymentTerms
}

type RegistrationRequestResponseRecord struct {
	Expressions  variant.VariantMap
	EntryFees    []EntryFeeItem
	HasDuplicate bool
}

type OPData struct {
	EntryFees decimal.Decimal
	Processed decimal.Decimal
	Received  decimal.Decimal
	Date      time.Time
	Payments  []OPDataPayment
}
type OPDataPayment struct {
	ID        int
	Created   time.Time
	Method    int
	Reference string
	Amount    decimal.Decimal
	Currency  string
	Received  decimal.Decimal
	Comment   string `json:",omitempty"`
}
