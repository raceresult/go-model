package website

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/raceresult/go-model/datetime"
)

type Element struct {
	ID            string
	Type          string // text, html, picture, columns, field, box, tabs, splits, legs, links, list, photos, certificates, comments
	Active        string // no, onlypopup,onlystandalone, yes
	EnabledFrom   datetime.DateTime
	EnabledTo     datetime.DateTime
	ShowIf        string
	Styles        []Style
	DynamicFormat string
	Children      []*Element
	Config        json.RawMessage
}

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
	DisplayMode        string // fieldonly -> "", badge, titleabove
	Title              string
	TitleStyles        []Style // for titleabove
	TitleDynamicFormat string
	Alignment          int
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

type ElementCommentsConfig struct {
	EmbedInBox bool
}

type ElementFavoriteConfig struct {
	Mode      string // "" standard with text, "notext": without text
	Alignment int
}

type ElementInlineblockConfig struct {
	Alignment int
}

type ElementMapConfig struct {
	ShowFavorites  bool   `json:"MapShowFavorites"`
	MarkerLabels   string `json:"MapMarkerLabels"`
	MarkerTooltips string `json:"MapMarkerTooltips"`
	MarkerFilter   string `json:"MapMarkerFilter"`
	MarkerWait     bool   `json:"MapMarkerWait"`
	EmbedInBox     bool   `json:"EmbedInBox"`
}

type MapData map[int]MapContest

type MapContest struct {
	Color        string `json:",omitempty"`
	Splits       []MapContestSplit
	Participants MapParticipants
}

type MapContestSplit struct {
	Distance int    // always in meters
	Position string `json:",omitempty"`
	Internal bool   `json:",omitempty"`
}

type MapParticipants map[int]MapParticipant

type MapParticipant struct {
	Label   string
	Tooltip string
	Splits  []*MapParticipantSplit
}

type MapParticipantSplit struct {
	Exists bool
	Time   int
}

// MarshalJSON implements the json.Marshaler interface.
func (m MapParticipants) MarshalJSON() ([]byte, error) {
	// Label/Tooltip columns exist if any participant uses them
	haveLabel := false
	haveTooltip := false
	for _, p := range m {
		haveLabel = haveLabel || p.Label != ""
		haveTooltip = haveTooltip || p.Tooltip != ""
		if haveLabel && haveTooltip {
			break
		}
	}

	// sort IDs for deterministic output
	ids := make([]int, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	// a priori, there are only bad guesses for the size of the resulting buffer
	var buf bytes.Buffer

	// write columns
	buf.WriteString(`{"Columns":["SplitsTime","SplitsExists"`)
	if haveLabel {
		buf.WriteString(`,"Label"`)
	}
	if haveTooltip {
		buf.WriteString(`,"Tooltip"`)
	}

	// write one row per participant
	buf.WriteString(`],"Data":{`)
	for i, id := range ids {
		p := m[id]
		if i > 0 {
			buf.WriteByte(',')
		}

		// key
		buf.WriteByte('"')
		buf.WriteString(strconv.Itoa(id))
		buf.WriteString(`":[[`)

		// SplitsTime
		for j, split := range p.Splits {
			if j > 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(strconv.Itoa(split.Time))
		}

		// SplitsExists
		buf.WriteString("],\"")
		for _, split := range p.Splits {
			if split.Exists {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
		}
		buf.WriteByte('"')

		// Label and Tooltip are arbitrary strings and need proper JSON escaping
		if haveLabel {
			buf.WriteByte(',')
			s, err := json.Marshal(p.Label)
			if err != nil {
				return nil, err
			}
			buf.Write(s)
		}
		if haveTooltip {
			buf.WriteByte(',')
			s, err := json.Marshal(p.Tooltip)
			if err != nil {
				return nil, err
			}
			buf.Write(s)
		}
		buf.WriteByte(']')
	}
	buf.WriteString("}}")
	return buf.Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *MapParticipants) UnmarshalJSON(data []byte) error {
	var mapParticipantJSON struct {
		Columns []string
		Data    map[int][]json.RawMessage
	}
	if err := json.Unmarshal(data, &mapParticipantJSON); err != nil {
		return err
	}

	// map column names to their position in each row
	colIndex := make(map[string]int, len(mapParticipantJSON.Columns))
	for i, c := range mapParticipantJSON.Columns {
		colIndex[c] = i
	}
	cell := func(row []json.RawMessage, col string, dst interface{}) error {
		i, ok := colIndex[col]
		if !ok || i >= len(row) {
			return nil
		}
		return json.Unmarshal(row[i], dst)
	}

	*m = make(MapParticipants, len(mapParticipantJSON.Data))
	for id, row := range mapParticipantJSON.Data {
		var p MapParticipant
		var exists string
		var times []int
		if err := cell(row, "SplitsExists", &exists); err != nil {
			return err
		}
		if err := cell(row, "SplitsTime", &times); err != nil {
			return err
		}
		if err := cell(row, "Label", &p.Label); err != nil {
			return err
		}
		if err := cell(row, "Tooltip", &p.Tooltip); err != nil {
			return err
		}

		p.Splits = make([]*MapParticipantSplit, 0, len(times))
		for j, t := range times {
			p.Splits = append(p.Splits, &MapParticipantSplit{
				Exists: j < len(exists) && exists[j] == '1',
				Time:   t,
			})
		}
		(*m)[id] = p
	}
	return nil
}
