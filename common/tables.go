package common

import "github.com/jasoncolburne/cesrgo/types"

const (
	Proto_ACDC = types.Proto("ACDC")
	Proto_KERI = types.Proto("KERI")

	Kind_JSON = types.Kind("JSON")
	Kind_CBOR = types.Kind("CBOR")
	Kind_MGPK = types.Kind("MGPK")
	Kind_CESR = types.Kind("CESR")
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
