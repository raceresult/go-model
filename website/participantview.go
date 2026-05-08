package website

import (
	"encoding/json"

	"github.com/raceresult/go-model/datetime"
)

type Element struct {
	ID          string
	Type        string // text, html, picture, columns, field, box, tabs, splits, legs, links, list, photos, certificates, comments
	Active      string // no, onlypopup,onlystandalone, yes
	EnabledFrom datetime.DateTime
	EnabledTo   datetime.DateTime
	ShowIf      string
	Styles      []Style
	ClassName   string
	Children    []*Element
	Config      json.RawMessage
}

type Style struct {
	Attribute string
	Value     string
}

type ElementBoxConfig struct {
	Label       string
	LabelPreset string
	TitleStyles []Style
}
type ElementTextConfig struct {
	Text      string
	Alignment string // left, center, right
}

type ElementHTMLConfig struct {
	Text string
}

type ElementPictureConfig struct {
	Src       string
	Alignment string // left, center, right
}

type ElementFieldConfig struct {
	Field       string
	DisplayMode string // fieldonly -> "", badge, titleabove
	Title       string
	TitleStyles []Style // for titleabove
	Alignment   string  // left, center, right
}

type ElementSplitsConfig struct {
	ShowTOD          bool
	ShowGunTime      bool
	ShowChipTime     bool
	ShowSectorTime   bool
	ShowPace         bool
	ShowOverallRank  bool
	ShowGenderRank   bool
	ShowAgeGroupRank bool
	ShowMaxRank      bool
}
type ElementLegsConfig struct {
	ShowOverallRank  bool
	ShowGenderRank   bool
	ShowAgeGroupRank bool
	ShowMaxRank      bool
	ShowPace         bool
}
type ElementLinksConfig struct {
	Links []ConfigLink
}
type ElementCertificatesConfig struct {
	CertificateSets []PublishedCertificateSet
}
type ElementPhotosConfig struct {
	PortalPhotographer    string
	PhotographerEventID   string
	PortalPhotographerBib string
}

type ElementListConfig struct {
	List string
}

type FavoriteConfig struct {
	Mode      string // "" standard with text, "notext": without text
	Alignment string // left, center, right
}
