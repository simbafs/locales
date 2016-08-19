package hsb_DE

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type hsb_DE struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositiveSuffix []byte
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

// New returns a new instance of translator for the 'hsb_DE' locale
func New() locales.Translator {
	return &hsb_DE{
		locale:                 "hsb_DE",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x2e}, {0x66, 0x65, 0x62, 0x2e}, {0x6d, 0xc4, 0x9b, 0x72, 0x2e}, {0x61, 0x70, 0x72, 0x2e}, {0x6d, 0x65, 0x6a, 0x2e}, {0x6a, 0x75, 0x6e, 0x2e}, {0x6a, 0x75, 0x6c, 0x2e}, {0x61, 0x77, 0x67, 0x2e}, {0x73, 0x65, 0x70, 0x2e}, {0x6f, 0x6b, 0x74, 0x2e}, {0x6e, 0x6f, 0x77, 0x2e}, {0x64, 0x65, 0x63, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x6a}, {0x66}, {0x6d}, {0x61}, {0x6d}, {0x6a}, {0x6a}, {0x61}, {0x73}, {0x6f}, {0x6e}, {0x64}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x61}, {0x66, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72, 0x61}, {0x6d, 0xc4, 0x9b, 0x72, 0x63, 0x61}, {0x61, 0x70, 0x72, 0x79, 0x6c, 0x61}, {0x6d, 0x65, 0x6a, 0x65}, {0x6a, 0x75, 0x6e, 0x69, 0x6a, 0x61}, {0x6a, 0x75, 0x6c, 0x69, 0x6a, 0x61}, {0x61, 0x77, 0x67, 0x75, 0x73, 0x74, 0x61}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x61}, {0x6f, 0x6b, 0x74, 0x6f, 0x62, 0x72, 0x61}, {0x6e, 0x6f, 0x77, 0x65, 0x6d, 0x62, 0x72, 0x61}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x61}},
		daysAbbreviated:        [][]uint8{{0x6e, 0x6a, 0x65}, {0x70, 0xc3, 0xb3, 0x6e}, {0x77, 0x75, 0x74}, {0x73, 0x72, 0x6a}, {0xc5, 0xa1, 0x74, 0x77}, {0x70, 0x6a, 0x61}, {0x73, 0x6f, 0x62}},
		daysNarrow:             [][]uint8{{0x6e}, {0x70}, {0x77}, {0x73}, {0xc5, 0xa1}, {0x70}, {0x73}},
		daysShort:              [][]uint8{{0x6e, 0x6a}, {0x70, 0xc3, 0xb3}, {0x77, 0x75}, {0x73, 0x72}, {0xc5, 0xa1, 0x74}, {0x70, 0x6a}, {0x73, 0x6f}},
		daysWide:               [][]uint8{{0x6e, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x65, 0x6c, 0x61}, {0x70, 0xc3, 0xb3, 0x6e, 0x64, 0xc5, 0xba, 0x65, 0x6c, 0x61}, {0x77, 0x75, 0x74, 0x6f, 0x72, 0x61}, {0x73, 0x72, 0x6a, 0x65, 0x64, 0x61}, {0xc5, 0xa1, 0x74, 0x77, 0xc3, 0xb3, 0x72, 0x74, 0x6b}, {0x70, 0x6a, 0x61, 0x74, 0x6b}, {0x73, 0x6f, 0x62, 0x6f, 0x74, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x64, 0x6f, 0x70, 0x6f, 0xc5, 0x82, 0x64, 0x6e, 0x6a, 0x61}, {0x70, 0x6f, 0x70, 0x6f, 0xc5, 0x82, 0x64, 0x6e, 0x6a, 0x75}},
		periodsNarrow:          [][]uint8{{0x64, 0x6f, 0x70, 0x2e}, {0x70, 0x6f, 0x70, 0x2e}},
		periodsWide:            [][]uint8{{0x64, 0x6f, 0x70, 0x6f, 0xc5, 0x82, 0x64, 0x6e, 0x6a, 0x61}, {0x70, 0x6f, 0x70, 0x6f, 0xc5, 0x82, 0x64, 0x6e, 0x6a, 0x75}},
		erasAbbreviated:        [][]uint8{{0x70, 0xc5, 0x99, 0x2e, 0x43, 0x68, 0x72, 0x2e, 0x6e, 0x2e}, {0x70, 0x6f, 0x20, 0x43, 0x68, 0x72, 0x2e, 0x6e, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x70, 0xc5, 0x99, 0x65, 0x64, 0x20, 0x43, 0x68, 0x72, 0x79, 0x73, 0x74, 0x6f, 0x77, 0x79, 0x6d, 0x20, 0x6e, 0x61, 0x72, 0x6f, 0x64, 0xc5, 0xba, 0x65, 0x6e, 0x6a, 0x6f, 0x6d}, {0x70, 0x6f, 0x20, 0x43, 0x68, 0x72, 0x79, 0x73, 0x74, 0x6f, 0x77, 0x79, 0x6d, 0x20, 0x6e, 0x61, 0x72, 0x6f, 0x64, 0xc5, 0xba, 0x65, 0x6e, 0x6a, 0x75}},
		timezones:              map[string][]uint8{"ACST": {0x73, 0x72, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ECT": {0x65, 0x6b, 0x77, 0x61, 0x64, 0x6f, 0x72, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MESZ": {0x73, 0x72, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ACWST": {0x73, 0x72, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x79, 0x20, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "IST": {0x69, 0x6e, 0x64, 0x69, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MYT": {0x6d, 0x61, 0x6c, 0x61, 0x6a, 0x7a, 0x69, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WITA": {0x73, 0x72, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "COT": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "VET": {0x76, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "JST": {0x6a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HNT": {0x6e, 0x6f, 0x77, 0x6f, 0x66, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MDT": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x68, 0xc3, 0xb3, 0x72, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AEDT": {0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WESZ": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CAT": {0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HADT": {0x68, 0x61, 0x77, 0x61, 0x69, 0x69, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "JDT": {0x6a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "EDT": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6e, 0x79, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "EST": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6e, 0x79, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HAT": {0x6e, 0x6f, 0x77, 0x6f, 0x66, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WIB": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CLT": {0x63, 0x68, 0x69, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "SRT": {0x73, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "COST": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "UYST": {0x75, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AKST": {0x61, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AWST": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "SAST": {0x6a, 0x75, 0xc5, 0xbe, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "OEZ": {0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ACWDT": {0x73, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x79, 0x20, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "LHDT": {0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73, 0x20, 0x6b, 0x75, 0x70, 0x79, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x6f, 0x77, 0x65}, "HAST": {0x68, 0x61, 0x77, 0x61, 0x69, 0x69, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ChST": {0x63, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ARST": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "BT": {0x62, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CST": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x79, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "∅∅∅": {0x61, 0x63, 0x6f, 0x72, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CHADT": {0x63, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ADT": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AKDT": {0x61, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CLST": {0x63, 0x68, 0x69, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "GFT": {0x66, 0x72, 0x61, 0x6e, 0x63, 0x6f, 0x73, 0x6b, 0x6f, 0x67, 0x75, 0x79, 0x61, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WEZ": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CDT": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x79, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "NZDT": {0x6e, 0x6f, 0x77, 0x6f, 0x73, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "LHST": {0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73, 0x20, 0x6b, 0x75, 0x70, 0x79, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x6f, 0x77, 0x65}, "WART": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "GYT": {0x67, 0x75, 0x79, 0x61, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "PST": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x70, 0x61, 0x63, 0x69, 0x66, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AEST": {0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WIT": {0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x6b, 0x69}, "CHAST": {0x63, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "TMST": {0x74, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "OESZ": {0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WAST": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "PDT": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x70, 0x61, 0x63, 0x69, 0x66, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ART": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MEZ": {0x73, 0x72, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WARST": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "EAT": {0x77, 0x75, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WAT": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "NZST": {0x6e, 0x6f, 0x77, 0x6f, 0x73, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "UYT": {0x75, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AWDT": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MST": {0x73, 0x65, 0x77, 0x6a, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x68, 0xc3, 0xb3, 0x72, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AST": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "BOT": {0x62, 0x6f, 0x6c, 0x69, 0x77, 0x69, 0x73, 0x6b, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "TMT": {0x74, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x73, 0x6b, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ACDT": {0x73, 0x72, 0x6a, 0x65, 0x64, 0xc5, 0xba, 0x6f, 0x61, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x69, 0x20, 0x6c, 0xc4, 0x9b, 0x74, 0x6e, 0x69, 0x20, 0xc4, 0x8d, 0x61, 0x73}},
	}
}

// Locale returns the current translators string locale
func (hsb *hsb_DE) Locale() string {
	return hsb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hsb_DE'
func (hsb *hsb_DE) PluralsCardinal() []locales.PluralRule {
	return hsb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hsb_DE'
func (hsb *hsb_DE) PluralsOrdinal() []locales.PluralRule {
	return hsb.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hsb_DE'
func (hsb *hsb_DE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod100 := i % 100
	fMod100 := f % 100

	if (v == 0 && iMod100 == 1) || (fMod100 == 1) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod100 == 2) || (fMod100 == 2) {
		return locales.PluralRuleTwo
	} else if (v == 0 && iMod100 >= 3 && iMod100 <= 4) || (fMod100 >= 3 && fMod100 <= 4) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hsb_DE'
func (hsb *hsb_DE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hsb_DE'
func (hsb *hsb_DE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (hsb *hsb_DE) MonthAbbreviated(month time.Month) []byte {
	return hsb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (hsb *hsb_DE) MonthsAbbreviated() [][]byte {
	return hsb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (hsb *hsb_DE) MonthNarrow(month time.Month) []byte {
	return hsb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (hsb *hsb_DE) MonthsNarrow() [][]byte {
	return hsb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (hsb *hsb_DE) MonthWide(month time.Month) []byte {
	return hsb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (hsb *hsb_DE) MonthsWide() [][]byte {
	return hsb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return hsb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (hsb *hsb_DE) WeekdaysAbbreviated() [][]byte {
	return hsb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayNarrow(weekday time.Weekday) []byte {
	return hsb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (hsb *hsb_DE) WeekdaysNarrow() [][]byte {
	return hsb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayShort(weekday time.Weekday) []byte {
	return hsb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (hsb *hsb_DE) WeekdaysShort() [][]byte {
	return hsb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayWide(weekday time.Weekday) []byte {
	return hsb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (hsb *hsb_DE) WeekdaysWide() [][]byte {
	return hsb.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hsb_DE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(hsb.decimal) + len(hsb.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hsb_DE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hsb *hsb_DE) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(hsb.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, hsb.percentSuffix...)

	b = append(b, hsb.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hsb.currencies[currency]
	l := len(s) + len(hsb.decimal) + len(hsb.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hsb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, hsb.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hsb_DE'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hsb.currencies[currency]
	l := len(s) + len(hsb.decimal) + len(hsb.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, hsb.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hsb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, hsb.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, hsb.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hsb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, hsb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hsb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x68, 0x6f, 0x64, 0xc5, 0xba, 0x27, 0x2e}...)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'hsb_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hsb *hsb_DE) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hsb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
