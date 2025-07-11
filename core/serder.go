package cesr

import (
	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/types"
)

var (
	justD = map[types.Field]types.Code{
		"d": codex.Blake3_256,
	}

	dAndI = map[types.Field]types.Code{
		"d": codex.Blake3_256,
		"i": codex.Blake3_256,
	}

	// both v1 and v2
	icp = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"kt", "k", "nt", "n", "bt", "b", "c", "a",
	}, []any{
		"", "", "", "", "0",
		"0", []any{}, "0", []any{}, "0", []any{}, []any{}, []any{},
	})

	rot1 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"p", "kt", "k", "nt", "n", "bt", "br", "ba", "a",
	}, []any{
		"", "", "", "", "0",
		"", "0", []any{}, "0", []any{}, "0", []any{}, []any{}, []any{},
	})

	rot2 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"p", "kt", "k", "nt", "n", "bt", "br", "ba", "c", "a",
	}, []any{
		"", "", "", "", "0",
		"", "0", []any{}, "0", []any{}, "0", []any{}, []any{}, []any{}, []any{},
	})

	// both v1 and v2
	ixn = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"p", "a",
	}, []any{
		"", "", "", "", "0",
		"", []any{},
	})

	dip = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"kt", "k", "nt", "n", "bt", "b", "c", "a", "di",
	}, []any{
		"", "", "", "", "0",
		"0", []any{}, "0", []any{}, "0", []any{}, []any{}, []any{}, "",
	})

	drt1 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"p", "kt", "k", "nt", "n", "bt", "br", "ba", "a",
	}, []any{
		"", "", "", "", "0",
		"", "0", []any{}, "0", []any{}, "0", []any{}, []any{}, []any{},
	})

	drt2 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"p", "kt", "k", "nt", "n", "bt", "br", "ba", "c", "a",
	}, []any{
		"", "", "", "", "0",
		"", "0", []any{}, "0", []any{}, "0", []any{}, []any{}, []any{}, []any{},
	})

	rct = types.NewMap([]string{
		"v", "t", "d", "i", "s",
	}, []any{
		"", "", "", "", "0",
	})

	qry1 = types.NewMap([]string{
		"v", "t", "d", "dt", "r",
		"rr", "q",
	}, []any{
		"", "", "", "", "",
		"", types.NewMap(nil, nil),
	})

	qry2 = types.NewMap([]string{
		"v", "t", "d", "i", "dt",
		"r", "rr", "q",
	}, []any{
		"", "", "", "", "",
		"", "", types.NewMap(nil, nil),
	})

	rpy1 = types.NewMap([]string{
		"v", "t", "d", "dt", "r",
		"a",
	}, []any{
		"", "", "", "", "",
		[]any{},
	})

	rpy2 = types.NewMap([]string{
		"v", "t", "d", "i", "dt",
		"r", "a",
	}, []any{
		"", "", "", "", "",
		"", types.NewMap(nil, []any{}),
	})

	pro1 = types.NewMap([]string{
		"v", "t", "d", "dt", "r",
		"rr", "q",
	}, []any{
		"", "", "", "", "",
		"", types.NewMap(nil, nil),
	})

	pro2 = types.NewMap([]string{
		"v", "t", "d", "i", "dt",
		"r", "rr", "q",
	}, []any{
		"", "", "", "", "",
		"", "", types.NewMap(nil, nil),
	})

	bar1 = types.NewMap([]string{
		"v", "t", "d", "dt", "r",
		"a",
	}, []any{
		"", "", "", "", "",
		[]any{},
	})

	bar2 = types.NewMap([]string{
		"v", "t", "d", "i", "dt",
		"r", "a",
	}, []any{
		"", "", "", "", "",
		"", types.NewMap(nil, []any{}),
	})

	xip = types.NewMap([]string{
		"v", "t", "d", "u", "i", "ri",
		"dt", "r", "q", "a",
	}, []any{
		"", "", "", "", "", "",
		"", "", types.NewMap(nil, nil), types.NewMap(nil, nil),
	})

	exn = types.NewMap([]string{
		"v", "t", "d", "i", "ri",
		"x", "p", "dt", "r", "q", "a",
	}, []any{
		"", "", "", "", "",
		"", "", "", "", types.NewMap(nil, nil), types.NewMap(nil, nil),
	})

	vcp1 = types.NewMap([]string{
		"v", "t", "d", "i", "ii",
		"s", "c", "bt", "b", "n",
	}, []any{
		"", "", "", "", "",
		"0", []any{}, "0", []any{}, "",
	})

	vrt1 = types.NewMap([]string{
		"v", "t", "d", "i", "p",
		"s", "bt", "br", "ba",
	}, []any{
		"", "", "", "", "",
		"0", "0", []any{}, []any{},
	})

	iss1 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"ri", "dt",
	}, []any{
		"", "", "", "", "0",
		"", "",
	})

	rev1 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"ri", "p", "dt",
	}, []any{
		"", "", "", "", "0",
		"", "", "",
	})

	bis1 = types.NewMap([]string{
		"v", "t", "d", "i", "ii",
		"s", "ra", "dt",
	}, []any{
		"", "", "", "", "",
		"0", types.NewMap(nil, nil), "",
	})

	brv1 = types.NewMap([]string{
		"v", "t", "d", "i", "s",
		"p", "ra", "dt",
	}, []any{
		"", "", "", "", "0",
		"", types.NewMap(nil, nil), "",
	})

	alts = types.NewMap([]string{
		"a", "A",
	}, []any{
		"A", "a",
	})

	noneAlls1 = types.NewMap([]string{
		"v", "d", "u", "i",
		"ri", "s", "a", "A", "e", "r",
	}, []any{
		"", "", "", "",
		"", "", "", "", "", "",
	})

	noneAlls2 = types.NewMap([]string{
		"v", "d", "u", "i",
		"rd", "s", "a", "A", "e", "r",
	}, []any{
		"", "", "", "",
		"", "", "", "", "", "",
	})

	noneOpts1 = types.NewMap([]string{
		"u", "ri", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "", "",
	})

	noneOpts2 = types.NewMap([]string{
		"u", "rd", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "", "",
	})

	aceAlls1 = types.NewMap([]string{
		"v", "t", "d", "u", "i",
		"ri", "s", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "",
		"", "", "", "", "", "",
	})

	aceOpts1 = types.NewMap([]string{
		"u", "ri", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "", "",
	})

	aceAlls2 = types.NewMap([]string{
		"v", "t", "d", "u", "i",
		"ri", "s", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "",
		"", "", "", "", "", "",
	})

	aceOpts2 = types.NewMap([]string{
		"u", "ri", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "", "",
	})

	acmAlls = types.NewMap([]string{
		"v", "t", "d", "u", "i",
		"rd", "s", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "",
		"", "", "", "", "", "",
	})

	acmOpts = types.NewMap([]string{
		"t", "u", "rd", "a", "A", "e", "r",
	}, []any{
		"", "", "", "", "", "", "",
	})

	actAlls = types.NewMap([]string{
		"v", "t", "d", "u", "i",
		"rd", "s", "a", "e", "r",
	}, []any{
		"", "", "", "", "",
		"", "", "", "", "",
	})

	acgAlls = types.NewMap([]string{
		"v", "t", "d", "u", "i",
		"rd", "s", "A", "e", "r",
	}, []any{
		"", "", "", "", "",
		"", "", "", "", "",
	})

	schAlls = types.NewMap([]string{
		"v", "t", "d", "s",
	}, []any{
		"", "", "", "",
	})

	attAlls = types.NewMap([]string{
		"v", "t", "d", "a",
	}, []any{
		"", "", "", "",
	})

	aggAlls = types.NewMap([]string{
		"v", "t", "d", "A",
	}, []any{
		"", "", "", "",
	})

	edgAlls = types.NewMap([]string{
		"v", "t", "d", "e",
	}, []any{
		"", "", "", "",
	})

	rulAlls = types.NewMap([]string{
		"v", "t", "d", "r",
	}, []any{
		"", "", "", "",
	})

	ripAlls = types.NewMap([]string{
		"v", "t", "d", "u", "i",
		"n", "dt",
	}, []any{
		"", "", "", "", "",
		"", "",
	})

	bupAlls = types.NewMap([]string{
		"v", "t", "d", "rd", "n",
		"p", "dt", "b",
	}, []any{
		"", "", "", "", "",
		"", "", "",
	})

	updAlls = types.NewMap([]string{
		"v", "t", "d", "rd", "n",
		"p", "dt", "td", "ts",
	}, []any{
		"", "", "", "", "",
		"", "", "", "",
	})
)

var fields = map[types.Proto]map[uint32]map[types.Ilk]struct {
	alls   types.Map
	opts   types.Map
	alts   types.Map
	saids  map[types.Field]types.Code
	strict bool
}{
	cesrgo.Proto_KERI: {
		cesrgo.VERSION_1_0.Major: {
			cesrgo.Ilk_ICP: {
				alls:  icp,
				saids: dAndI,
			},
			cesrgo.Ilk_ROT: {
				alls:  rot1,
				saids: justD,
			},
			cesrgo.Ilk_IXN: {
				alls:  ixn,
				saids: justD,
			},
			cesrgo.Ilk_DIP: {
				alls:  dip,
				saids: dAndI,
			},
			cesrgo.Ilk_DRT: {
				alls:  drt1,
				saids: justD,
			},
			cesrgo.Ilk_RCT: {
				alls:  rct,
				saids: justD,
			},
			cesrgo.Ilk_QRY: {
				alls:  qry1,
				saids: justD,
			},
			cesrgo.Ilk_RPY: {
				alls:  rpy1,
				saids: justD,
			},
			cesrgo.Ilk_PRO: {
				alls:  pro1,
				saids: justD,
			},
			cesrgo.Ilk_BAR: {
				alls:  bar1,
				saids: justD,
			},
			cesrgo.Ilk_VCP: {
				alls:  vcp1,
				saids: dAndI,
			},
			cesrgo.Ilk_VRT: {
				alls:  vrt1,
				saids: justD,
			},
			cesrgo.Ilk_ISS: {
				alls:  iss1,
				saids: justD,
			},
			cesrgo.Ilk_REV: {
				alls:  rev1,
				saids: justD,
			},
			cesrgo.Ilk_BIS: {
				alls:  bis1,
				saids: justD,
			},
			cesrgo.Ilk_BRV: {
				alls:  brv1,
				saids: justD,
			},
		},
		cesrgo.VERSION_2_0.Major: {
			cesrgo.Ilk_ICP: {
				alls:  icp,
				saids: dAndI,
			},
			cesrgo.Ilk_ROT: {
				alls:  rot2,
				saids: justD,
			},
			cesrgo.Ilk_IXN: {
				alls:  ixn,
				saids: justD,
			},
			cesrgo.Ilk_DIP: {
				alls:  dip,
				saids: dAndI,
			},
			cesrgo.Ilk_DRT: {
				alls:  drt2,
				saids: justD,
			},
			cesrgo.Ilk_RCT: {
				alls:  rct,
				saids: justD,
			},
			cesrgo.Ilk_QRY: {
				alls:  qry2,
				saids: justD,
			},
			cesrgo.Ilk_RPY: {
				alls:  rpy2,
				saids: justD,
			},
			cesrgo.Ilk_PRO: {
				alls:  pro2,
				saids: justD,
			},
			cesrgo.Ilk_BAR: {
				alls:  bar2,
				saids: justD,
			},
			cesrgo.Ilk_XIP: {
				alls:  xip,
				saids: justD,
			},
			cesrgo.Ilk_EXN: {
				alls:  exn,
				saids: justD,
			},
		},
	},
	cesrgo.Proto_ACDC: {
		cesrgo.VERSION_1_0.Major: {
			cesrgo.Ilk_NONE: {
				alls:   noneAlls1,
				opts:   noneOpts1,
				alts:   alts,
				saids:  justD,
				strict: true,
			},
			cesrgo.Ilk_ACE: {
				alls:   aceAlls1,
				opts:   aceOpts1,
				alts:   alts,
				saids:  justD,
				strict: false,
			},
		},
		cesrgo.VERSION_2_0.Major: {
			cesrgo.Ilk_NONE: {
				alls:  noneAlls2,
				opts:  noneOpts2,
				alts:  alts,
				saids: justD,
			},
			cesrgo.Ilk_ACM: {
				alls:  acmAlls,
				opts:  acmOpts,
				alts:  alts,
				saids: justD,
			},
			cesrgo.Ilk_ACE: {
				alls:  aceAlls2,
				opts:  aceOpts2,
				alts:  alts,
				saids: justD,
			},
			cesrgo.Ilk_ACT: {
				alls:  actAlls,
				saids: justD,
			},
			cesrgo.Ilk_ACG: {
				alls:  acgAlls,
				saids: justD,
			},
			cesrgo.Ilk_SCH: {
				alls:  schAlls,
				saids: justD,
			},
			cesrgo.Ilk_ATT: {
				alls:  attAlls,
				saids: justD,
			},
			cesrgo.Ilk_AGG: {
				alls:  aggAlls,
				saids: justD,
			},
			cesrgo.Ilk_EDG: {
				alls:  edgAlls,
				saids: justD,
			},
			cesrgo.Ilk_RUL: {
				alls:  rulAlls,
				saids: justD,
			},
			cesrgo.Ilk_RIP: {
				alls:  ripAlls,
				saids: justD,
			},
			cesrgo.Ilk_BUP: {
				alls:  bupAlls,
				saids: justD,
			},
			cesrgo.Ilk_UPD: {
				alls:  updAlls,
				saids: justD,
			},
		},
	},
}

type Serder struct {
	raw   types.Raw
	sad   types.Map
	proto types.Proto
	pvrsn types.Version
	genus string
	gvrsn types.Version
	kind  types.Kind
	size  types.Size
}
