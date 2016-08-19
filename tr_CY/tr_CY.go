package tr_CY

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type tr_CY struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'tr_CY' locale
func New() locales.Translator {
	return &tr_CY{
		locale:                 "tr_CY",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4f, 0x63, 0x61}, {0xc5, 0x9e, 0x75, 0x62}, {0x4d, 0x61, 0x72}, {0x4e, 0x69, 0x73}, {0x4d, 0x61, 0x79}, {0x48, 0x61, 0x7a}, {0x54, 0x65, 0x6d}, {0x41, 0xc4, 0x9f, 0x75}, {0x45, 0x79, 0x6c}, {0x45, 0x6b, 0x69}, {0x4b, 0x61, 0x73}, {0x41, 0x72, 0x61}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4f}, {0xc5, 0x9e}, {0x4d}, {0x4e}, {0x4d}, {0x48}, {0x54}, {0x41}, {0x45}, {0x45}, {0x4b}, {0x41}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4f, 0x63, 0x61, 0x6b}, {0xc5, 0x9e, 0x75, 0x62, 0x61, 0x74}, {0x4d, 0x61, 0x72, 0x74}, {0x4e, 0x69, 0x73, 0x61, 0x6e}, {0x4d, 0x61, 0x79, 0xc4, 0xb1, 0x73}, {0x48, 0x61, 0x7a, 0x69, 0x72, 0x61, 0x6e}, {0x54, 0x65, 0x6d, 0x6d, 0x75, 0x7a}, {0x41, 0xc4, 0x9f, 0x75, 0x73, 0x74, 0x6f, 0x73}, {0x45, 0x79, 0x6c, 0xc3, 0xbc, 0x6c}, {0x45, 0x6b, 0x69, 0x6d}, {0x4b, 0x61, 0x73, 0xc4, 0xb1, 0x6d}, {0x41, 0x72, 0x61, 0x6c, 0xc4, 0xb1, 0x6b}},
		daysAbbreviated:        [][]uint8{{0x50, 0x61, 0x7a}, {0x50, 0x7a, 0x74}, {0x53, 0x61, 0x6c}, {0xc3, 0x87, 0x61, 0x72}, {0x50, 0x65, 0x72}, {0x43, 0x75, 0x6d}, {0x43, 0x6d, 0x74}},
		daysNarrow:             [][]uint8{{0x50}, {0x50}, {0x53}, {0xc3, 0x87}, {0x50}, {0x43}, {0x43}},
		daysShort:              [][]uint8{{0x50, 0x61}, {0x50, 0x74}, {0x53, 0x61}, {0xc3, 0x87, 0x61}, {0x50, 0x65}, {0x43, 0x75}, {0x43, 0x74}},
		daysWide:               [][]uint8{{0x50, 0x61, 0x7a, 0x61, 0x72}, {0x50, 0x61, 0x7a, 0x61, 0x72, 0x74, 0x65, 0x73, 0x69}, {0x53, 0x61, 0x6c, 0xc4, 0xb1}, {0xc3, 0x87, 0x61, 0x72, 0xc5, 0x9f, 0x61, 0x6d, 0x62, 0x61}, {0x50, 0x65, 0x72, 0xc5, 0x9f, 0x65, 0x6d, 0x62, 0x65}, {0x43, 0x75, 0x6d, 0x61}, {0x43, 0x75, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x73, 0x69}},
		periodsAbbreviated:     [][]uint8{{0xc3, 0x96, 0xc3, 0x96}, {0xc3, 0x96, 0x53}},
		periodsNarrow:          [][]uint8{{0xc3, 0xb6, 0xc3, 0xb6}, {0xc3, 0xb6, 0x73}},
		periodsWide:            [][]uint8{{0xc3, 0x96, 0xc3, 0x96}, {0xc3, 0x96, 0x53}},
		erasAbbreviated:        [][]uint8{{0x4d, 0xc3, 0x96}, {0x4d, 0x53}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x4d, 0x69, 0x6c, 0x61, 0x74, 0x74, 0x61, 0x6e, 0x20, 0xc3, 0x96, 0x6e, 0x63, 0x65}, {0x4d, 0x69, 0x6c, 0x61, 0x74, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x6f, 0x6e, 0x72, 0x61}},
		timezones:              map[string][]uint8{"AEDT": {0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CLST": {0xc5, 0x9e, 0x69, 0x6c, 0x69, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WARST": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x72, 0x6a, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "PDT": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x50, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "MESZ": {0x4f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x76, 0x72, 0x75, 0x70, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CDT": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "EAT": {0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "MEZ": {0x4f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x76, 0x72, 0x75, 0x70, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "JST": {0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "NZST": {0x59, 0x65, 0x6e, 0x69, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CLT": {0xc5, 0x9e, 0x69, 0x6c, 0x69, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "MST": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x44, 0x61, 0xc4, 0x9f, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WART": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x72, 0x6a, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "COST": {0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "AWDT": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WIB": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x45, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x79, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ACWST": {0xc4, 0xb0, 0xc3, 0xa7, 0x62, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "OEZ": {0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x41, 0x76, 0x72, 0x75, 0x70, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "EDT": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ACST": {0x4f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ACWDT": {0xc4, 0xb0, 0xc3, 0xa7, 0x62, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WEZ": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x76, 0x72, 0x75, 0x70, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WAT": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ACDT": {0x4f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "PST": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x50, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CST": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x79, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WIT": {0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x45, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x79, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "MDT": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x44, 0x61, 0xc4, 0x9f, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CAT": {0x4f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ChST": {0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0x73, 0xc4, 0xb1, 0x7a, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x73, 0xc4, 0xb1, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x4f, 0x72, 0x74, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ART": {0x41, 0x72, 0x6a, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "TMT": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "TMST": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "OESZ": {0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x41, 0x76, 0x72, 0x75, 0x70, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "IST": {0x48, 0x69, 0x6e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WAST": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "AEST": {0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ARST": {0x41, 0x72, 0x6a, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "COT": {0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "JDT": {0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x79, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "EST": {0x4b, 0x75, 0x7a, 0x65, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x44, 0x6f, 0xc4, 0x9f, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WESZ": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x76, 0x72, 0x75, 0x70, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "HNT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "MYT": {0x4d, 0x61, 0x6c, 0x65, 0x7a, 0x79, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "NZDT": {0x59, 0x65, 0x6e, 0x69, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "AWST": {0x42, 0x61, 0x74, 0xc4, 0xb1, 0x20, 0x41, 0x76, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "∅∅∅": {0x41, 0x63, 0x72, 0x65, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "SAST": {0x47, 0xc3, 0xbc, 0x6e, 0x65, 0x79, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "ECT": {0x45, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "WITA": {0x4f, 0x72, 0x74, 0x61, 0x20, 0x45, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x79, 0x61, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}, "HAT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x59, 0x61, 0x7a, 0x20, 0x53, 0x61, 0x61, 0x74, 0x69}},
	}
}

// Locale returns the current translators string locale
func (tr *tr_CY) Locale() string {
	return tr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'tr_CY'
func (tr *tr_CY) PluralsCardinal() []locales.PluralRule {
	return tr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'tr_CY'
func (tr *tr_CY) PluralsOrdinal() []locales.PluralRule {
	return tr.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'tr_CY'
func (tr *tr_CY) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'tr_CY'
func (tr *tr_CY) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'tr_CY'
func (tr *tr_CY) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := tr.CardinalPluralRule(num1, v1)
	end := tr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (tr *tr_CY) MonthAbbreviated(month time.Month) []byte {
	return tr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (tr *tr_CY) MonthsAbbreviated() [][]byte {
	return tr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (tr *tr_CY) MonthNarrow(month time.Month) []byte {
	return tr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (tr *tr_CY) MonthsNarrow() [][]byte {
	return tr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (tr *tr_CY) MonthWide(month time.Month) []byte {
	return tr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (tr *tr_CY) MonthsWide() [][]byte {
	return tr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (tr *tr_CY) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return tr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (tr *tr_CY) WeekdaysAbbreviated() [][]byte {
	return tr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (tr *tr_CY) WeekdayNarrow(weekday time.Weekday) []byte {
	return tr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (tr *tr_CY) WeekdaysNarrow() [][]byte {
	return tr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (tr *tr_CY) WeekdayShort(weekday time.Weekday) []byte {
	return tr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (tr *tr_CY) WeekdaysShort() [][]byte {
	return tr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (tr *tr_CY) WeekdayWide(weekday time.Weekday) []byte {
	return tr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (tr *tr_CY) WeekdaysWide() [][]byte {
	return tr.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'tr_CY' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(tr.decimal) + len(tr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'tr_CY' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (tr *tr_CY) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(tr.decimal) + len(tr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tr.minus[0])
	}

	b = append(b, tr.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tr.currencies[currency]
	l := len(s) + len(tr.decimal) + len(tr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, tr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'tr_CY'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tr.currencies[currency]
	l := len(s) + len(tr.decimal) + len(tr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(tr.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, tr.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, tr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, tr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.daysWide[t.Weekday()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, tr.periodsAbbreviated[0]...)
	} else {
		b = append(b, tr.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, tr.periodsAbbreviated[0]...)
	} else {
		b = append(b, tr.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, tr.periodsAbbreviated[0]...)
	} else {
		b = append(b, tr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'tr_CY'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tr *tr_CY) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, tr.periodsAbbreviated[0]...)
	} else {
		b = append(b, tr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := tr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
