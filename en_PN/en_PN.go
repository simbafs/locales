package en_PN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type en_PN struct {
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

// New returns a new instance of translator for the 'en_PN' locale
func New() locales.Translator {
	return &en_PN{
		locale:                 "en_PN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x46, 0x65, 0x62}, {0x4d, 0x61, 0x72}, {0x41, 0x70, 0x72}, {0x4d, 0x61, 0x79}, {0x4a, 0x75, 0x6e}, {0x4a, 0x75, 0x6c}, {0x41, 0x75, 0x67}, {0x53, 0x65, 0x70}, {0x4f, 0x63, 0x74}, {0x4e, 0x6f, 0x76}, {0x44, 0x65, 0x63}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x79}, {0x46, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72, 0x79}, {0x4d, 0x61, 0x72, 0x63, 0x68}, {0x41, 0x70, 0x72, 0x69, 0x6c}, {0x4d, 0x61, 0x79}, {0x4a, 0x75, 0x6e, 0x65}, {0x4a, 0x75, 0x6c, 0x79}, {0x41, 0x75, 0x67, 0x75, 0x73, 0x74}, {0x53, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x4f, 0x63, 0x74, 0x6f, 0x62, 0x65, 0x72}, {0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x44, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x53, 0x75, 0x6e}, {0x4d, 0x6f, 0x6e}, {0x54, 0x75, 0x65}, {0x57, 0x65, 0x64}, {0x54, 0x68, 0x75}, {0x46, 0x72, 0x69}, {0x53, 0x61, 0x74}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x54}, {0x57}, {0x54}, {0x46}, {0x53}},
		daysShort:              [][]uint8{{0x53, 0x75}, {0x4d, 0x6f}, {0x54, 0x75}, {0x57, 0x65}, {0x54, 0x68}, {0x46, 0x72}, {0x53, 0x61}},
		daysWide:               [][]uint8{{0x53, 0x75, 0x6e, 0x64, 0x61, 0x79}, {0x4d, 0x6f, 0x6e, 0x64, 0x61, 0x79}, {0x54, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79}, {0x57, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79}, {0x54, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79}, {0x46, 0x72, 0x69, 0x64, 0x61, 0x79}, {0x53, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x6d}, {0x70, 0x6d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x61, 0x6d}, {0x70, 0x6d}},
		erasAbbreviated:        [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		erasNarrow:             [][]uint8{{0x42}, {0x41}},
		erasWide:               [][]uint8{{0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x43, 0x68, 0x72, 0x69, 0x73, 0x74}, {0x41, 0x6e, 0x6e, 0x6f, 0x20, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x69}},
		timezones:              map[string][]uint8{"EST": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "MDT": {0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "NZDT": {0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WAST": {0x57, 0x65, 0x73, 0x74, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "EAT": {0x45, 0x61, 0x73, 0x74, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AEST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "COST": {0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WEZ": {0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ChST": {0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CLST": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "SAST": {0x53, 0x6f, 0x75, 0x74, 0x68, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WITA": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "PST": {0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ECT": {0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AEDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "OEZ": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "NZST": {0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CLT": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x6e, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WART": {0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WAT": {0x57, 0x65, 0x73, 0x74, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "MST": {0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "EDT": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CAT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ACST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WIB": {0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65}, "PDT": {0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "MESZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "JST": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CST": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "GFT": {0x46, 0x72, 0x65, 0x6e, 0x63, 0x68, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ACDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "OESZ": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "HNT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "JDT": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WIT": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WARST": {0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ACWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "MEZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "WESZ": {0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "ACWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "∅∅∅": {0x50, 0x65, 0x72, 0x75, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x20, 0x54, 0x69, 0x6d, 0x65}, "HAT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x65}, "COT": {0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}},
	}
}

// Locale returns the current translators string locale
func (en *en_PN) Locale() string {
	return en.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'en_PN'
func (en *en_PN) PluralsCardinal() []locales.PluralRule {
	return en.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'en_PN'
func (en *en_PN) PluralsOrdinal() []locales.PluralRule {
	return en.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'en_PN'
func (en *en_PN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'en_PN'
func (en *en_PN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 1 && nMod100 != 11 {
		return locales.PluralRuleOne
	} else if nMod10 == 2 && nMod100 != 12 {
		return locales.PluralRuleTwo
	} else if nMod10 == 3 && nMod100 != 13 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'en_PN'
func (en *en_PN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (en *en_PN) MonthAbbreviated(month time.Month) []byte {
	return en.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (en *en_PN) MonthsAbbreviated() [][]byte {
	return en.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (en *en_PN) MonthNarrow(month time.Month) []byte {
	return en.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (en *en_PN) MonthsNarrow() [][]byte {
	return en.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (en *en_PN) MonthWide(month time.Month) []byte {
	return en.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (en *en_PN) MonthsWide() [][]byte {
	return en.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (en *en_PN) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return en.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (en *en_PN) WeekdaysAbbreviated() [][]byte {
	return en.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (en *en_PN) WeekdayNarrow(weekday time.Weekday) []byte {
	return en.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (en *en_PN) WeekdaysNarrow() [][]byte {
	return en.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (en *en_PN) WeekdayShort(weekday time.Weekday) []byte {
	return en.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (en *en_PN) WeekdaysShort() [][]byte {
	return en.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (en *en_PN) WeekdayWide(weekday time.Weekday) []byte {
	return en.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (en *en_PN) WeekdaysWide() [][]byte {
	return en.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'en_PN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(en.decimal) + len(en.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, en.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, en.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'en_PN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (en *en_PN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(en.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, en.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, en.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := en.currencies[currency]
	l := len(s) + len(en.decimal) + len(en.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, en.group[0])
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
		b = append(b, en.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, en.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'en_PN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := en.currencies[currency]
	l := len(s) + len(en.decimal) + len(en.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, en.group[0])
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

		for j := len(en.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, en.currencyNegativePrefix[j])
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
			b = append(b, en.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, en.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, en.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, en.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, en.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, en.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, en.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, en.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'en_PN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (en *en_PN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, en.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := en.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
