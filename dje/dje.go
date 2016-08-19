package dje

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type dje struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	decimal            []byte
	group              []byte
	minus              []byte
	percent            []byte
	perMille           []byte
	timeSeparator      []byte
	inifinity          []byte
	currencies         [][]byte // idx = enum of currency code
	monthsAbbreviated  [][]byte
	monthsNarrow       [][]byte
	monthsWide         [][]byte
	daysAbbreviated    [][]byte
	daysNarrow         [][]byte
	daysShort          [][]byte
	daysWide           [][]byte
	periodsAbbreviated [][]byte
	periodsNarrow      [][]byte
	periodsShort       [][]byte
	periodsWide        [][]byte
	erasAbbreviated    [][]byte
	erasNarrow         [][]byte
	erasWide           [][]byte
	timezones          map[string][]byte
}

// New returns a new instance of translator for the 'dje' locale
func New() locales.Translator {
	return &dje{
		locale:             "dje",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		decimal:            []byte{0x2e},
		group:              []byte{0xc2, 0xa0},
		minus:              []byte{},
		percent:            []byte{},
		perMille:           []byte{},
		timeSeparator:      []byte{0x3a},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0xc5, 0xbd, 0x61, 0x6e}, {0x46, 0x65, 0x65}, {0x4d, 0x61, 0x72}, {0x41, 0x77, 0x69}, {0x4d, 0x65}, {0xc5, 0xbd, 0x75, 0x77}, {0xc5, 0xbd, 0x75, 0x79}, {0x55, 0x74}, {0x53, 0x65, 0x6b}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x6f}, {0x44, 0x65, 0x65}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0xc5, 0xbd}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0xc5, 0xbd}, {0xc5, 0xbd}, {0x55}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:         [][]uint8{[]uint8(nil), {0xc5, 0xbd, 0x61, 0x6e, 0x77, 0x69, 0x79, 0x65}, {0x46, 0x65, 0x65, 0x77, 0x69, 0x72, 0x69, 0x79, 0x65}, {0x4d, 0x61, 0x72, 0x73, 0x69}, {0x41, 0x77, 0x69, 0x72, 0x69, 0x6c}, {0x4d, 0x65}, {0xc5, 0xbd, 0x75, 0x77, 0x65, 0xc5, 0x8b}, {0xc5, 0xbd, 0x75, 0x79, 0x79, 0x65}, {0x55, 0x74}, {0x53, 0x65, 0x6b, 0x74, 0x61, 0x6e, 0x62, 0x75, 0x72}, {0x4f, 0x6b, 0x74, 0x6f, 0x6f, 0x62, 0x75, 0x72}, {0x4e, 0x6f, 0x6f, 0x77, 0x61, 0x6e, 0x62, 0x75, 0x72}, {0x44, 0x65, 0x65, 0x73, 0x61, 0x6e, 0x62, 0x75, 0x72}},
		daysAbbreviated:    [][]uint8{{0x41, 0x6c, 0x68}, {0x41, 0x74, 0x69}, {0x41, 0x74, 0x61}, {0x41, 0x6c, 0x61}, {0x41, 0x6c, 0x6d}, {0x41, 0x6c, 0x7a}, {0x41, 0x73, 0x69}},
		daysNarrow:         [][]uint8{{0x48}, {0x54}, {0x54}, {0x4c}, {0x4d}, {0x5a}, {0x53}},
		daysWide:           [][]uint8{{0x41, 0x6c, 0x68, 0x61, 0x64, 0x69}, {0x41, 0x74, 0x69, 0x6e, 0x6e, 0x69}, {0x41, 0x74, 0x61, 0x6c, 0x61, 0x61, 0x74, 0x61}, {0x41, 0x6c, 0x61, 0x72, 0x62, 0x61}, {0x41, 0x6c, 0x68, 0x61, 0x6d, 0x69, 0x73, 0x69}, {0x41, 0x6c, 0x7a, 0x75, 0x6d, 0x61}, {0x41, 0x73, 0x69, 0x62, 0x74, 0x69}},
		periodsAbbreviated: [][]uint8{{0x53, 0x75, 0x62, 0x62, 0x61, 0x61, 0x68, 0x69}, {0x5a, 0x61, 0x61, 0x72, 0x69, 0x6b, 0x61, 0x79, 0x20, 0x62}},
		periodsWide:        [][]uint8{{0x53, 0x75, 0x62, 0x62, 0x61, 0x61, 0x68, 0x69}, {0x5a, 0x61, 0x61, 0x72, 0x69, 0x6b, 0x61, 0x79, 0x20, 0x62}},
		erasAbbreviated:    [][]uint8{{0x49, 0x4a}, {0x49, 0x5a}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{{0x49, 0x73, 0x61, 0x61, 0x20, 0x6a, 0x69, 0x6e, 0x65}, {0x49, 0x73, 0x61, 0x61, 0x20, 0x7a, 0x61, 0x6d, 0x61, 0x6e, 0x6f, 0x6f}},
		timezones:          map[string][]uint8{"LHST": {0x4c, 0x48, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "ART": {0x41, 0x52, 0x54}, "IST": {0x49, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "OEZ": {0x4f, 0x45, 0x5a}, "VET": {0x56, 0x45, 0x54}, "SGT": {0x53, 0x47, 0x54}, "WAT": {0x57, 0x41, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "AWST": {0x41, 0x57, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "MST": {0x4d, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WIT": {0x57, 0x49, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "COT": {0x43, 0x4f, 0x54}, "EST": {0x45, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "PST": {0x50, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "BT": {0x42, 0x54}, "CST": {0x43, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "EAT": {0x45, 0x41, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "CAT": {0x43, 0x41, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (dje *dje) Locale() string {
	return dje.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'dje'
func (dje *dje) PluralsCardinal() []locales.PluralRule {
	return dje.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'dje'
func (dje *dje) PluralsOrdinal() []locales.PluralRule {
	return dje.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'dje'
func (dje *dje) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'dje'
func (dje *dje) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'dje'
func (dje *dje) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (dje *dje) MonthAbbreviated(month time.Month) []byte {
	return dje.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (dje *dje) MonthsAbbreviated() [][]byte {
	return dje.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (dje *dje) MonthNarrow(month time.Month) []byte {
	return dje.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (dje *dje) MonthsNarrow() [][]byte {
	return dje.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (dje *dje) MonthWide(month time.Month) []byte {
	return dje.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (dje *dje) MonthsWide() [][]byte {
	return dje.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (dje *dje) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return dje.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (dje *dje) WeekdaysAbbreviated() [][]byte {
	return dje.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (dje *dje) WeekdayNarrow(weekday time.Weekday) []byte {
	return dje.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (dje *dje) WeekdaysNarrow() [][]byte {
	return dje.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (dje *dje) WeekdayShort(weekday time.Weekday) []byte {
	return dje.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (dje *dje) WeekdaysShort() [][]byte {
	return dje.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (dje *dje) WeekdayWide(weekday time.Weekday) []byte {
	return dje.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (dje *dje) WeekdaysWide() [][]byte {
	return dje.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'dje' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(dje.decimal) + len(dje.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dje.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(dje.group) - 1; j >= 0; j-- {
					b = append(b, dje.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(dje.minus) - 1; j >= 0; j-- {
			b = append(b, dje.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'dje' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (dje *dje) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(dje.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dje.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(dje.minus) - 1; j >= 0; j-- {
			b = append(b, dje.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, dje.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dje.currencies[currency]
	l := len(s) + len(dje.decimal) + len(dje.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dje.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(dje.group) - 1; j >= 0; j-- {
					b = append(b, dje.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(dje.minus) - 1; j >= 0; j-- {
			b = append(b, dje.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dje.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'dje'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dje.currencies[currency]
	l := len(s) + len(dje.decimal) + len(dje.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dje.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(dje.group) - 1; j >= 0; j-- {
					b = append(b, dje.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(dje.minus) - 1; j >= 0; j-- {
			b = append(b, dje.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dje.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, dje.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, dje.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, dje.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, dje.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'dje'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (dje *dje) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dje.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := dje.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
