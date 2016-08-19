package et_EE

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type et_EE struct {
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

// New returns a new instance of translator for the 'et_EE' locale
func New() locales.Translator {
	return &et_EE{
		locale:                 "et_EE",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0xe2, 0x88, 0x92},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x61, 0x6e}, {0x76, 0x65, 0x65, 0x62, 0x72}, {0x6d, 0xc3, 0xa4, 0x72, 0x74, 0x73}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0x69}, {0x6a, 0x75, 0x75, 0x6e, 0x69}, {0x6a, 0x75, 0x75, 0x6c, 0x69}, {0x61, 0x75, 0x67}, {0x73, 0x65, 0x70, 0x74}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x74, 0x73}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x56}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x61, 0x6e, 0x75, 0x61, 0x72}, {0x76, 0x65, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72}, {0x6d, 0xc3, 0xa4, 0x72, 0x74, 0x73}, {0x61, 0x70, 0x72, 0x69, 0x6c, 0x6c}, {0x6d, 0x61, 0x69}, {0x6a, 0x75, 0x75, 0x6e, 0x69}, {0x6a, 0x75, 0x75, 0x6c, 0x69}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x6f, 0x6b, 0x74, 0x6f, 0x6f, 0x62, 0x65, 0x72}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x64, 0x65, 0x74, 0x73, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x50}, {0x45}, {0x54}, {0x4b}, {0x4e}, {0x52}, {0x4c}},
		daysNarrow:             [][]uint8{{0x50}, {0x45}, {0x54}, {0x4b}, {0x4e}, {0x52}, {0x4c}},
		daysShort:              [][]uint8{{0x50}, {0x45}, {0x54}, {0x4b}, {0x4e}, {0x52}, {0x4c}},
		daysWide:               [][]uint8{{0x70, 0xc3, 0xbc, 0x68, 0x61, 0x70, 0xc3, 0xa4, 0x65, 0x76}, {0x65, 0x73, 0x6d, 0x61, 0x73, 0x70, 0xc3, 0xa4, 0x65, 0x76}, {0x74, 0x65, 0x69, 0x73, 0x69, 0x70, 0xc3, 0xa4, 0x65, 0x76}, {0x6b, 0x6f, 0x6c, 0x6d, 0x61, 0x70, 0xc3, 0xa4, 0x65, 0x76}, {0x6e, 0x65, 0x6c, 0x6a, 0x61, 0x70, 0xc3, 0xa4, 0x65, 0x76}, {0x72, 0x65, 0x65, 0x64, 0x65}, {0x6c, 0x61, 0x75, 0x70, 0xc3, 0xa4, 0x65, 0x76}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x65, 0x4b, 0x72}, {0x70, 0x4b, 0x72}},
		erasNarrow:             [][]uint8{{0x65, 0x4b, 0x72}, {0x70, 0x4b, 0x72}},
		erasWide:               [][]uint8{{0x65, 0x6e, 0x6e, 0x65, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x75, 0x73, 0x74}, {0x70, 0xc3, 0xa4, 0x72, 0x61, 0x73, 0x74, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x75, 0x73, 0x74}},
		timezones:              map[string][]uint8{"MST": {0x4d, 0x53, 0x54}, "MEZ": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "OESZ": {0x49, 0x64, 0x61, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "JDT": {0x4a, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "WAT": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x41, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "AWST": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ACWDT": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "CLT": {0x54, 0xc5, 0xa1, 0x69, 0x69, 0x6c, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "WESZ": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "SAST": {0x4c, 0xc3, 0xb5, 0x75, 0x6e, 0x61, 0x2d, 0x41, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "MESZ": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "EDT": {0x49, 0x64, 0x61, 0x72, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x75, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x75, 0x64, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ECT": {0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x69, 0x20, 0x61, 0x65, 0x67}, "COST": {0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "ACDT": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "∅∅∅": {0x42, 0x72, 0x61, 0x73, 0x69, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "COT": {0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x61, 0x65, 0x67}, "EST": {0x49, 0x64, 0x61, 0x72, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x75, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x69, 0x73, 0x69, 0x61, 0x20, 0xe2, 0x80, 0x8b, 0xe2, 0x80, 0x8b, 0x61, 0x65, 0x67}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x69, 0x20, 0x61, 0x65, 0x67}, "PDT": {0x56, 0x61, 0x69, 0x6b, 0x73, 0x65, 0x20, 0x6f, 0x6f, 0x6b, 0x65, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "OEZ": {0x49, 0x64, 0x61, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "AEST": {0x49, 0x64, 0x61, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "AEDT": {0x49, 0x64, 0x61, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "NZDT": {0x55, 0x75, 0x73, 0x2d, 0x4d, 0x65, 0x72, 0x65, 0x6d, 0x61, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "WIB": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x65, 0x73, 0x69, 0x61, 0x20, 0x61, 0x65, 0x67}, "PST": {0x56, 0x61, 0x69, 0x6b, 0x73, 0x65, 0x20, 0x6f, 0x6f, 0x6b, 0x65, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x75, 0x64, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x69, 0x20, 0x61, 0x65, 0x67}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "WITA": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x65, 0x73, 0x69, 0x61, 0x20, 0x61, 0x65, 0x67}, "CLST": {0x54, 0xc5, 0xa1, 0x69, 0x69, 0x6c, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "WEZ": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "HAT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "WIT": {0x49, 0x64, 0x61, 0x2d, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x65, 0x73, 0x69, 0x61, 0x20, 0x61, 0x65, 0x67}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x61, 0x65, 0x67}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x61, 0x20, 0x61, 0x65, 0x67}, "HNT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "AWDT": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "WART": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "JST": {0x4a, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "CAT": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x41, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x61, 0x65, 0x67}, "CST": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x41, 0x6d, 0x65, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x69, 0x76, 0x69, 0x61, 0x20, 0x61, 0x65, 0x67}, "WAST": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x41, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "NZST": {0x55, 0x75, 0x73, 0x2d, 0x4d, 0x65, 0x72, 0x65, 0x6d, 0x61, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "MDT": {0x4d, 0x44, 0x54}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x20, 0x61, 0x65, 0x67}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "TMST": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "ACWST": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ACST": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ChST": {0x54, 0xc5, 0xa1, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "WARST": {0x4c, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x65, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "TMT": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x61, 0x65, 0x67}, "EAT": {0x49, 0x64, 0x61, 0x2d, 0x41, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x61, 0x65, 0x67}, "CDT": {0x4b, 0x65, 0x73, 0x6b, 0x2d, 0x41, 0x6d, 0x65, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x75, 0x76, 0x65, 0x61, 0x65, 0x67}, "GFT": {0x50, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x75, 0x73, 0x65, 0x20, 0x47, 0x75, 0x61, 0x6a, 0x61, 0x61, 0x6e, 0x61, 0x20, 0x61, 0x65, 0x67}},
	}
}

// Locale returns the current translators string locale
func (et *et_EE) Locale() string {
	return et.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'et_EE'
func (et *et_EE) PluralsCardinal() []locales.PluralRule {
	return et.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'et_EE'
func (et *et_EE) PluralsOrdinal() []locales.PluralRule {
	return et.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'et_EE'
func (et *et_EE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'et_EE'
func (et *et_EE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'et_EE'
func (et *et_EE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (et *et_EE) MonthAbbreviated(month time.Month) []byte {
	return et.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (et *et_EE) MonthsAbbreviated() [][]byte {
	return et.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (et *et_EE) MonthNarrow(month time.Month) []byte {
	return et.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (et *et_EE) MonthsNarrow() [][]byte {
	return et.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (et *et_EE) MonthWide(month time.Month) []byte {
	return et.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (et *et_EE) MonthsWide() [][]byte {
	return et.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (et *et_EE) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return et.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (et *et_EE) WeekdaysAbbreviated() [][]byte {
	return et.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (et *et_EE) WeekdayNarrow(weekday time.Weekday) []byte {
	return et.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (et *et_EE) WeekdaysNarrow() [][]byte {
	return et.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (et *et_EE) WeekdayShort(weekday time.Weekday) []byte {
	return et.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (et *et_EE) WeekdaysShort() [][]byte {
	return et.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (et *et_EE) WeekdayWide(weekday time.Weekday) []byte {
	return et.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (et *et_EE) WeekdaysWide() [][]byte {
	return et.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'et_EE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(et.decimal) + len(et.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(et.group) - 1; j >= 0; j-- {
					b = append(b, et.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(et.minus) - 1; j >= 0; j-- {
			b = append(b, et.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'et_EE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (et *et_EE) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(et.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(et.minus) - 1; j >= 0; j-- {
			b = append(b, et.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, et.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := et.currencies[currency]
	l := len(s) + len(et.decimal) + len(et.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(et.group) - 1; j >= 0; j-- {
					b = append(b, et.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(et.minus) - 1; j >= 0; j-- {
			b = append(b, et.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, et.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, et.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'et_EE'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := et.currencies[currency]
	l := len(s) + len(et.decimal) + len(et.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(et.group) - 1; j >= 0; j-- {
					b = append(b, et.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(et.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, et.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, et.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, et.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, et.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, et.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, et.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, et.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, et.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

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

// FmtTimeLong returns the long time representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

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

// FmtTimeFull returns the full time representation of 't' for 'et_EE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (et *et_EE) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

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

	if btz, ok := et.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
