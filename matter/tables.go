package matter

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/types"
)

const (
	Ed25519_Seed              = types.Code("A")    // Ed25519 256 bit random seed for private key
	Ed25519N                  = types.Code("B")    // Ed25519 verification key non-transferable, basic derivation.
	X25519                    = types.Code("C")    // X25519 public encryption key, may be converted from Ed25519 or Ed25519N.
	Ed25519                   = types.Code("D")    // Ed25519 verification key basic derivation
	Blake3_256                = types.Code("E")    // Blake3 256 bit digest self-addressing derivation.
	Blake2b_256               = types.Code("F")    // Blake2b 256 bit digest self-addressing derivation.
	Blake2s_256               = types.Code("G")    // Blake2s 256 bit digest self-addressing derivation.
	SHA3_256                  = types.Code("H")    // SHA3 256 bit digest self-addressing derivation.
	SHA2_256                  = types.Code("I")    // SHA2 256 bit digest self-addressing derivation.
	ECDSA_256k1_Seed          = types.Code("J")    // ECDSA secp256k1 256 bit random Seed for private key
	Ed448_Seed                = types.Code("K")    // Ed448 448 bit random Seed for private key
	X448                      = types.Code("L")    // X448 public encryption key, converted from Ed448
	Short                     = types.Code("M")    // Short 2 byte b2 number
	Big                       = types.Code("N")    // Big 8 byte b2 number
	X25519_Private            = types.Code("O")    // X25519 private decryption key/seed, may be converted from Ed25519
	X25519_Cipher_Seed        = types.Code("P")    // X25519 sealed box 124 char qb64 Cipher of 44 char qb64 Seed
	ECDSA_256r1_Seed          = types.Code("Q")    // ECDSA secp256r1 256 bit random Seed for private key
	Tall                      = types.Code("R")    // Tall 5 byte b2 number
	Large                     = types.Code("S")    // Large 11 byte b2 number
	Great                     = types.Code("T")    // Great 14 byte b2 number
	Vast                      = types.Code("U")    // Vast 17 byte b2 number
	Label1                    = types.Code("V")    // Label1 1 bytes for label lead size 1
	Label2                    = types.Code("W")    // Label2 2 bytes for label lead size 0
	Tag3                      = types.Code("X")    // Tag3  3 B64 encoded chars for special values
	Tag7                      = types.Code("Y")    // Tag7  7 B64 encoded chars for special values
	Tag11                     = types.Code("Z")    // Tag11  11 B64 encoded chars for special values
	Salt_256                  = types.Code("a")    // Salt/seed/nonce/blind 256 bits
	Salt_128                  = types.Code("0A")   // Salt/seed/nonce 128 bits or number of length 128 bits (Huge)
	Ed25519_Sig               = types.Code("0B")   // Ed25519 signature.
	ECDSA_256k1_Sig           = types.Code("0C")   // ECDSA secp256k1 signature.
	Blake3_512                = types.Code("0D")   // Blake3 512 bit digest self-addressing derivation.
	Blake2b_512               = types.Code("0E")   // Blake2b 512 bit digest self-addressing derivation.
	SHA3_512                  = types.Code("0F")   // SHA3 512 bit digest self-addressing derivation.
	SHA2_512                  = types.Code("0G")   // SHA2 512 bit digest self-addressing derivation.
	Long                      = types.Code("0H")   // Long 4 byte b2 number
	ECDSA_256r1_Sig           = types.Code("0I")   // ECDSA secp256r1 signature.
	Tag1                      = types.Code("0J")   // Tag1 1 B64 encoded char + 1 prepad for special values
	Tag2                      = types.Code("0K")   // Tag2 2 B64 encoded chars for for special values
	Tag5                      = types.Code("0L")   // Tag5 5 B64 encoded chars + 1 prepad for special values
	Tag6                      = types.Code("0M")   // Tag6 6 B64 encoded chars for special values
	Tag9                      = types.Code("0N")   // Tag9 9 B64 encoded chars + 1 prepad for special values
	Tag10                     = types.Code("0O")   // Tag10 10 B64 encoded chars for special values
	GramHeadNeck              = types.Code("0P")   // GramHeadNeck 32 B64 chars memogram head with neck
	GramHead                  = types.Code("0Q")   // GramHead 28 B64 chars memogram head only
	GramHeadAIDNeck           = types.Code("0R")   // GramHeadAIDNeck 76 B64 chars memogram head with AID and neck
	GramHeadAID               = types.Code("0S")   // GramHeadAID 72 B64 chars memogram head with AID only
	ECDSA_256k1N              = types.Code("1AAA") // ECDSA secp256k1 verification key non-transferable, basic derivation.
	ECDSA_256k1               = types.Code("1AAB") // ECDSA public verification or encryption key, basic derivation
	Ed448N                    = types.Code("1AAC") // Ed448 non-transferable prefix public signing verification key. Basic derivation.
	Ed448                     = types.Code("1AAD") // Ed448 public signing verification key. Basic derivation.
	Ed448_Sig                 = types.Code("1AAE") // Ed448 signature. Self-signing derivation.
	Tag4                      = types.Code("1AAF") // Tag4 4 B64 encoded chars for special values
	DateTime                  = types.Code("1AAG") // Base64 custom encoded 32 char ISO-8601 DateTime
	X25519_Cipher_Salt        = types.Code("1AAH") // X25519 sealed box 100 char qb64 Cipher of 24 char qb64 Salt
	ECDSA_256r1N              = types.Code("1AAI") // ECDSA secp256r1 verification key non-transferable, basic derivation.
	ECDSA_256r1               = types.Code("1AAJ") // ECDSA secp256r1 verification or encryption key, basic derivation
	Null                      = types.Code("1AAK") // Null None or empty value
	No                        = types.Code("1AAL") // No Falsey Boolean value
	Yes                       = types.Code("1AAM") // Yes Truthy Boolean value
	Tag8                      = types.Code("1AAN") // Tag8 8 B64 encoded chars for special values
	Escape                    = types.Code("1AAO") // Escape code for escaping special map fields
	Empty                     = types.Code("1AAP") // Empty value for Nonce, UUID, or related fields
	TBD0S                     = types.Code("1__-") // Testing purposes only, fixed special values with non-empty raw lead size 0
	TBD0                      = types.Code("1___") // Testing purposes only, fixed with lead size 0
	TBD1S                     = types.Code("2__-") // Testing purposes only, fixed special values with non-empty raw lead size 1
	TBD1                      = types.Code("2___") // Testing purposes only, fixed with lead size 1
	TBD2S                     = types.Code("3__-") // Testing purposes only, fixed special values with non-empty raw lead size 2
	TBD2                      = types.Code("3___") // Testing purposes only, fixed with lead size 2
	StrB64_L0                 = types.Code("4A")   // String Base64 only lead size 0
	StrB64_L1                 = types.Code("5A")   // String Base64 only lead size 1
	StrB64_L2                 = types.Code("6A")   // String Base64 only lead size 2
	StrB64_Big_L0             = types.Code("7AAA") // String Base64 only big lead size 0
	StrB64_Big_L1             = types.Code("8AAA") // String Base64 only big lead size 1
	StrB64_Big_L2             = types.Code("9AAA") // String Base64 only big lead size 2
	Bytes_L0                  = types.Code("4B")   // Byte String lead size 0
	Bytes_L1                  = types.Code("5B")   // Byte String lead size 1
	Bytes_L2                  = types.Code("6B")   // Byte String lead size 2
	Bytes_Big_L0              = types.Code("7AAB") // Byte String big lead size 0
	Bytes_Big_L1              = types.Code("8AAB") // Byte String big lead size 1
	Bytes_Big_L2              = types.Code("9AAB") // Byte String big lead size 2
	X25519_Cipher_L0          = types.Code("4C")   // X25519 sealed box cipher bytes of sniffable stream plaintext lead size 0
	X25519_Cipher_L1          = types.Code("5C")   // X25519 sealed box cipher bytes of sniffable stream plaintext lead size 1
	X25519_Cipher_L2          = types.Code("6C")   // X25519 sealed box cipher bytes of sniffable stream plaintext lead size 2
	X25519_Cipher_Big_L0      = types.Code("7AAC") // X25519 sealed box cipher bytes of sniffable stream plaintext big lead size 0
	X25519_Cipher_Big_L1      = types.Code("8AAC") // X25519 sealed box cipher bytes of sniffable stream plaintext big lead size 1
	X25519_Cipher_Big_L2      = types.Code("9AAC") // X25519 sealed box cipher bytes of sniffable stream plaintext big lead size 2
	X25519_Cipher_QB64_L0     = types.Code("4D")   // X25519 sealed box cipher bytes of QB64 plaintext lead size 0
	X25519_Cipher_QB64_L1     = types.Code("5D")   // X25519 sealed box cipher bytes of QB64 plaintext lead size 1
	X25519_Cipher_QB64_L2     = types.Code("6D")   // X25519 sealed box cipher bytes of QB64 plaintext lead size 2
	X25519_Cipher_QB64_Big_L0 = types.Code("7AAD") // X25519 sealed box cipher bytes of QB64 plaintext big lead size 0
	X25519_Cipher_QB64_Big_L1 = types.Code("8AAD") // X25519 sealed box cipher bytes of QB64 plaintext big lead size 1
	X25519_Cipher_QB64_Big_L2 = types.Code("9AAD") // X25519 sealed box cipher bytes of QB64 plaintext big lead size 2
	X25519_Cipher_QB2_L0      = types.Code("4E")   // X25519 sealed box cipher bytes of QB2 plaintext lead size 0
	X25519_Cipher_QB2_L1      = types.Code("5E")   // X25519 sealed box cipher bytes of QB2 plaintext lead size 1
	X25519_Cipher_QB2_L2      = types.Code("6E")   // X25519 sealed box cipher bytes of QB2 plaintext lead size 2
	X25519_Cipher_QB2_Big_L0  = types.Code("7AAE") // X25519 sealed box cipher bytes of QB2 plaintext big lead size 0
	X25519_Cipher_QB2_Big_L1  = types.Code("8AAE") // X25519 sealed box cipher bytes of QB2 plaintext big lead size 1
	X25519_Cipher_QB2_Big_L2  = types.Code("9AAE") // X25519 sealed box cipher bytes of QB2 plaintext big lead size 2
	HPKEBase_Cipher_L0        = types.Code("4F")   // HPKE Base cipher bytes of sniffable stream plaintext lead size 0
	HPKEBase_Cipher_L1        = types.Code("5F")   // HPKE Base cipher bytes of sniffable stream plaintext lead size 1
	HPKEBase_Cipher_L2        = types.Code("6F")   // HPKE Base cipher bytes of sniffable stream plaintext lead size 2
	HPKEBase_Cipher_Big_L0    = types.Code("7AAF") // HPKE Base cipher bytes of sniffable stream plaintext big lead size 0
	HPKEBase_Cipher_Big_L1    = types.Code("8AAF") // HPKE Base cipher bytes of sniffable stream plaintext big lead size 1
	HPKEBase_Cipher_Big_L2    = types.Code("9AAF") // HPKE Base cipher bytes of sniffable stream plaintext big lead size 2
	HPKEAuth_Cipher_L0        = types.Code("4G")   // HPKE Auth cipher bytes of sniffable stream plaintext lead size 0
	HPKEAuth_Cipher_L1        = types.Code("5G")   // HPKE Auth cipher bytes of sniffable stream plaintext lead size 1
	HPKEAuth_Cipher_L2        = types.Code("6G")   // HPKE Auth cipher bytes of sniffable stream plaintext lead size 2
	HPKEAuth_Cipher_Big_L0    = types.Code("7AAG") // HPKE Auth cipher bytes of sniffable stream plaintext big lead size 0
	HPKEAuth_Cipher_Big_L1    = types.Code("8AAG") // HPKE Auth cipher bytes of sniffable stream plaintext big lead size 1
	HPKEAuth_Cipher_Big_L2    = types.Code("9AAG") // HPKE Auth cipher bytes of sniffable stream plaintext big lead size 2
	Decimal_L0                = types.Code("4H")   // Decimal B64 string float and int lead size 0
	Decimal_L1                = types.Code("5H")   // Decimal B64 string float and int lead size 1
	Decimal_L2                = types.Code("6H")   // Decimal B64 string float and intlead size 2
	Decimal_Big_L0            = types.Code("7AAH") // Decimal B64 string float and int big lead size 0
	Decimal_Big_L1            = types.Code("8AAH") // Decimal B64 string float and int big lead size 1
	Decimal_Big_L2            = types.Code("9AAH") // Decimal B64 string float and int big lead size 2
)

var SMALL_VRZ_DEX = []byte{4, 5, 6}
var LARGE_VRZ_DEX = []byte{7, 8, 9}
var SMALL_VRZ_BYTES = uint32(3)
var LARGE_VRZ_BYTES = uint32(6)

type Sizage struct {
	Hs uint32
	Ss uint32
	Xs uint32
	Fs *uint32
	Ls uint32
}

func GetSizage(code types.Code) (Sizage, error) {
	var fs uint32

	switch code {
	case Ed25519_Seed, Ed25519N, Ed25519,
		X25519, X25519_Private,
		Blake3_256, Blake2b_256, Blake2s_256,
		SHA3_256, SHA2_256,
		ECDSA_256k1_Seed, ECDSA_256r1_Seed,
		Salt_256:
		fs = 44
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Ed448_Seed, X448:
		fs = 76
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Short:
		fs = 4
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Big:
		fs = 12
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case X25519_Cipher_Seed:
		fs = 124
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tall:
		fs = 8
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Large:
		fs = 16
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Great:
		fs = 20
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Vast:
		fs = 24
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Label1:
		fs = 4
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 1}, nil

	case Label2:
		fs = 4
		return Sizage{Hs: 1, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag3:
		fs = 4
		return Sizage{Hs: 1, Ss: 3, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag7:
		fs = 8
		return Sizage{Hs: 1, Ss: 7, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag11:
		fs = 12
		return Sizage{Hs: 1, Ss: 11, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Salt_128:
		fs = 24
		return Sizage{Hs: 2, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Ed25519_Sig, ECDSA_256k1_Sig, ECDSA_256r1_Sig,
		Blake3_512, Blake2b_512, SHA3_512, SHA2_512:
		fs = 88
		return Sizage{Hs: 2, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Long:
		fs = 8
		return Sizage{Hs: 2, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag1:
		fs = 4
		return Sizage{Hs: 2, Ss: 2, Xs: 1, Fs: &fs, Ls: 0}, nil

	case Tag2:
		fs = 4
		return Sizage{Hs: 2, Ss: 2, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag5:
		fs = 8
		return Sizage{Hs: 2, Ss: 6, Xs: 1, Fs: &fs, Ls: 0}, nil

	case Tag6:
		fs = 8
		return Sizage{Hs: 2, Ss: 6, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag9:
		fs = 12
		return Sizage{Hs: 2, Ss: 10, Xs: 1, Fs: &fs, Ls: 0}, nil

	case Tag10:
		fs = 12
		return Sizage{Hs: 2, Ss: 10, Xs: 0, Fs: &fs, Ls: 0}, nil

	case GramHeadNeck:
		fs = 32
		return Sizage{Hs: 2, Ss: 22, Xs: 0, Fs: &fs, Ls: 0}, nil

	case GramHead:
		fs = 28
		return Sizage{Hs: 2, Ss: 22, Xs: 0, Fs: &fs, Ls: 0}, nil

	case GramHeadAIDNeck:
		fs = 76
		return Sizage{Hs: 2, Ss: 22, Xs: 0, Fs: &fs, Ls: 0}, nil

	case GramHeadAID:
		fs = 72
		return Sizage{Hs: 2, Ss: 22, Xs: 0, Fs: &fs, Ls: 0}, nil

	case ECDSA_256k1N, ECDSA_256k1, ECDSA_256r1N, ECDSA_256r1:
		fs = 48
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Ed448N, Ed448:
		fs = 80
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Ed448_Sig:
		fs = 156
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag4:
		fs = 8
		return Sizage{Hs: 4, Ss: 4, Xs: 0, Fs: &fs, Ls: 0}, nil

	case DateTime:
		fs = 36
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case X25519_Cipher_Salt:
		fs = 100
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Null, Yes, No, Escape, Empty:
		fs = 4
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case Tag8:
		fs = 12
		return Sizage{Hs: 4, Ss: 8, Xs: 0, Fs: &fs, Ls: 0}, nil

	case TBD0S:
		fs = 12
		return Sizage{Hs: 4, Ss: 2, Xs: 0, Fs: &fs, Ls: 0}, nil

	case TBD0:
		fs = 8
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 0}, nil

	case TBD1S:
		fs = 12
		return Sizage{Hs: 4, Ss: 2, Xs: 1, Fs: &fs, Ls: 1}, nil

	case TBD1:
		fs = 8
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 1}, nil

	case TBD2S:
		fs = 12
		return Sizage{Hs: 4, Ss: 2, Xs: 0, Fs: &fs, Ls: 2}, nil

	case TBD2:
		fs = 8
		return Sizage{Hs: 4, Ss: 0, Xs: 0, Fs: &fs, Ls: 2}, nil

	case StrB64_L0, Bytes_L0, X25519_Cipher_L0, X25519_Cipher_QB64_L0,
		X25519_Cipher_QB2_L0, HPKEBase_Cipher_L0, HPKEAuth_Cipher_L0,
		Decimal_L0:
		return Sizage{Hs: 2, Ss: 2, Xs: 0, Ls: 0}, nil

	case StrB64_L1, Bytes_L1, X25519_Cipher_L1, X25519_Cipher_QB64_L1,
		X25519_Cipher_QB2_L1, HPKEBase_Cipher_L1, HPKEAuth_Cipher_L1,
		Decimal_L1:
		return Sizage{Hs: 2, Ss: 2, Xs: 0, Ls: 1}, nil

	case StrB64_L2, Bytes_L2, X25519_Cipher_L2, X25519_Cipher_QB64_L2,
		X25519_Cipher_QB2_L2, HPKEBase_Cipher_L2, HPKEAuth_Cipher_L2,
		Decimal_L2:
		return Sizage{Hs: 2, Ss: 2, Xs: 0, Ls: 2}, nil

	case StrB64_Big_L0, Bytes_Big_L0, X25519_Cipher_Big_L0, X25519_Cipher_QB64_Big_L0,
		X25519_Cipher_QB2_Big_L0, HPKEBase_Cipher_Big_L0, HPKEAuth_Cipher_Big_L0,
		Decimal_Big_L0:
		return Sizage{Hs: 4, Ss: 4, Xs: 0, Ls: 0}, nil

	case StrB64_Big_L1, Bytes_Big_L1, X25519_Cipher_Big_L1, X25519_Cipher_QB64_Big_L1,
		X25519_Cipher_QB2_Big_L1, HPKEBase_Cipher_Big_L1, HPKEAuth_Cipher_Big_L1,
		Decimal_Big_L1:
		return Sizage{Hs: 4, Ss: 4, Xs: 0, Ls: 1}, nil

	case StrB64_Big_L2, Bytes_Big_L2, X25519_Cipher_Big_L2, X25519_Cipher_QB64_Big_L2,
		X25519_Cipher_QB2_Big_L2, HPKEBase_Cipher_Big_L2, HPKEAuth_Cipher_Big_L2,
		Decimal_Big_L2:
		return Sizage{Hs: 4, Ss: 4, Xs: 0, Ls: 2}, nil

	default:
		return Sizage{}, fmt.Errorf("unknown code: %s", code)
	}
}

func Hardage(c string) (uint32, error) {
	if c >= "A" && c <= "Z" || c >= "a" && c <= "z" {
		return 1, nil
	}
	if c == "0" || c == "4" || c == "5" || c == "6" {
		return 2, nil
	}
	if c == "1" || c == "2" || c == "3" || c == "7" || c == "8" || c == "9" {
		return 4, nil
	}

	if c == "-" {
		return 0, fmt.Errorf("count code start")
	}
	if c == "_" {
		return 0, fmt.Errorf("op code start")
	}

	return 0, fmt.Errorf("unknown hardage: %s", c)
}

func Bardage(b byte) (uint32, error) {
	if b <= 0x33 {
		return 1, nil
	}
	if b == 0x34 || b >= 0x38 && b <= 0x3a {
		return 2, nil
	}
	if b >= 0x35 && b <= 0x37 || b >= 0x3b && b <= 0x3d {
		return 4, nil
	}

	if b == 0x3e {
		return 0, fmt.Errorf("count code start")
	}
	if b == 0x3f {
		return 0, fmt.Errorf("op code start")
	}

	return 0, fmt.Errorf("unknown bardage: %x", b)
}
