package indexer

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/types"
)

var SMALL_VRZ_BYTES = uint32(3)
var LARGE_VRZ_BYTES = uint32(6)

var (
	Ed25519             = types.Code("A")  // Ed25519 sig appears same in both lists if any.
	Ed25519_Crt         = types.Code("B")  // Ed25519 sig appears in current list only.
	ECDSA_256k1         = types.Code("C")  // ECDSA secp256k1 sig appears same in both lists if any.
	ECDSA_256k1_Crt     = types.Code("D")  // ECDSA secp256k1 sig appears in current list.
	ECDSA_256r1         = types.Code("E")  // ECDSA secp256r1 sig appears same in both lists if any.
	ECDSA_256r1_Crt     = types.Code("F")  // ECDSA secp256r1 sig appears in current list.
	Ed448               = types.Code("0A") // Ed448 signature appears in both lists.
	Ed448_Crt           = types.Code("0B") // Ed448 signature appears in current list only.
	Ed25519_Big         = types.Code("2A") // Ed25519 sig appears in both lists.
	Ed25519_Big_Crt     = types.Code("2B") // Ed25519 sig appears in current list only.
	ECDSA_256k1_Big     = types.Code("2C") // ECDSA secp256k1 sig appears in both lists.
	ECDSA_256k1_Big_Crt = types.Code("2D") // ECDSA secp256k1 sig appears in current list only.
	ECDSA_256r1_Big     = types.Code("2E") // ECDSA secp256r1 sig appears in both lists.
	ECDSA_256r1_Big_Crt = types.Code("2F") // ECDSA secp256r1 sig appears in current list only.
	Ed448_Big           = types.Code("3A") // Ed448 signature appears in both lists.
	Ed448_Big_Crt       = types.Code("3B") // Ed448 signature appears in current list only.
	TBD0                = types.Code("0z") // Test of Var len label L=N*4 <= 4095 char quadlets includes code
	TBD1                = types.Code("1z") // Test of index sig lead 1
	TBD4                = types.Code("4z") // Test of index sig lead 1 big
)

var ValidSigCodes = []types.Code{
	Ed25519,
	Ed25519_Crt,
	ECDSA_256k1,
	ECDSA_256k1_Crt,
	ECDSA_256r1,
	ECDSA_256r1_Crt,
	Ed448,
	Ed448_Crt,
	Ed25519_Big,
	Ed25519_Big_Crt,
	ECDSA_256k1_Big,
	ECDSA_256k1_Big_Crt,
	ECDSA_256r1_Big,
	ECDSA_256r1_Big_Crt,
	Ed448_Big,
	Ed448_Big_Crt,
}

var ValidCurrentSigCodes = []types.Code{
	Ed25519_Crt,
	ECDSA_256k1_Crt,
	ECDSA_256r1_Crt,
	Ed448_Crt,
	Ed25519_Big_Crt,
	ECDSA_256k1_Big_Crt,
	ECDSA_256r1_Big_Crt,
	Ed448_Big_Crt,
}

var ValidBothSigCodes = []types.Code{
	Ed25519,
	ECDSA_256k1,
	ECDSA_256r1,
	Ed448,
	Ed25519_Big,
	ECDSA_256k1_Big,
	ECDSA_256r1_Big,
	Ed448_Big,
}

type Sizage struct {
	Hs uint32
	Ss uint32
	Os uint32
	Fs *uint32
	Ls uint32
}

func GetSizage(code types.Code) (Sizage, error) {
	switch code {
	case Ed25519, Ed25519_Crt, ECDSA_256k1, ECDSA_256k1_Crt, ECDSA_256r1, ECDSA_256r1_Crt:
		fs := uint32(88)
		return Sizage{Hs: 1, Ss: 1, Os: 0, Fs: &fs, Ls: 0}, nil
	case Ed448, Ed448_Crt:
		fs := uint32(156)
		return Sizage{Hs: 2, Ss: 2, Os: 1, Fs: &fs, Ls: 0}, nil
	case Ed25519_Big, Ed25519_Big_Crt, ECDSA_256k1_Big, ECDSA_256k1_Big_Crt, ECDSA_256r1_Big, ECDSA_256r1_Big_Crt:
		fs := uint32(92)
		return Sizage{Hs: 2, Ss: 4, Os: 2, Fs: &fs, Ls: 0}, nil
	case Ed448_Big, Ed448_Big_Crt:
		fs := uint32(160)
		return Sizage{Hs: 2, Ss: 6, Os: 3, Fs: &fs, Ls: 0}, nil
	case TBD0:
		return Sizage{Hs: 2, Ss: 2, Os: 0, Ls: 0}, nil
	case TBD1:
		fs := uint32(76)
		return Sizage{Hs: 2, Ss: 2, Os: 1, Fs: &fs, Ls: 1}, nil
	case TBD4:
		fs := uint32(80)
		return Sizage{Hs: 2, Ss: 6, Os: 3, Fs: &fs, Ls: 1}, nil
	default:
		return Sizage{}, fmt.Errorf("unknown sizage: %s", code)
	}
}

func GetHardage(c byte) (uint32, error) {
	if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
		return 1, nil
	}
	if c >= '0' && c <= '4' {
		return 2, nil
	}

	if c == '-' {
		return 0, fmt.Errorf("count code start")
	}
	if c == '_' {
		return 0, fmt.Errorf("op code start")
	}

	return 0, fmt.Errorf("unknown hardage: %c", c)
}

func GetBardage(b byte) (uint32, error) {
	if b <= 0x33 {
		return 1, nil
	}
	if b >= 0x34 && b <= 0x38 {
		return 2, nil
	}

	if b == 0x3e {
		return 0, fmt.Errorf("count code start")
	}
	if b == 0x3f {
		return 0, fmt.Errorf("op code start")
	}

	return 0, fmt.Errorf("unknown bardage: %x", b)
}
