package website

import (
	"encoding/json"

	"github.com/raceresult/go-model/datetime"
)

type Element struct {
	ID            string
	Type          string
	Active        string
	EnabledFrom   datetime.DateTime
	EnabledTo     datetime.DateTime
	ShowIf        string
	Styles        []Style
	DynamicFormat string
	Children      []*Element
	Config        json.RawMessage
}

// Element.Type values
const (
	ElementTypeText         = "text"
	ElementTypeHTML         = "html"
	ElementTypePicture      = "picture"
	ElementTypeColumns      = "columns"
	ElementTypeField        = "field"
	ElementTypeBox          = "box"
	ElementTypeTabs         = "tabs"
	ElementTypeSplits       = "splits"
	ElementTypeLegs         = "legs"
	ElementTypeLinks        = "links"
	ElementTypeList         = "list"
	ElementTypePhotos       = "photos"
	ElementTypeCertificates = "certificates"
	ElementTypeChart        = "chart"
	ElementTypeInlineBlock  = "inlineblock"
	ElementTypeMap          = "map"
)

// Element.Active values
const (
	ElementActiveNo             = "no"
	ElementActiveOnlyPopup      = "onlypopup"
	ElementActiveOnlyStandalone = "onlystandalone"
	ElementActiveYes            = "yes"
)

type Style struct {
	Attribute string
	Value     string
}

type ElementBoxConfig struct {
	Label              string
	LabelPreset        string
	TitleStyles        []Style
	TitleDynamicFormat string
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
	Field              string
	DisplayMode        string
	Title              string
	TitleStyles        []Style // for titleabove
	TitleDynamicFormat string
	Alignment          int
}

// ElementFieldConfig.DisplayMode values
const (
	FieldDisplayModeFieldOnly  = ""
	FieldDisplayModeBadge      = "badge"
	FieldDisplayModeTitleAbove = "titleabove"
)

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
	HidePrediction   bool
}
type ElementLegsConfig struct {
	EmbedInBox       bool
	ShowOverallRank  bool
	ShowGenderRank   bool
	ShowAgeGroupRank bool
	ShowMaxRank      bool
	ShowPace         bool
	HidePrediction   bool
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

type ElementInlineblockConfig struct {
	Alignment int
}

type ElementChartConfig struct {
	EmbedInBox bool
	ChartType  string
	RankScope  string `json:",omitempty"`
}

// ElementChartConfig.ChartType values
const (
	ChartTypeSplitPace = "splitpace"
	ChartTypeSplitRank = "splitrank"
)

// ElementChartConfig.RankScope values
const (
	RankScopeOverall  = ""
	RankScopeGender   = "gender"
	RankScopeAgeGroup = "agegroup"
)

type ElementMapConfig struct {
	ShowFavorites  bool   `json:"MapShowFavorites"`
	MarkerLabels   string `json:"MapMarkerLabels"`
	MarkerTooltips string `json:"MapMarkerTooltips"`
	MarkerFilter   string `json:"MapMarkerFilter"`
	MarkerWait     bool   `json:"MapMarkerWait"`
	MarkerHideIcon bool   `json:"MapMarkerHideIcon"`
	EmbedInBox     bool   `json:"EmbedInBox"`
}

type MapData map[int]*MapContest

type MapContest struct {
	Color        string `json:",omitempty"`
	Splits       []*MapContestSplit
	Participants MapParticipants
}

type MapContestSplit struct {
	Name     string
	Distance int    // always in meters
	Position string `json:",omitempty"`
	Internal bool   `json:",omitempty"`
	ID       int    `json:"-"`
	OrderPos int    `json:"-"`
}

type MapParticipants struct {
	Columns []string
	Data    map[int][]interface{}
}
