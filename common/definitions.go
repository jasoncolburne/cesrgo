package common

import (
	"regexp"

	"github.com/jasoncolburne/cesrgo/core/types"
)

var (
	VER1FULLSPAN = 17
	VER1TERM     = '_'
	VEREX1       = "([A-Z]{4})([0-9a-f])([0-9a-f])([A-Z]{4})([0-9a-f]{6})_"

	VER2FULLSPAN = 19
	VER2TERM     = '.'
	VEREX2       = "([A-Z]{4})([0-9A-Za-z_-])([0-9A-Za-z_-]{2})([0-9A-Za-z_-])([0-9A-Za-z_-]{2})([A-Z]{4})([0-9A-Za-z_-]{4})\\."

	MAXVERFULLSPAN = max(VER1FULLSPAN, VER2FULLSPAN)
	MAXVSOFFSET    = 12

	SMELLSIZE = MAXVSOFFSET + MAXVERFULLSPAN

	B64RUNES   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	B64INDICES = map[rune]uint8{}

	B64EX   = "^[A-Za-z0-9_-]*$"
	PATHREX = "^[a-zA-Z0-9_]*$"
	ATREX   = "^[a-zA-Z_][a-zA-Z0-9_]*$"

	REVER  *regexp.Regexp
	REB64  *regexp.Regexp
	REATT  *regexp.Regexp
	REPATH *regexp.Regexp

	TIER_LOW  = types.Tier("low")
	TIER_MED  = types.Tier("med")
	TIER_HIGH = types.Tier("high")
)
