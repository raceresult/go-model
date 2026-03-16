package portal

import (
	"encoding/json"

	"github.com/raceresult/go-model/datetime"
)

type SiteConfig struct {
	Label              string
	Publish            bool
	CoverPicture       string // URL of the cover picture for the event site
	CoverPictureMobile string // URL of the cover picture optimized for mobile devices
	HideEventLogo      bool   // Whether to hide the event logo on the site
	AdditionalCode     string // Additional HTML or JavaScript code to be included in the site
	BrandColorDark     string
	PortalTestKey      string
	Organizer          Company      // Information about the event organizer
	Timer              Company      // Information about the timekeeper of the event
	PayProc            Company      // Information about the payment processor for the event
	Tabs               []*TabConfig // List of tabs to be displayed on the event site, including their properties, as represented in the settings
}

type Company struct {
	Name           string
	Mail           string
	Web            string
	Picture        string
	DataProtection string
}

type TabConfig struct {
	Label       string
	URLName     string
	Enabled     bool
	ShowInMenu  bool
	ActiveFrom  datetime.DateTime
	ActiveUntil datetime.DateTime
	Type        string // "text", "externalcontent", "registration", "lists", "reviews", "contact", "details"
	Config      json.RawMessage
}

func (q *TabConfig) IsVisible() bool {
	return q.Enabled && q.ShowInMenu
}

func (q *TabConfig) IsDisabled(t datetime.DateTime) bool {
	if !q.ActiveFrom.IsZero() && t.Before(q.ActiveFrom) {
		return true
	}
	if !q.ActiveUntil.IsZero() && t.After(q.ActiveUntil) {
		return true
	}
	return false
}

// Text Tab
type TabTextConfig struct {
	InfoText        string // HTML, plain Text should already be HTML here
	ShowEventGroups bool
}

// ExternalContent Tab
type TabExternalContentConfig struct {
	Data            string // HTML, plain Text should already be HTML here
	ShowEventGroups bool
	SubType         string
	Config          json.RawMessage
}

// Tab Contact
type TabContactConfig struct {
	ComplaintsDeactivated bool
	ComplaintsEmail       string
}

// Registration Tab
type TabRegistrationConfig struct {
	InfoText          string // HTML, plain Text should already be HTML here
	RegTestModeUntil  datetime.DateTime
	RegistrationLogin RegistrationLogin
	RegistrationForms []RegConfig
}

type RegConfig struct {
	Form          string
	InfoText      string
	ShowContests  bool
	HideEntryFees bool
	ButtonText    string
	IfLoggedIn    int // 0=always, 1=if logged in, 2=if not logged in
}

// Lists Tab
type TabListsConfig struct {
	InfoText             string // HTML, plain Text should already be HTML here
	Subtype              string // boxes or dropdown
	Lists                []ConfigList
	LeaderboardHideCount bool
	ShowCommentIcon      bool
}

type ConfigList struct {
	Name     string
	Mode     string
	Contest  interface{}
	ShowAs   string
	Format   string
	Live     interface{} `json:",omitempty"`
	Sortable interface{} `json:",omitempty"`
	Leader   interface{} `json:",omitempty"`
	Details  string      `json:",omitempty"`
	ID       string      `json:",omitempty"`
}

// Details Tab
type TabDetailsConfig struct {
	Head                  string
	HeadText              string
	SplitConfig           string
	PortalPhotographer    string
	PhotographerEventID   string
	PortalPhotographerBib string
	CertificateSets       []PublishedCertificateSet
	Links                 []ConfigLink
	CertificatesShowFrom  datetime.DateTime
	CertificatesShowUntil datetime.DateTime
	EnableComments        bool
	List                  string
}

type PublishedCertificateSet struct {
	Mode    string
	Name    string
	ShowAs  string
	Picture string
	Contest interface{}
}

type ConfigLink struct {
	Disabled bool
	URL      string
	Filter   string
	Picture  string
	Label    string
}
