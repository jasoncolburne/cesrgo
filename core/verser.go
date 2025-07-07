package cesr

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Verser struct {
	Tagger
}

func NewVerser(
	versage *types.Versage,
	proto *types.Proto,
	pvrsn *types.Version,
	gvrsn *types.Version,
	opts ...options.MatterOption,
) (*Verser, error) {
	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	var tag *string
	if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if versage != nil {
			if proto != nil || pvrsn != nil || gvrsn != nil {
				return nil, fmt.Errorf("versage must be nil if proto, pvrsn, or gvrsn is not nil")
			}

			proto = &versage.Proto
			pvrsn = &versage.Pvrsn
			gvrsn = versage.Gvrsn
		}

		pvtag, err := verToB64(pvrsn, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to convert pvrsn to b64: %v", err)
		}

		tagStr := string(*proto) + string(pvtag)

		if gvrsn != nil {
			gvtag, err := verToB64(gvrsn, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to convert gvrsn to b64: %v", err)
			}

			tagStr += string(gvtag)
		}

		tag = &tagStr
	}

	tagger, err := NewTagger(tag, opts...)
	if err != nil {
		return nil, err
	}

	i := &Verser{Tagger: *tagger}

	if !slices.Contains([]types.Code{codex.Tag7, codex.Tag10}, i.GetCode()) {
		return nil, fmt.Errorf("verser must be coded as tag7 or tag10")
	}

	return i, nil
}

func (i *Verser) Versage() (types.Versage, error) {
	tag, err := i.Tag()
	if err != nil {
		return types.Versage{}, fmt.Errorf("failed to get tag: %v", err)
	}

	proto := types.Proto(tag[:4])
	pvrsn, err := b64ToVer(tag[4:7])

	var gvrsn *types.Version
	if len(tag) == 10 {
		ver, err := b64ToVer(tag[7:10])
		if err != nil {
			return types.Versage{}, fmt.Errorf("failed to convert gvrsn to version: %v", err)
		}

		gvrsn = &ver
	}

	versage := types.Versage{
		Proto: proto,
		Pvrsn: pvrsn,
		Gvrsn: gvrsn,
	}

	return versage, nil
}

func verToB64(version *types.Version, text *string) (types.Qb64, error) {
	if version == nil && text == nil {
		return "", fmt.Errorf("version or text is required")
	}

	var major uint32
	var minor uint32

	if version != nil {
		major = version.Major
		minor = version.Minor
	}

	if text != nil {
		splitStrs := strings.Split(*text, ".")
		if len(splitStrs) > 2 {
			return "", fmt.Errorf("invalid version text: %s", *text)
		}

		splits := []uint32{}
		for _, splitStr := range splitStrs {
			split, err := strconv.ParseUint(splitStr, 10, 32)
			if err != nil {
				return "", fmt.Errorf("invalid version text: %s", *text)
			}

			splits = append(splits, uint32(split))
		}

		parts := []uint32{major, minor}

		for i := 2 - len(splits); i > 0; i-- {
			splits = append(splits, parts[i])
		}

		major = splits[0]
		minor = splits[1]
	}

	if major > 63 || minor > 4095 {
		return "", fmt.Errorf("invalid version: %d.%d", major, minor)
	}

	_major, err := common.BigIntToB64(big.NewInt(int64(major)), 1)
	if err != nil {
		return "", fmt.Errorf("failed to convert major to b64: %v", err)
	}

	_minor, err := common.BigIntToB64(big.NewInt(int64(minor)), 2)
	if err != nil {
		return "", fmt.Errorf("failed to convert minor to b64: %v", err)
	}

	qb64 := fmt.Sprintf("%s%s", _major, _minor)
	return types.Qb64(qb64), nil
}

func b64ToVer(b64 string) (types.Version, error) {
	re, err := common.Reb64()
	if err != nil {
		return types.Version{}, fmt.Errorf("failed to create re: %v", err)
	}

	if !re.MatchString(b64) {
		return types.Version{}, fmt.Errorf("invalid b64: %s", b64)
	}

	major, err := common.B64ToU32(b64[:1])
	if err != nil {
		return types.Version{}, fmt.Errorf("failed to convert major to int: %v", err)
	}

	minor, err := common.B64ToU32(b64[1:3])
	if err != nil {
		return types.Version{}, fmt.Errorf("failed to convert minor to int: %v", err)
	}

	return types.Version{
		Major: major,
		Minor: minor,
	}, nil
}
