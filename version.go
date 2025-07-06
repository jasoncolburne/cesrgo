// Package cesrgo provides [describe your library's functionality here].
package cesrgo

import "github.com/jasoncolburne/cesrgo/core/types"

// Version represents the current version of the library.
const Version = "2.0.0"

const (
	Proto_ACDC = types.Proto("ACDC")
	Proto_KERI = types.Proto("KERI")

	Kind_JSON = types.Kind("JSON")
	Kind_CBOR = types.Kind("CBOR")
	Kind_MGPK = types.Kind("MGPK")
	Kind_CESR = types.Kind("CESR")

	Ilk_ICP = types.Ilk("icp")
	Ilk_ROT = types.Ilk("rot")
	Ilk_IXN = types.Ilk("ixn")
	Ilk_DIP = types.Ilk("dip")
	Ilk_DRT = types.Ilk("drt")
	Ilk_RCT = types.Ilk("rct")
	Ilk_QRY = types.Ilk("qry")
	Ilk_RPY = types.Ilk("rpy")
	Ilk_XIP = types.Ilk("xip")
	Ilk_EXN = types.Ilk("exn")
	Ilk_PRO = types.Ilk("pro")
	Ilk_BAR = types.Ilk("bar")
	Ilk_VCP = types.Ilk("vcp")
	Ilk_VRT = types.Ilk("vrt")
	Ilk_ISS = types.Ilk("iss")
	Ilk_REV = types.Ilk("rev")
	Ilk_BIS = types.Ilk("bis")
	Ilk_BRV = types.Ilk("brv")
	Ilk_RIP = types.Ilk("rip")
	Ilk_BUP = types.Ilk("bup")
	Ilk_UPD = types.Ilk("upd")
	Ilk_ACM = types.Ilk("acm")
	Ilk_ACT = types.Ilk("act")
	Ilk_ACG = types.Ilk("acg")
	Ilk_ACE = types.Ilk("ace")
	Ilk_SCH = types.Ilk("sch")
	Ilk_ATT = types.Ilk("att")
	Ilk_AGG = types.Ilk("agg")
	Ilk_EDG = types.Ilk("edg")
	Ilk_RUL = types.Ilk("rul")
)

var (
	PROTOS = []types.Proto{
		Proto_ACDC,
		Proto_KERI,
	}

	KINDS = []types.Kind{
		Kind_JSON,
		Kind_CBOR,
		Kind_MGPK,
		Kind_CESR,
	}

	ILKS = []types.Ilk{
		Ilk_ICP,
		Ilk_ROT,
		Ilk_IXN,
		Ilk_DIP,
		Ilk_DRT,
		Ilk_RCT,
		Ilk_QRY,
		Ilk_RPY,
		Ilk_XIP,
		Ilk_EXN,
		Ilk_PRO,
		Ilk_BAR,
		Ilk_VCP,
		Ilk_VRT,
		Ilk_ISS,
		Ilk_REV,
		Ilk_BIS,
		Ilk_BRV,
		Ilk_RIP,
		Ilk_BUP,
		Ilk_UPD,
		Ilk_ACM,
		Ilk_ACT,
		Ilk_ACG,
		Ilk_ACE,
		Ilk_SCH,
		Ilk_ATT,
		Ilk_AGG,
		Ilk_EDG,
		Ilk_RUL,
	}

	VERSION_1_0 = types.Version{
		Major: 1,
		Minor: 0,
	}

	VERSION_2_0 = types.Version{
		Major: 2,
		Minor: 0,
	}

	VERSION = VERSION_2_0
)
