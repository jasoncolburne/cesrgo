package indexer

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/types"
	"github.com/jasoncolburne/cesrgo/util"
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

var (
	_76  = uint32(76)
	_80  = uint32(80)
	_88  = uint32(88)
	_92  = uint32(92)
	_156 = uint32(156)
	_160 = uint32(160)
)

var Sizes = map[types.Code]Sizage{
	Ed25519:         {Hs: 1, Ss: 1, Os: 0, Fs: &_88, Ls: 0},
	Ed25519_Crt:     {Hs: 1, Ss: 1, Os: 0, Fs: &_88, Ls: 0},
	ECDSA_256k1:     {Hs: 1, Ss: 1, Os: 0, Fs: &_88, Ls: 0},
	ECDSA_256k1_Crt: {Hs: 1, Ss: 1, Os: 0, Fs: &_88, Ls: 0},
	ECDSA_256r1:     {Hs: 1, Ss: 1, Os: 0, Fs: &_88, Ls: 0},
	ECDSA_256r1_Crt: {Hs: 1, Ss: 1, Os: 0, Fs: &_88, Ls: 0},

	Ed25519_Big:         {Hs: 2, Ss: 4, Os: 2, Fs: &_92, Ls: 0},
	Ed25519_Big_Crt:     {Hs: 2, Ss: 4, Os: 2, Fs: &_92, Ls: 0},
	ECDSA_256k1_Big:     {Hs: 2, Ss: 4, Os: 2, Fs: &_92, Ls: 0},
	ECDSA_256k1_Big_Crt: {Hs: 2, Ss: 4, Os: 2, Fs: &_92, Ls: 0},
	ECDSA_256r1_Big:     {Hs: 2, Ss: 4, Os: 2, Fs: &_92, Ls: 0},
	ECDSA_256r1_Big_Crt: {Hs: 2, Ss: 4, Os: 2, Fs: &_92, Ls: 0},

	Ed448:     {Hs: 2, Ss: 2, Os: 1, Fs: &_156, Ls: 0},
	Ed448_Crt: {Hs: 2, Ss: 2, Os: 1, Fs: &_156, Ls: 0},

	Ed448_Big:     {Hs: 2, Ss: 6, Os: 3, Fs: &_160, Ls: 0},
	Ed448_Big_Crt: {Hs: 2, Ss: 6, Os: 3, Fs: &_160, Ls: 0},

	TBD0: {Hs: 2, Ss: 2, Os: 0, Ls: 0},
	TBD1: {Hs: 2, Ss: 2, Os: 1, Fs: &_76, Ls: 1},
	TBD4: {Hs: 2, Ss: 6, Os: 3, Fs: &_80, Ls: 0},
}

var Hards = map[byte]int{}
var Bards = map[byte]int{}

func generateHards() {
	if len(Hards) > 0 {
		return
	}

	for c := 'A'; c <= 'Z'; c++ {
		Hards[byte(c)] = 1
	}

	for c := 'a'; c <= 'z'; c++ {
		Hards[byte(c)] = 1
	}

	for c := '0'; c <= '4'; c++ {
		Hards[byte(c)] = 2
	}
}

func generateBards() error {
	if len(Bards) > 0 {
		return nil
	}

	generateHards()

	for hard, i := range Hards {
		bard, err := util.CodeB64ToB2(string(hard))
		if err != nil {
			return err
		}

		if len(bard) != 1 {
			return fmt.Errorf("unexpected bard length: %d", len(bard))
		}
		Bards[bard[0]] = i
	}

	return nil
}

func Hardage(c byte) (int, bool) {
	generateHards()

	n, ok := Hards[c]
	return n, ok
}

func Bardage(b byte) (int, bool) {
	err := generateBards()
	if err != nil {
		return -1, false
	}

	n, ok := Bards[b]
	return n, ok
}
