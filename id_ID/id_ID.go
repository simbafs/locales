package id_ID

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type id_ID struct {
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

// New returns a new instance of translator for the 'id_ID' locale
func New() locales.Translator {
	return &id_ID{
		locale:             "id_ID",
		pluralsCardinal:    []locales.PluralRule{6},
		pluralsOrdinal:     []locales.PluralRule{6},
		decimal:            []byte{0x2c},
		group:              []byte{0x2e},
		minus:              []byte{0x2d},
		percent:            []byte{0x25},
		perMille:           []byte{0xe2, 0x80, 0xb0},
		timeSeparator:      []byte{0x2e},
		inifinity:          []byte{0xe2, 0x88, 0x9e},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x46, 0x65, 0x62}, {0x4d, 0x61, 0x72}, {0x41, 0x70, 0x72}, {0x4d, 0x65, 0x69}, {0x4a, 0x75, 0x6e}, {0x4a, 0x75, 0x6c}, {0x41, 0x67, 0x74}, {0x53, 0x65, 0x70}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x76}, {0x44, 0x65, 0x73}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:         [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x69}, {0x46, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72, 0x69}, {0x4d, 0x61, 0x72, 0x65, 0x74}, {0x41, 0x70, 0x72, 0x69, 0x6c}, {0x4d, 0x65, 0x69}, {0x4a, 0x75, 0x6e, 0x69}, {0x4a, 0x75, 0x6c, 0x69}, {0x41, 0x67, 0x75, 0x73, 0x74, 0x75, 0x73}, {0x53, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x4f, 0x6b, 0x74, 0x6f, 0x62, 0x65, 0x72}, {0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x44, 0x65, 0x73, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:    [][]uint8{{0x4d, 0x69, 0x6e}, {0x53, 0x65, 0x6e}, {0x53, 0x65, 0x6c}, {0x52, 0x61, 0x62}, {0x4b, 0x61, 0x6d}, {0x4a, 0x75, 0x6d}, {0x53, 0x61, 0x62}},
		daysNarrow:         [][]uint8{{0x4d}, {0x53}, {0x53}, {0x52}, {0x4b}, {0x4a}, {0x53}},
		daysShort:          [][]uint8{{0x4d, 0x69, 0x6e}, {0x53, 0x65, 0x6e}, {0x53, 0x65, 0x6c}, {0x52, 0x61, 0x62}, {0x4b, 0x61, 0x6d}, {0x4a, 0x75, 0x6d}, {0x53, 0x61, 0x62}},
		daysWide:           [][]uint8{{0x4d, 0x69, 0x6e, 0x67, 0x67, 0x75}, {0x53, 0x65, 0x6e, 0x69, 0x6e}, {0x53, 0x65, 0x6c, 0x61, 0x73, 0x61}, {0x52, 0x61, 0x62, 0x75}, {0x4b, 0x61, 0x6d, 0x69, 0x73}, {0x4a, 0x75, 0x6d, 0x61, 0x74}, {0x53, 0x61, 0x62, 0x74, 0x75}},
		periodsAbbreviated: [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:      [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:        [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:    [][]uint8{{0x53, 0x4d}, {0x4d}},
		erasNarrow:         [][]uint8{{0x53, 0x4d}, {0x4d}},
		erasWide:           [][]uint8{{0x53, 0x65, 0x62, 0x65, 0x6c, 0x75, 0x6d, 0x20, 0x4d, 0x61, 0x73, 0x65, 0x68, 0x69}, {0x4d}},
		timezones:          map[string][]uint8{"GYT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "HKST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "SRT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "SAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x65, 0x6c, 0x61, 0x74, 0x61, 0x6e}, "ChST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "JST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x4a, 0x65, 0x70, 0x61, 0x6e, 0x67}, "LHDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "ACDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "MESZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "HAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "VET": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "WAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "ECT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x45, 0x6b, 0x75, 0x61, 0x64, 0x6f, 0x72}, "ACWST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "WARST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x42, 0x61, 0x67, 0x69, 0x61, 0x6e, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "ARST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "OESZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "WEZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "CAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "∅∅∅": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x63, 0x72, 0x65}, "WART": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x42, 0x61, 0x67, 0x69, 0x61, 0x6e, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "AWDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "CDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "GFT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x50, 0x72, 0x61, 0x6e, 0x63, 0x69, 0x73}, "WIT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "CLST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x43, 0x69, 0x6c, 0x65}, "SGT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x61}, "UYT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "ADT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b}, "MEZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "LHST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "MDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x50, 0x65, 0x67, 0x75, 0x6e, 0x75, 0x6e, 0x67, 0x61, 0x6e}, "PST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x50, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b}, "AEDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "BOT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "UYST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "GMT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x52, 0x61, 0x74, 0x61, 0x2d, 0x72, 0x61, 0x74, 0x61, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "TMT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "NZDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x53, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x61, 0x20, 0x42, 0x61, 0x72, 0x75}, "AWST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "WITA": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "MYT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "PDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x50, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b}, "CST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "EAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "HADT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "HKT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "WAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "AKST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "WIB": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "IST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "HNT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "ACST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "ART": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "AEST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "WESZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "EST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "EDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "HAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "BT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "AST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b}, "ACWDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "COST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "JDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x4a, 0x65, 0x70, 0x61, 0x6e, 0x67}, "MST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x50, 0x65, 0x67, 0x75, 0x6e, 0x75, 0x6e, 0x67, 0x61, 0x6e}, "CHADT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "TMST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "CLT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x43, 0x69, 0x6c, 0x65}, "OEZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "COT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "NZST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x53, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x61, 0x20, 0x42, 0x61, 0x72, 0x75}, "CHAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "AKDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}},
	}
}

// Locale returns the current translators string locale
func (id *id_ID) Locale() string {
	return id.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'id_ID'
func (id *id_ID) PluralsCardinal() []locales.PluralRule {
	return id.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'id_ID'
func (id *id_ID) PluralsOrdinal() []locales.PluralRule {
	return id.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'id_ID'
func (id *id_ID) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'id_ID'
func (id *id_ID) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'id_ID'
func (id *id_ID) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (id *id_ID) MonthAbbreviated(month time.Month) []byte {
	return id.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (id *id_ID) MonthsAbbreviated() [][]byte {
	return id.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (id *id_ID) MonthNarrow(month time.Month) []byte {
	return id.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (id *id_ID) MonthsNarrow() [][]byte {
	return id.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (id *id_ID) MonthWide(month time.Month) []byte {
	return id.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (id *id_ID) MonthsWide() [][]byte {
	return id.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (id *id_ID) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return id.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (id *id_ID) WeekdaysAbbreviated() [][]byte {
	return id.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (id *id_ID) WeekdayNarrow(weekday time.Weekday) []byte {
	return id.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (id *id_ID) WeekdaysNarrow() [][]byte {
	return id.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (id *id_ID) WeekdayShort(weekday time.Weekday) []byte {
	return id.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (id *id_ID) WeekdaysShort() [][]byte {
	return id.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (id *id_ID) WeekdayWide(weekday time.Weekday) []byte {
	return id.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (id *id_ID) WeekdaysWide() [][]byte {
	return id.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'id_ID' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(id.decimal) + len(id.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, id.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, id.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'id_ID' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (id *id_ID) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(id.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, id.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, id.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := id.currencies[currency]
	l := len(s) + len(id.decimal) + len(id.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, id.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, id.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, id.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'id_ID'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := id.currencies[currency]
	l := len(s) + len(id.decimal) + len(id.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, id.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, id.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, id.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, id.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, id.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, id.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, id.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'id_ID'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (id *id_ID) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := id.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
