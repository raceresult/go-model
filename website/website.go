package website

import (
	"encoding/json"

	"github.com/raceresult/go-model/datetime"
)

const WebsiteDefault = "default"

type Website struct {
	Label              string
	Publish            bool
	CoverPicture       string // URL of the cover picture for the event site
	CoverPictureMobile string // URL of the cover picture optimized for mobile devices
	HideEventLogo      bool   // Whether to hide the event logo on the site
	AdditionalCode     string // Additional HTML or JavaScript code to be included in the site
	AdditionalFeatures []AdditionalFeature
	BrandColorDark     string
	ApplyContestColor  string
	PortalTestKey      string
	CommentsMinimumAge int      // 0-21 // Minimum age required to enable comments on the participant view
	Organizer          Company  // Information about the event organizer
	Timer              Company  // Information about the timekeeper of the event
	PayProc            Company  // Information about the payment processor for the event
	Tabs               []*Tab   // List of tabs to be displayed on the event site, including their properties, as represented in the settings
	HomeTab            *HomeTab // Information about the home tab of the event site, including its properties
}

const ApplyContestColorYes = "yes"

type HomeTab struct {
	// Info Tab
	InfoText string // HTML, plain Text should already be HTML here
	// Sidebar
	ShowEventGroups     bool
	ExternalWebsiteLink string
	// Forum Tab
	ForumDeactivated bool
	// Contact Tab
	ComplaintsDeactivated bool
	ComplaintsEmail       string
}

type AdditionalFeature struct {
	Type   AdditionalFeatureType
	ID     string
	Active bool
}

type AdditionalFeatureType string

// AdditionalFeature.Type values
const (
	AdditionalFeatureTypeNone            AdditionalFeatureType = ""
	AdditionalFeatureTypeGoogleAnalytics AdditionalFeatureType = "GoogleAnalytics"
	AdditionalFeatureTypeMetaPixel       AdditionalFeatureType = "MetaPixel"
)

type Company struct {
	Name           string
	Mail           string
	Web            string
	Picture        string
	DataProtection string
}

type Tab struct {
	Label       string
	URLName     string
	Enabled     bool
	ShowInMenu  bool
	ShowInMenu2 ShowInMenu
	ActiveFrom  datetime.DateTime
	ActiveUntil datetime.DateTime
	Type        TabType
	Config      json.RawMessage
}

type TabType string

// Tab.Type values
const (
	TabTypeText            TabType = "text"
	TabTypeExternalContent TabType = "externalcontent"
	TabTypeRegistration    TabType = "registration"
	TabTypeLists           TabType = "lists"
	TabTypeFavorites       TabType = "favorites"
	TabTypeParticipantView TabType = "participantview"

	// legacy
	TabTypeContact TabType = "contact"
	TabTypeDetails TabType = "details"
	TabTypeReviews TabType = "reviews"
)

type ShowInMenu string

// Tab.ShowInMenu2 values
const (
	ShowInMenuDefault    ShowInMenu = "" // fall back to Tab.ShowInMenu
	ShowInMenuAlways     ShowInMenu = "always"
	ShowInMenuNever      ShowInMenu = "never"
	ShowInMenuWhenActive ShowInMenu = "whenactive"
)

func (q *Tab) IsWithinActiveTimeFrame(t datetime.DateTime) bool {
	if !q.ActiveFrom.IsZero() && t.Before(q.ActiveFrom) {
		return false
	}
	if !q.ActiveUntil.IsZero() && t.After(q.ActiveUntil) {
		return false
	}
	return true
}

func (q *Tab) IsVisibleInMenu(t datetime.DateTime) bool {
	switch q.ShowInMenu2 {
	case ShowInMenuAlways:
		return true
	case ShowInMenuNever:
		return false
	case ShowInMenuWhenActive:
		return q.IsWithinActiveTimeFrame(t)
	default:
		return q.ShowInMenu
	}
}

// Text Tab
type TabTextConfig struct {
	InfoText                string // HTML, plain Text should already be HTML here
	ShowEventGroups         bool
	ShowEventNews           bool
	ShowExternalWebsiteLink bool
	ExternalWebsiteLink     string
}

// ExternalContent Tab
type TabExternalContentConfig struct {
	Data                    string // HTML, plain Text should already be HTML here
	ShowEventGroups         bool
	ShowEventNews           bool
	ShowExternalWebsiteLink bool
	ExternalWebsiteLink     string
	SubType                 string
	Config                  json.RawMessage
}

// TabExternalContentConfig.SubType values
const (
	ExtContentSubTypeIFrame  = ""
	ExtContentSubTypeYoutube = "youtube"
)

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
	Subtype              ListsSubtype
	StandardDetails      string
	Lists                []ConfigList
	LeaderboardHideCount bool
}

type ListsSubtype string

// TabListsConfig.Subtype values
const (
	ListsSubtypeBoxes    ListsSubtype = "boxes"
	ListsSubtypeDropdown ListsSubtype = "dropdown"
)

type ConfigList struct {
	Name                string
	Mode                string
	Contest             interface{}
	ShowAs              string
	Format              string
	Live                interface{} `json:",omitempty"`
	Sortable            interface{} `json:",omitempty"`
	Leader              interface{} `json:",omitempty"`
	Details             string      `json:",omitempty"`
	ID                  string      `json:",omitempty"`
	HideProfilePictures bool        `json:",omitempty"`
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

// Participantsview Tab
type TabParticipantviewConfig struct {
	LiveUpdate bool
	Elements   []*Element
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

// Favorites Tab
type TabFavoritesConfig struct {
	Participantview      string
	Search               SearchConfig
	SearchResultElements []*Element
	FavoriteElements     []*Element
	ListSortedField      string
	ListSortAsc          bool
}

type SearchConfig struct {
	Filter      string
	Fields      []string
	SortedField string
	SortAsc     bool
}
