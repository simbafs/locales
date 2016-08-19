package gsw_CH

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type gsw_CH struct {
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

// New returns a new instance of translator for the 'gsw_CH' locale
func New() locales.Translator {
	return &gsw_CH{
		locale:                 "gsw_CH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0xe2, 0x80, 0x99},
		minus:                  []byte{0xe2, 0x88, 0x92},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x46, 0x65, 0x62}, {0x4d, 0xc3, 0xa4, 0x72}, {0x41, 0x70, 0x72}, {0x4d, 0x61, 0x69}, {0x4a, 0x75, 0x6e}, {0x4a, 0x75, 0x6c}, {0x41, 0x75, 0x67}, {0x53, 0x65, 0x70}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x76}, {0x44, 0x65, 0x7a}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x75, 0x61, 0x72}, {0x46, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72}, {0x4d, 0xc3, 0xa4, 0x72, 0x7a}, {0x41, 0x70, 0x72, 0x69, 0x6c}, {0x4d, 0x61, 0x69}, {0x4a, 0x75, 0x6e, 0x69}, {0x4a, 0x75, 0x6c, 0x69}, {0x41, 0x75, 0x67, 0x75, 0x73, 0x63, 0x68, 0x74}, {0x53, 0x65, 0x70, 0x74, 0xc3, 0xa4, 0x6d, 0x62, 0x65, 0x72}, {0x4f, 0x6b, 0x74, 0x6f, 0x6f, 0x62, 0x65, 0x72}, {0x4e, 0x6f, 0x76, 0xc3, 0xa4, 0x6d, 0x62, 0x65, 0x72}, {0x44, 0x65, 0x7a, 0xc3, 0xa4, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x53, 0x75, 0x2e}, {0x4d, 0xc3, 0xa4, 0x2e}, {0x5a, 0x69, 0x2e}, {0x4d, 0x69, 0x2e}, {0x44, 0x75, 0x2e}, {0x46, 0x72, 0x2e}, {0x53, 0x61, 0x2e}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x44}, {0x4d}, {0x44}, {0x46}, {0x53}},
		daysWide:               [][]uint8{{0x53, 0x75, 0x6e, 0x6e, 0x74, 0x69, 0x67}, {0x4d, 0xc3, 0xa4, 0xc3, 0xa4, 0x6e, 0x74, 0x69, 0x67}, {0x5a, 0x69, 0x69, 0x73, 0x63, 0x68, 0x74, 0x69, 0x67}, {0x4d, 0x69, 0x74, 0x74, 0x77, 0x75, 0x63, 0x68}, {0x44, 0x75, 0x6e, 0x73, 0x63, 0x68, 0x74, 0x69, 0x67}, {0x46, 0x72, 0x69, 0x69, 0x74, 0x69, 0x67}, {0x53, 0x61, 0x6d, 0x73, 0x63, 0x68, 0x74, 0x69, 0x67}},
		periodsAbbreviated:     [][]uint8{{0x76, 0x2e, 0x6d, 0x2e}, {0x6e, 0x2e, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x76, 0x6f, 0x72, 0x6d, 0x2e}, {0x6e, 0x61, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x76, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}, {0x6e, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x76, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}, {0x6e, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}},
		erasWide:               [][]uint8{{0x76, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}, {0x6e, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}},
		timezones:              map[string][]uint8{"CLST": {0x43, 0x4c, 0x53, 0x54}, "OESZ": {0x4f, 0x73, 0x63, 0x68, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x69, 0x69, 0x74}, "BT": {0x42, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "SAST": {0x53, 0xc3, 0xbc, 0xc3, 0xbc, 0x64, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x7a, 0x69, 0x69, 0x74}, "GYT": {0x47, 0x59, 0x54}, "WEZ": {0x57, 0x65, 0x73, 0x63, 0x68, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x63, 0x68, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x7a, 0x69, 0x69, 0x74}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "HAT": {0x48, 0x41, 0x54}, "WIT": {0x57, 0x49, 0x54}, "ART": {0x41, 0x52, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "CDT": {0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x2d, 0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x69, 0x69, 0x74}, "MDT": {0x4d, 0x44, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WESZ": {0x57, 0x65, 0x73, 0x63, 0x68, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x69, 0x69, 0x74}, "CST": {0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x2d, 0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x61, 0x6c, 0x20, 0x53, 0x63, 0x68, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x7a, 0x69, 0x69, 0x74}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x69, 0x69, 0x74}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "EAT": {0x4f, 0x73, 0x63, 0x68, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x5a, 0x69, 0x69, 0x74}, "JDT": {0x4a, 0x44, 0x54}, "CAT": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x5a, 0x69, 0x69, 0x74}, "PST": {0x50, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "WAST": {0x57, 0x65, 0x73, 0x63, 0x68, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x69, 0x69, 0x74}, "HNT": {0x48, 0x4e, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "MESZ": {0x4d, 0x69, 0x74, 0x74, 0x65, 0x6c, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x69, 0x69, 0x74}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "SGT": {0x53, 0x47, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "UYT": {0x55, 0x59, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "SRT": {0x53, 0x52, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "MEZ": {0x4d, 0x69, 0x74, 0x74, 0x65, 0x6c, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x63, 0x68, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x7a, 0x69, 0x69, 0x74}, "IST": {0x49, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "MST": {0x4d, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "WAT": {0x57, 0x65, 0x73, 0x63, 0x68, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x63, 0x68, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x7a, 0x69, 0x69, 0x74}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x53, 0x63, 0x68, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x7a, 0x69, 0x69, 0x74}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "OEZ": {0x4f, 0x73, 0x63, 0x68, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x69, 0x20, 0x53, 0x63, 0x68, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x7a, 0x69, 0x69, 0x74}, "PDT": {0x50, 0x44, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "AST": {0x41, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "ADT": {0x41, 0x44, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (gsw *gsw_CH) Locale() string {
	return gsw.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'gsw_CH'
func (gsw *gsw_CH) PluralsCardinal() []locales.PluralRule {
	return gsw.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'gsw_CH'
func (gsw *gsw_CH) PluralsOrdinal() []locales.PluralRule {
	return gsw.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'gsw_CH'
func (gsw *gsw_CH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'gsw_CH'
func (gsw *gsw_CH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'gsw_CH'
func (gsw *gsw_CH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (gsw *gsw_CH) MonthAbbreviated(month time.Month) []byte {
	return gsw.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (gsw *gsw_CH) MonthsAbbreviated() [][]byte {
	return gsw.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (gsw *gsw_CH) MonthNarrow(month time.Month) []byte {
	return gsw.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (gsw *gsw_CH) MonthsNarrow() [][]byte {
	return gsw.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (gsw *gsw_CH) MonthWide(month time.Month) []byte {
	return gsw.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (gsw *gsw_CH) MonthsWide() [][]byte {
	return gsw.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (gsw *gsw_CH) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return gsw.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (gsw *gsw_CH) WeekdaysAbbreviated() [][]byte {
	return gsw.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (gsw *gsw_CH) WeekdayNarrow(weekday time.Weekday) []byte {
	return gsw.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (gsw *gsw_CH) WeekdaysNarrow() [][]byte {
	return gsw.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (gsw *gsw_CH) WeekdayShort(weekday time.Weekday) []byte {
	return gsw.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (gsw *gsw_CH) WeekdaysShort() [][]byte {
	return gsw.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (gsw *gsw_CH) WeekdayWide(weekday time.Weekday) []byte {
	return gsw.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (gsw *gsw_CH) WeekdaysWide() [][]byte {
	return gsw.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'gsw_CH' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gsw.decimal) + len(gsw.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gsw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(gsw.group) - 1; j >= 0; j-- {
					b = append(b, gsw.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(gsw.minus) - 1; j >= 0; j-- {
			b = append(b, gsw.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'gsw_CH' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (gsw *gsw_CH) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gsw.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gsw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(gsw.minus) - 1; j >= 0; j-- {
			b = append(b, gsw.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, gsw.percentSuffix...)

	b = append(b, gsw.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gsw.currencies[currency]
	l := len(s) + len(gsw.decimal) + len(gsw.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gsw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(gsw.group) - 1; j >= 0; j-- {
					b = append(b, gsw.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(gsw.minus) - 1; j >= 0; j-- {
			b = append(b, gsw.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gsw.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, gsw.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'gsw_CH'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gsw.currencies[currency]
	l := len(s) + len(gsw.decimal) + len(gsw.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gsw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(gsw.group) - 1; j >= 0; j-- {
					b = append(b, gsw.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(gsw.minus) - 1; j >= 0; j-- {
			b = append(b, gsw.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gsw.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, gsw.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, gsw.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtDateMedium(t time.Time) []byte {

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
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, gsw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, gsw.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, gsw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'gsw_CH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gsw *gsw_CH) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gsw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := gsw.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
