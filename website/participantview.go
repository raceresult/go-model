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
	Alignment int
}

type ElementHTMLConfig struct {
	Text string
}

type ElementPictureConfig struct {
	Src       string
	Alignment int
}

type ElementFieldConfig struct {
	Field       string
	DisplayMode string // fieldonly -> "", badge, titleabove
	Title       string
	TitleStyles []Style // for titleabove
	Alignment   int
}

type ElementSplitsConfig struct {
	EmbedInBox       bool
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
	EmbedInBox       bool
	ShowOverallRank  bool
	ShowGenderRank   bool
	ShowAgeGroupRank bool
	ShowMaxRank      bool
	ShowPace         bool
}
type ElementLinksConfig struct {
	EmbedInBox bool
	Links      []ConfigLink
}
type ElementCertificatesConfig struct {
	EmbedInBox      bool
	CertificateSets []PublishedCertificateSet
}
type ElementPhotosConfig struct {
	EmbedInBox            bool
	PortalPhotographer    string
	PhotographerEventID   string
	PortalPhotographerBib string
}

type ElementListConfig struct {
	EmbedInBox bool
	List       string
}

type ElementCommentsConfig struct {
	EmbedInBox bool
}

type ElementFavoriteConfig struct {
	Mode      string // "" standard with text, "notext": without text
	Alignment int
}
