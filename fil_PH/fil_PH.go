package fil_PH

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type fil_PH struct {
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

// New returns a new instance of translator for the 'fil_PH' locale
func New() locales.Translator {
	return &fil_PH{
		locale:                 "fil_PH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x45, 0x6e, 0x65}, {0x50, 0x65, 0x62}, {0x4d, 0x61, 0x72}, {0x41, 0x62, 0x72}, {0x4d, 0x61, 0x79}, {0x48, 0x75, 0x6e}, {0x48, 0x75, 0x6c}, {0x41, 0x67, 0x6f}, {0x53, 0x65, 0x74}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x62}, {0x44, 0x69, 0x73}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x45, 0x6e, 0x65}, {0x50, 0x65, 0x62}, {0x4d, 0x61, 0x72}, {0x41, 0x62, 0x72}, {0x4d, 0x61, 0x79}, {0x48, 0x75, 0x6e}, {0x48, 0x75, 0x6c}, {0x41, 0x67, 0x6f}, {0x53, 0x65, 0x74}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x62}, {0x44, 0x69, 0x73}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x45, 0x6e, 0x65, 0x72, 0x6f}, {0x50, 0x65, 0x62, 0x72, 0x65, 0x72, 0x6f}, {0x4d, 0x61, 0x72, 0x73, 0x6f}, {0x41, 0x62, 0x72, 0x69, 0x6c}, {0x4d, 0x61, 0x79, 0x6f}, {0x48, 0x75, 0x6e, 0x79, 0x6f}, {0x48, 0x75, 0x6c, 0x79, 0x6f}, {0x41, 0x67, 0x6f, 0x73, 0x74, 0x6f}, {0x53, 0x65, 0x74, 0x79, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x4f, 0x6b, 0x74, 0x75, 0x62, 0x72, 0x65}, {0x4e, 0x6f, 0x62, 0x79, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x44, 0x69, 0x73, 0x79, 0x65, 0x6d, 0x62, 0x72, 0x65}},
		daysAbbreviated:        [][]uint8{{0x4c, 0x69, 0x6e}, {0x4c, 0x75, 0x6e}, {0x4d, 0x61, 0x72}, {0x4d, 0x69, 0x79}, {0x48, 0x75, 0x77}, {0x42, 0x69, 0x79}, {0x53, 0x61, 0x62}},
		daysNarrow:             [][]uint8{{0x4c, 0x69, 0x6e}, {0x4c, 0x75, 0x6e}, {0x4d, 0x61, 0x72}, {0x4d, 0x69, 0x79}, {0x48, 0x75, 0x77}, {0x42, 0x69, 0x79}, {0x53, 0x61, 0x62}},
		daysShort:              [][]uint8{{0x4c, 0x69}, {0x4c, 0x75}, {0x4d, 0x61}, {0x4d, 0x69}, {0x48, 0x75}, {0x42, 0x69}, {0x53, 0x61}},
		daysWide:               [][]uint8{{0x4c, 0x69, 0x6e, 0x67, 0x67, 0x6f}, {0x4c, 0x75, 0x6e, 0x65, 0x73}, {0x4d, 0x61, 0x72, 0x74, 0x65, 0x73}, {0x4d, 0x69, 0x79, 0x65, 0x72, 0x6b, 0x75, 0x6c, 0x65, 0x73}, {0x48, 0x75, 0x77, 0x65, 0x62, 0x65, 0x73}, {0x42, 0x69, 0x79, 0x65, 0x72, 0x6e, 0x65, 0x73}, {0x53, 0x61, 0x62, 0x61, 0x64, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61, 0x6d}, {0x70, 0x6d}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		timezones:              map[string][]uint8{"AKST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "HKT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "ACWST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x20, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "TMT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "HADT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "OEZ": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x53, 0x69, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x67, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "JDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x4a, 0x61, 0x70, 0x61, 0x6e}, "BOT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "SRT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "WEZ": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "WART": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "ACDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "ACST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "WAST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61}, "EAT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x53, 0x69, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61}, "ART": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "ARST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "JST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4a, 0x61, 0x70, 0x61, 0x6e}, "EST": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73}, "COT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "MEZ": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "MESZ": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "CDT": {0x53, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x6e, 0x61, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "GFT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x46, 0x72, 0x65, 0x6e, 0x63, 0x68, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61}, "AWDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "UYST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "PST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x50, 0x61, 0x73, 0x69, 0x70, 0x69, 0x6b, 0x6f}, "PDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x50, 0x61, 0x73, 0x69, 0x70, 0x69, 0x6b, 0x6f}, "GYT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "ECT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "NZDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "MYT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x6e, 0x67, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "CHAST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "OESZ": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x67, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "SGT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65}, "BT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "AEST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "MST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x42, 0x75, 0x6e, 0x64, 0x6f, 0x6b}, "HKST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "AEDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "COST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "HNT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "UYT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x6e, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ADT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x6f}, "CLT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "ChST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "VET": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "WESZ": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "CAT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61}, "TMST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "CLST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "ACWDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x20, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "HAT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "AKDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "LHDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WIB": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x79, 0x61}, "IST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "AWST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x79, 0x61}, "LHST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "MDT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x42, 0x75, 0x6e, 0x64, 0x6f, 0x6b}, "WARST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "WIT": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x67, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x79, 0x61}, "∅∅∅": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x61, 0x67, 0x2d, 0x69, 0x6e, 0x69, 0x74, 0x20, 0x6e, 0x67, 0x20, 0x41, 0x7a, 0x6f, 0x72, 0x65, 0x73}, "WITA": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x6e, 0x67, 0x20, 0x47, 0x69, 0x74, 0x6e, 0x61, 0x6e, 0x67, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x79, 0x61}, "WAT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4b, 0x61, 0x6e, 0x6c, 0x75, 0x72, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61}, "HAST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "CHADT": {0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "SAST": {0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x6f, 0x67, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61}, "AST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x6f}, "EDT": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CST": {0x53, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73}, "NZST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x72, 0x61, 0x73, 0x20, 0x73, 0x61, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}},
	}
}

// Locale returns the current translators string locale
func (fil *fil_PH) Locale() string {
	return fil.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fil_PH'
func (fil *fil_PH) PluralsCardinal() []locales.PluralRule {
	return fil.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fil_PH'
func (fil *fil_PH) PluralsOrdinal() []locales.PluralRule {
	return fil.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fil_PH'
func (fil *fil_PH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	fMod10 := f % 10

	if (v == 0 && (i == 1 || i == 2 || i == 3)) || (v == 0 && (iMod10 != 4 && iMod10 != 6 && iMod10 != 9)) || (v != 0 && (fMod10 != 4 && fMod10 != 6 && fMod10 != 9)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fil_PH'
func (fil *fil_PH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fil_PH'
func (fil *fil_PH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := fil.CardinalPluralRule(num1, v1)
	end := fil.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fil *fil_PH) MonthAbbreviated(month time.Month) []byte {
	return fil.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fil *fil_PH) MonthsAbbreviated() [][]byte {
	return fil.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fil *fil_PH) MonthNarrow(month time.Month) []byte {
	return fil.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fil *fil_PH) MonthsNarrow() [][]byte {
	return fil.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fil *fil_PH) MonthWide(month time.Month) []byte {
	return fil.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fil *fil_PH) MonthsWide() [][]byte {
	return fil.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fil *fil_PH) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return fil.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fil *fil_PH) WeekdaysAbbreviated() [][]byte {
	return fil.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fil *fil_PH) WeekdayNarrow(weekday time.Weekday) []byte {
	return fil.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fil *fil_PH) WeekdaysNarrow() [][]byte {
	return fil.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fil *fil_PH) WeekdayShort(weekday time.Weekday) []byte {
	return fil.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fil *fil_PH) WeekdaysShort() [][]byte {
	return fil.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fil *fil_PH) WeekdayWide(weekday time.Weekday) []byte {
	return fil.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fil *fil_PH) WeekdaysWide() [][]byte {
	return fil.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fil_PH' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fil.decimal) + len(fil.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fil.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fil.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fil_PH' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fil *fil_PH) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fil.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fil.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fil.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fil.currencies[currency]
	l := len(s) + len(fil.decimal) + len(fil.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fil.group[0])
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
		b = append(b, fil.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fil.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fil_PH'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fil.currencies[currency]
	l := len(s) + len(fil.decimal) + len(fil.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fil.group[0])
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

		for j := len(fil.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, fil.currencyNegativePrefix[j])
		}

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
			b = append(b, fil.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fil.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, fil.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, fil.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, fil.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, fil.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fil.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fil.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'fil_PH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fil *fil_PH) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fil.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fil.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
