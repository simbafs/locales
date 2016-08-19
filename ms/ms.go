package ms

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ms struct {
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

// New returns a new instance of translator for the 'ms' locale
func New() locales.Translator {
	return &ms{
		locale:                 "ms",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x52, 0x4d}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x46, 0x65, 0x62}, {0x4d, 0x61, 0x63}, {0x41, 0x70, 0x72}, {0x4d, 0x65, 0x69}, {0x4a, 0x75, 0x6e}, {0x4a, 0x75, 0x6c}, {0x4f, 0x67, 0x6f}, {0x53, 0x65, 0x70}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x76}, {0x44, 0x69, 0x73}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x4f}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x69}, {0x46, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72, 0x69}, {0x4d, 0x61, 0x63}, {0x41, 0x70, 0x72, 0x69, 0x6c}, {0x4d, 0x65, 0x69}, {0x4a, 0x75, 0x6e}, {0x4a, 0x75, 0x6c, 0x61, 0x69}, {0x4f, 0x67, 0x6f, 0x73}, {0x53, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x4f, 0x6b, 0x74, 0x6f, 0x62, 0x65, 0x72}, {0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x44, 0x69, 0x73, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x41, 0x68, 0x64}, {0x49, 0x73, 0x6e}, {0x53, 0x65, 0x6c}, {0x52, 0x61, 0x62}, {0x4b, 0x68, 0x61}, {0x4a, 0x75, 0x6d}, {0x53, 0x61, 0x62}},
		daysNarrow:             [][]uint8{{0x41}, {0x49}, {0x53}, {0x52}, {0x4b}, {0x4a}, {0x53}},
		daysShort:              [][]uint8{{0x41, 0x68}, {0x49, 0x73}, {0x53, 0x65}, {0x52, 0x61}, {0x4b, 0x68}, {0x4a, 0x75}, {0x53, 0x61}},
		daysWide:               [][]uint8{{0x41, 0x68, 0x61, 0x64}, {0x49, 0x73, 0x6e, 0x69, 0x6e}, {0x53, 0x65, 0x6c, 0x61, 0x73, 0x61}, {0x52, 0x61, 0x62, 0x75}, {0x4b, 0x68, 0x61, 0x6d, 0x69, 0x73}, {0x4a, 0x75, 0x6d, 0x61, 0x61, 0x74}, {0x53, 0x61, 0x62, 0x74, 0x75}},
		periodsAbbreviated:     [][]uint8{{0x50, 0x47}, {0x50, 0x54, 0x47}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x50, 0x47}, {0x50, 0x54, 0x47}},
		erasAbbreviated:        [][]uint8{{0x53, 0x2e, 0x4d, 0x2e}, {0x54, 0x4d}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x53, 0x2e, 0x4d, 0x2e}, {0x54, 0x4d}},
		timezones:              map[string][]uint8{"ADT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b}, "ACDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "CLT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "COST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "HAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "MDT": {0x4d, 0x44, 0x54}, "ACWST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "VET": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "HNT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "BOT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "TMST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "EST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "UYST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "OEZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x68, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "WEZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x68, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "JDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x4a, 0x65, 0x70, 0x75, 0x6e}, "WAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "MYT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "MST": {0x4d, 0x53, 0x54}, "HAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "BT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "ACWDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "AEDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "SGT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x61}, "WAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "CDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "NZST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "CHADT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "ECT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "IST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "EDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "WESZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x68, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "WITA": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "SRT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "ACST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "HKST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "WIB": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "TMT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "PST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x50, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b}, "PDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x50, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b}, "AST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b}, "ARST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "LHST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "SAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x65, 0x6c, 0x61, 0x74, 0x61, 0x6e}, "GYT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "JST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x4a, 0x65, 0x70, 0x75, 0x6e}, "LHDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "ChST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "HKT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CLST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "NZDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "GFT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x50, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x69, 0x73}, "UYT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "OESZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x68, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "CST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x50, 0x75, 0x73, 0x61, 0x74}, "WIT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "AKDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "∅∅∅": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x42, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x61}, "WART": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "HADT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "MESZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x68, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "AKST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "CHAST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "WARST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x6d, 0x20, 0x50, 0x61, 0x6e, 0x61, 0x73, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "EAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72}, "MEZ": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x45, 0x72, 0x6f, 0x70, 0x61, 0x68, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "AEST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x54, 0x69, 0x6d, 0x75, 0x72, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "ART": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "COT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "CAT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x54, 0x65, 0x6e, 0x67, 0x61, 0x68}, "AWDT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x53, 0x69, 0x61, 0x6e, 0x67, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}, "GMT": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x4d, 0x69, 0x6e, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "AWST": {0x57, 0x61, 0x6b, 0x74, 0x75, 0x20, 0x50, 0x69, 0x61, 0x77, 0x61, 0x69, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x42, 0x61, 0x72, 0x61, 0x74}},
	}
}

// Locale returns the current translators string locale
func (ms *ms) Locale() string {
	return ms.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ms'
func (ms *ms) PluralsCardinal() []locales.PluralRule {
	return ms.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ms'
func (ms *ms) PluralsOrdinal() []locales.PluralRule {
	return ms.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ms'
func (ms *ms) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ms'
func (ms *ms) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ms'
func (ms *ms) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ms *ms) MonthAbbreviated(month time.Month) []byte {
	return ms.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ms *ms) MonthsAbbreviated() [][]byte {
	return ms.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ms *ms) MonthNarrow(month time.Month) []byte {
	return ms.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ms *ms) MonthsNarrow() [][]byte {
	return ms.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ms *ms) MonthWide(month time.Month) []byte {
	return ms.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ms *ms) MonthsWide() [][]byte {
	return ms.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ms *ms) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ms.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ms *ms) WeekdaysAbbreviated() [][]byte {
	return ms.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ms *ms) WeekdayNarrow(weekday time.Weekday) []byte {
	return ms.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ms *ms) WeekdaysNarrow() [][]byte {
	return ms.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ms *ms) WeekdayShort(weekday time.Weekday) []byte {
	return ms.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ms *ms) WeekdaysShort() [][]byte {
	return ms.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ms *ms) WeekdayWide(weekday time.Weekday) []byte {
	return ms.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ms *ms) WeekdaysWide() [][]byte {
	return ms.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ms' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ms.decimal) + len(ms.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ms.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ms.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ms' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ms *ms) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ms.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ms.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ms.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ms.currencies[currency]
	l := len(s) + len(ms.decimal) + len(ms.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ms.group[0])
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
		b = append(b, ms.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ms.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ms'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ms.currencies[currency]
	l := len(s) + len(ms.decimal) + len(ms.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ms.group[0])
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

		for j := len(ms.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ms.currencyNegativePrefix[j])
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
			b = append(b, ms.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ms.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

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

// FmtDateMedium returns the medium date representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ms.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ms.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ms.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ms.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ms.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ms.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ms'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ms *ms) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ms.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ms.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
