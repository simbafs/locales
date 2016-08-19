package to

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type to struct {
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
	currencyPositivePrefix []byte
	currencyNegativePrefix []byte
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

// New returns a new instance of translator for the 'to' locale
func New() locales.Translator {
	return &to{
		locale:                 "to",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0xe2, 0x82, 0xbe}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x24}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x53, 0xc4, 0x81, 0x6e}, {0x46, 0xc4, 0x93, 0x70}, {0x4d, 0x61, 0xca, 0xbb, 0x61}, {0xca, 0xbb, 0x45, 0x70, 0x65}, {0x4d, 0xc4, 0x93}, {0x53, 0x75, 0x6e}, {0x53, 0x69, 0x75}, {0xca, 0xbb, 0x41, 0x6f, 0x6b}, {0x53, 0x65, 0x70}, {0xca, 0xbb, 0x4f, 0x6b, 0x61}, {0x4e, 0xc5, 0x8d, 0x76}, {0x54, 0xc4, 0xab, 0x73}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x53}, {0x46}, {0x4d}, {0x45}, {0x4d}, {0x53}, {0x53}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x54}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x53, 0xc4, 0x81, 0x6e, 0x75, 0x61, 0x6c, 0x69}, {0x46, 0xc4, 0x93, 0x70, 0x75, 0x65, 0x6c, 0x69}, {0x4d, 0x61, 0xca, 0xbb, 0x61, 0x73, 0x69}, {0xca, 0xbb, 0x45, 0x70, 0x65, 0x6c, 0x65, 0x6c, 0x69}, {0x4d, 0xc4, 0x93}, {0x53, 0x75, 0x6e, 0x65}, {0x53, 0x69, 0x75, 0x6c, 0x61, 0x69}, {0xca, 0xbb, 0x41, 0x6f, 0x6b, 0x6f, 0x73, 0x69}, {0x53, 0x65, 0x70, 0x69, 0x74, 0x65, 0x6d, 0x61}, {0xca, 0xbb, 0x4f, 0x6b, 0x61, 0x74, 0x6f, 0x70, 0x61}, {0x4e, 0xc5, 0x8d, 0x76, 0x65, 0x6d, 0x61}, {0x54, 0xc4, 0xab, 0x73, 0x65, 0x6d, 0x61}},
		daysAbbreviated:        [][]uint8{{0x53, 0xc4, 0x81, 0x70}, {0x4d, 0xc5, 0x8d, 0x6e}, {0x54, 0xc5, 0xab, 0x73}, {0x50, 0x75, 0x6c}, {0x54, 0x75, 0xca, 0xbb, 0x61}, {0x46, 0x61, 0x6c}, {0x54, 0x6f, 0x6b}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x54}, {0x50}, {0x54}, {0x46}, {0x54}},
		daysShort:              [][]uint8{{0x53, 0xc4, 0x81, 0x70}, {0x4d, 0xc5, 0x8d, 0x6e}, {0x54, 0xc5, 0xab, 0x73}, {0x50, 0x75, 0x6c}, {0x54, 0x75, 0xca, 0xbb, 0x61}, {0x46, 0x61, 0x6c}, {0x54, 0x6f, 0x6b}},
		daysWide:               [][]uint8{{0x53, 0xc4, 0x81, 0x70, 0x61, 0x74, 0x65}, {0x4d, 0xc5, 0x8d, 0x6e, 0x69, 0x74, 0x65}, {0x54, 0xc5, 0xab, 0x73, 0x69, 0x74, 0x65}, {0x50, 0x75, 0x6c, 0x65, 0x6c, 0x75, 0x6c, 0x75}, {0x54, 0x75, 0xca, 0xbb, 0x61, 0x70, 0x75, 0x6c, 0x65, 0x6c, 0x75, 0x6c, 0x75}, {0x46, 0x61, 0x6c, 0x61, 0x69, 0x74, 0x65}, {0x54, 0x6f, 0x6b, 0x6f, 0x6e, 0x61, 0x6b, 0x69}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x4b, 0x4d}, {0x54, 0x53}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x6b, 0x69, 0x20, 0x6d, 0x75, 0xca, 0xbb, 0x61}, {0x74, 0x61, 0xca, 0xbb, 0x75, 0x20, 0xca, 0xbb, 0x6f, 0x20, 0x53, 0xc4, 0xab, 0x73, 0xc5, 0xab}},
		timezones:              map[string][]uint8{"AWDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "CHADT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x61, 0x74, 0x69, 0x68, 0x61, 0x6d, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "TMT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x74, 0xc5, 0xab, 0x6b, 0x69, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x69, 0x74, 0x61, 0x6e, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "CLST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "HAST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x68, 0x61, 0x75, 0x61, 0xca, 0xbb, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "BT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x70, 0xc5, 0xab, 0x74, 0x61, 0x6e, 0x69}, "SGT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x61}, "GMT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6b, 0x69, 0x6c, 0x69, 0x6e, 0x69, 0x75, 0x69, 0x73, 0x69, 0x20, 0x6d, 0xc4, 0x81, 0x6c, 0x69, 0x65}, "WART": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x73, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "LHST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x75, 0xca, 0xbb, 0x65, 0x69, 0x6b, 0x69, 0x68, 0x6f, 0x75, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "CHAST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x61, 0x74, 0x69, 0x68, 0x61, 0x6d, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "ACWDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x6c, 0x6f, 0x74, 0x6f, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "CLT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "WESZ": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x75, 0x6c, 0x6f, 0x70, 0x65, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "AKST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6c, 0x61, 0x73, 0x69, 0x6b, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "WIB": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x69, 0x6e, 0x69, 0x74, 0x6f, 0x6e, 0x69, 0x73, 0x69, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f}, "VET": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x61}, "EST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "WAST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x66, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "NZST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6e, 0x75, 0xca, 0xbb, 0x75, 0x73, 0x69, 0x6c, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "WITA": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x69, 0x6e, 0x69, 0x74, 0x6f, 0x6e, 0x69, 0x73, 0x69, 0x61, 0x2d, 0x6c, 0x6f, 0x74, 0x6f}, "AST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0xca, 0xbb, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x6b, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "GYT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6b, 0x75, 0x69, 0x61, 0x6e, 0x61}, "MEZ": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x75, 0x6c, 0x6f, 0x70, 0x65, 0x2d, 0x6c, 0x6f, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "LHDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x75, 0xca, 0xbb, 0x65, 0x69, 0x6b, 0x69, 0x68, 0x6f, 0x75, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "SRT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x75, 0x6c, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "OEZ": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x75, 0x6c, 0x6f, 0x70, 0x65, 0x2d, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "JST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x69, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "MYT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x6c, 0x65, 0x69, 0x73, 0x69, 0x61}, "EAT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x66, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65}, "AEST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "HKT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x68, 0x6f, 0x6e, 0x67, 0x69, 0x2d, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "IST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61}, "WEZ": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x75, 0x6c, 0x6f, 0x70, 0x65, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "ART": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x73, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "MDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x6d, 0x6f, 0xca, 0xbb, 0x75, 0x6e, 0x67, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "UYST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x75, 0x6c, 0x75, 0x6b, 0x75, 0x61, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "WARST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x73, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "TMST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x74, 0xc5, 0xab, 0x6b, 0x69, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x69, 0x74, 0x61, 0x6e, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "HADT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x68, 0x61, 0x75, 0x61, 0xca, 0xbb, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "ECT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x6b, 0x75, 0x65, 0x74, 0x6f, 0x61}, "CAT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x66, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x6c, 0x6f, 0x74, 0x6f}, "MST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x6d, 0x6f, 0xca, 0xbb, 0x75, 0x6e, 0x67, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "PDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x70, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "MESZ": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x75, 0x6c, 0x6f, 0x70, 0x65, 0x2d, 0x6c, 0x6f, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "ARST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x73, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "COST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6b, 0x6f, 0x6c, 0x6f, 0x6d, 0x69, 0x70, 0x69, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "BOT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x70, 0x6f, 0x6c, 0xc4, 0xab, 0x76, 0x69, 0x61}, "JDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x73, 0x69, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "HKST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x68, 0x6f, 0x6e, 0x67, 0x69, 0x2d, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "NZDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6e, 0x75, 0xca, 0xbb, 0x75, 0x73, 0x69, 0x6c, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "SAST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x66, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6e, 0x67, 0x61}, "PST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x70, 0x61, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "ADT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0xca, 0xbb, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x6b, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "ACST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x6c, 0x6f, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "ChST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6b, 0x61, 0x6d, 0x6f, 0x6c, 0x6f}, "CDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x6c, 0x6f, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "∅∅∅": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0xc4, 0x81, 0x73, 0x6f, 0x6c, 0x65, 0x73, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "OESZ": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x65, 0x75, 0x6c, 0x6f, 0x70, 0x65, 0x2d, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "ACWST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x6c, 0x6f, 0x74, 0x6f, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "COT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6b, 0x6f, 0x6c, 0x6f, 0x6d, 0x69, 0x70, 0x69, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "HAT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x66, 0x6f, 0x6e, 0x75, 0x61, 0xca, 0xbb, 0x69, 0x6c, 0x6f, 0x66, 0x6f, 0xca, 0xbb, 0x6f, 0x75, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "AWST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "UYT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x75, 0x6c, 0x75, 0x6b, 0x75, 0x61, 0x69, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "CST": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x6c, 0x6f, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "GFT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x6b, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x2d, 0x66, 0x61, 0x6b, 0x61, 0x66, 0x61, 0x6c, 0x61, 0x6e, 0x69, 0x73, 0xc4, 0x93}, "HNT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0x66, 0x6f, 0x6e, 0x75, 0x61, 0xca, 0xbb, 0x69, 0x6c, 0x6f, 0x66, 0x6f, 0xca, 0xbb, 0x6f, 0x75, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "WAT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x66, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x68, 0x69, 0x68, 0x69, 0x66, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x74, 0x6f, 0x74, 0x6f, 0x6e, 0x75}, "WIT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x69, 0x6e, 0x69, 0x74, 0x6f, 0x6e, 0x69, 0x73, 0x69, 0x61, 0x2d, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65}, "AKDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6c, 0x61, 0x73, 0x69, 0x6b, 0x61, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "ACDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x6c, 0x6f, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "AEDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0xc4, 0x93, 0x6c, 0x69, 0x61, 0x2d, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}, "EDT": {0x68, 0x6f, 0x75, 0x61, 0x20, 0x66, 0x61, 0x6b, 0x61, 0xca, 0xbb, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6b, 0x61, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6c, 0x61, 0x75, 0x20, 0x68, 0x61, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x74, 0x61, 0x69, 0x6d, 0x69, 0x20, 0x6c, 0x69, 0x6c, 0x69, 0x75}},
	}
}

// Locale returns the current translators string locale
func (to *to) Locale() string {
	return to.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'to'
func (to *to) PluralsCardinal() []locales.PluralRule {
	return to.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'to'
func (to *to) PluralsOrdinal() []locales.PluralRule {
	return to.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'to'
func (to *to) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'to'
func (to *to) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'to'
func (to *to) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (to *to) MonthAbbreviated(month time.Month) []byte {
	return to.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (to *to) MonthsAbbreviated() [][]byte {
	return to.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (to *to) MonthNarrow(month time.Month) []byte {
	return to.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (to *to) MonthsNarrow() [][]byte {
	return to.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (to *to) MonthWide(month time.Month) []byte {
	return to.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (to *to) MonthsWide() [][]byte {
	return to.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (to *to) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return to.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (to *to) WeekdaysAbbreviated() [][]byte {
	return to.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (to *to) WeekdayNarrow(weekday time.Weekday) []byte {
	return to.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (to *to) WeekdaysNarrow() [][]byte {
	return to.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (to *to) WeekdayShort(weekday time.Weekday) []byte {
	return to.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (to *to) WeekdaysShort() [][]byte {
	return to.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (to *to) WeekdayWide(weekday time.Weekday) []byte {
	return to.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (to *to) WeekdaysWide() [][]byte {
	return to.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'to' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(to.decimal) + len(to.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, to.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, to.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'to' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (to *to) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(to.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, to.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, to.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := to.currencies[currency]
	l := len(s) + len(to.decimal) + len(to.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, to.group[0])
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

	for j := len(to.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, to.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, to.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, to.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'to'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := to.currencies[currency]
	l := len(s) + len(to.decimal) + len(to.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, to.group[0])
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

		for j := len(to.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, to.currencyNegativePrefix[j])
		}

		b = append(b, to.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(to.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, to.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, to.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, to.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, to.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, to.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, to.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, to.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, to.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'to'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (to *to) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, to.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := to.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
