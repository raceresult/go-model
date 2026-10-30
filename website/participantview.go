package website

import (
	"encoding/json"

	"github.com/raceresult/go-model/datetime"
)

type Element struct {
	ID            string
	Type          ElementType
	Active        ElementActive
	EnabledFrom   datetime.DateTime
	EnabledTo     datetime.DateTime
	ShowIf        string
	Styles        []Style
	DynamicFormat string
	Children      []*Element
	Config        json.RawMessage
}

type ElementType string

// Element.Type values
const (
	ElementTypeText         ElementType = "text"
	ElementTypeHTML         ElementType = "html"
	ElementTypePicture      ElementType = "picture"
	ElementTypeColumns      ElementType = "columns"
	ElementTypeField        ElementType = "field"
	ElementTypeBox          ElementType = "box"
	ElementTypeTabs         ElementType = "tabs"
	ElementTypeSplits       ElementType = "splits"
	ElementTypeLegs         ElementType = "legs"
	ElementTypeLinks        ElementType = "links"
	ElementTypeList         ElementType = "list"
	ElementTypePhotos       ElementType = "photos"
	ElementTypeCertificates ElementType = "certificates"
	ElementTypeChart        ElementType = "chart"
	ElementTypeInlineBlock  ElementType = "inlineblock"
	ElementTypeMap          ElementType = "map"

	// legacy
	ElementTypeFavorite ElementType = "favorite"
	ElementTypeComments ElementType = "comments"
)

type ElementActive string

// Element.Active values
const (
	ElementActiveNo             ElementActive = "no"
	ElementActiveOnlyPopup      ElementActive = "onlypopup"
	ElementActiveOnlyStandalone ElementActive = "onlystandalone"
	ElementActiveYes            ElementActive = "yes"
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
	DisplayMode        FieldDisplayMode
	Title              string
	TitleStyles        []Style // for titleabove
	TitleDynamicFormat string
	Alignment          int
}

type FieldDisplayMode string

// ElementFieldConfig.DisplayMode values
const (
	FieldDisplayModeFieldOnly  FieldDisplayMode = ""
	FieldDisplayModeBadge      FieldDisplayMode = "badge"
	FieldDisplayModeTitleAbove FieldDisplayMode = "titleabove"
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
	ChartType  ChartType
	RankScope  RankScope `json:",omitempty"`
}

type ChartType string

// ElementChartConfig.ChartType values
const (
	ChartTypeSplitPace ChartType = "splitpace"
	ChartTypeSplitRank ChartType = "splitrank"
)

type RankScope string

// ElementChartConfig.RankScope values
const (
	RankScopeOverall  RankScope = ""
	RankScopeGender   RankScope = "gender"
	RankScopeAgeGroup RankScope = "agegroup"
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
