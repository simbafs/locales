package kok

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type kok struct {
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

// New returns a new instance of translator for the 'kok' locale
func New() locales.Translator {
	return &kok{
		locale:                 "kok",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xa4, 0x9c, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xa8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xab, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9a}, {0xe0, 0xa4, 0x8f, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb2}, {0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x87}, {0xe0, 0xa4, 0x9c, 0xe0, 0xa5, 0x82, 0xe0, 0xa4, 0xa8}, {0xe0, 0xa4, 0x9c, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb2, 0xe0, 0xa5, 0x88}, {0xe0, 0xa4, 0x93, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0x82, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0x93, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xa8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xb5, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0x82, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xa1, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0x82, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbf}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xae}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0x82, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb3}, {0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa7}, {0xe0, 0xa4, 0x97, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x81}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xbf}},
		daysWide:               [][]uint8{{0xe0, 0xa4, 0x86, 0xe0, 0xa4, 0xa6, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xa4, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xaf, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0x82, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb3, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa7, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0x97, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0xe0, 0xa4, 0xae, 0x2e, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x82, 0x2e}, {0xe0, 0xa4, 0xae, 0x2e, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0x82, 0x2e}},
		periodsWide:            [][]uint8{{0xe0, 0xa4, 0xae, 0x2e, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x82, 0x2e}, {0xe0, 0xa4, 0xae, 0x2e, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0x82, 0x2e}},
		erasAbbreviated:        [][]uint8{{0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xa4, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x82, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb5}, {0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xa4, 0xe0, 0xa4, 0xb6, 0xe0, 0xa4, 0x96, 0xe0, 0xa4, 0xbe}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"OEZ": {0x4f, 0x45, 0x5a}, "EST": {0x45, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "CST": {0x43, 0x53, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "CAT": {0x43, 0x41, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "BT": {0x42, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "VET": {0x56, 0x45, 0x54}, "IST": {0xe0, 0xa4, 0xad, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xa4, 0xe0, 0xa5, 0x80, 0xe0, 0xa4, 0xaf, 0x20, 0xe0, 0xa4, 0xb8, 0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0xaf}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "ADT": {0x41, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "CDT": {0x43, 0x44, 0x54}, "WIT": {0x57, 0x49, 0x54}, "UYT": {0x55, 0x59, 0x54}, "WIB": {0x57, 0x49, 0x42}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "EAT": {0x45, 0x41, 0x54}, "JST": {0x4a, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "ART": {0x41, 0x52, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "SRT": {0x53, 0x52, 0x54}, "MST": {0x4d, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "JDT": {0x4a, 0x44, 0x54}},
	}
}

// Locale returns the current translators string locale
func (kok *kok) Locale() string {
	return kok.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kok'
func (kok *kok) PluralsCardinal() []locales.PluralRule {
	return kok.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kok'
func (kok *kok) PluralsOrdinal() []locales.PluralRule {
	return kok.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kok'
func (kok *kok) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kok'
func (kok *kok) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kok'
func (kok *kok) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kok *kok) MonthAbbreviated(month time.Month) []byte {
	return kok.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kok *kok) MonthsAbbreviated() [][]byte {
	return kok.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kok *kok) MonthNarrow(month time.Month) []byte {
	return kok.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kok *kok) MonthsNarrow() [][]byte {
	return kok.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kok *kok) MonthWide(month time.Month) []byte {
	return kok.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kok *kok) MonthsWide() [][]byte {
	return kok.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kok *kok) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return kok.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kok *kok) WeekdaysAbbreviated() [][]byte {
	return kok.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kok *kok) WeekdayNarrow(weekday time.Weekday) []byte {
	return kok.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kok *kok) WeekdaysNarrow() [][]byte {
	return kok.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kok *kok) WeekdayShort(weekday time.Weekday) []byte {
	return kok.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kok *kok) WeekdaysShort() [][]byte {
	return kok.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kok *kok) WeekdayWide(weekday time.Weekday) []byte {
	return kok.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kok *kok) WeekdaysWide() [][]byte {
	return kok.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kok' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(kok.decimal) + len(kok.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(kok.decimal) - 1; j >= 0; j-- {
				b = append(b, kok.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(kok.group) - 1; j >= 0; j-- {
					b = append(b, kok.group[j])
				}

				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(kok.minus) - 1; j >= 0; j-- {
			b = append(b, kok.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kok' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kok *kok) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(kok.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(kok.decimal) - 1; j >= 0; j-- {
				b = append(b, kok.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(kok.minus) - 1; j >= 0; j-- {
			b = append(b, kok.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kok.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kok.currencies[currency]
	l := len(s) + len(kok.decimal) + len(kok.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(kok.decimal) - 1; j >= 0; j-- {
				b = append(b, kok.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(kok.group) - 1; j >= 0; j-- {
					b = append(b, kok.group[j])
				}

				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(kok.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kok.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(kok.minus) - 1; j >= 0; j-- {
			b = append(b, kok.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kok.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kok'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kok.currencies[currency]
	l := len(s) + len(kok.decimal) + len(kok.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(kok.decimal) - 1; j >= 0; j-- {
				b = append(b, kok.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(kok.group) - 1; j >= 0; j-- {
					b = append(b, kok.group[j])
				}

				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(kok.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kok.currencyNegativePrefix[j])
		}

		for j := len(kok.minus) - 1; j >= 0; j-- {
			b = append(b, kok.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kok.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kok.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kok.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kok.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, kok.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kok.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'kok'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kok *kok) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kok.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
