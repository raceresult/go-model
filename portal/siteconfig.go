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
	Type        string // "text", "iframe", "registration", "lists", "reviews", "contact", "details"
	Config      json.RawMessage
}

// Registration Tab
type TabRegistrationConfig struct {
	InfoText          string // HTML, plain Text should already be HTML here
	RegTestModeUntil  datetime.DateTime
	RegistrationLogin struct {
		Enabled        bool
		InfoText       string
		LoginNameField string
		Fields         []struct {
			Name  string
			Label string
		}
	}
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
	Live     interface{}
	Sortable interface{}
	Leader   interface{}
	Details  string
	ID       string
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
