package matter

import (
	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/core/types"
)

const (
	Ed25519_Seed       = types.Code("A") // Ed25519 256 bit random seed for private key
	Ed25519N           = types.Code("B") // Ed25519 verification key non-transferable, basic derivation.
	X25519             = types.Code("C") // X25519 public encryption key, may be converted from Ed25519 or Ed25519N.
	Ed25519            = types.Code("D") // Ed25519 verification key basic derivation
	Blake3_256         = types.Code("E") // Blake3 256 bit digest self-addressing derivation.
	Blake2b_256        = types.Code("F") // Blake2b 256 bit digest self-addressing derivation.
	Blake2s_256        = types.Code("G") // Blake2s 256 bit digest self-addressing derivation.
	SHA3_256           = types.Code("H") // SHA3 256 bit digest self-addressing derivation.
	SHA2_256           = types.Code("I") // SHA2 256 bit digest self-addressing derivation.
	ECDSA_256k1_Seed   = types.Code("J") // ECDSA secp256k1 256 bit random Seed for private key
	Ed448_Seed         = types.Code("K") // Ed448 448 bit random Seed for private key
	X448               = types.Code("L") // X448 public encryption key, converted from Ed448
	Short              = types.Code("M") // Short 2 byte b2 number
	Big                = types.Code("N") // Big 8 byte b2 number
	X25519_Private     = types.Code("O") // X25519 private decryption key/seed, may be converted from Ed25519
	X25519_Cipher_Seed = types.Code("P") // X25519 sealed box 124 char qb64 Cipher of 44 char qb64 Seed
	ECDSA_256r1_Seed   = types.Code("Q") // ECDSA secp256r1 256 bit random Seed for private key
	Tall               = types.Code("R") // Tall 5 byte b2 number
	Large              = types.Code("S") // Large 11 byte b2 number
	Great              = types.Code("T") // Great 14 byte b2 number
	Vast               = types.Code("U") // Vast 17 byte b2 number
	Label1             = types.Code("V") // Label1 1 bytes for label lead size 1
	Label2             = types.Code("W") // Label2 2 bytes for label lead size 0
	Tag3               = types.Code("X") // Tag3  3 B64 encoded chars for special values
	Tag7               = types.Code("Y") // Tag7  7 B64 encoded chars for special values
	Tag11              = types.Code("Z") // Tag11  11 B64 encoded chars for special values
	Salt_256           = types.Code("a") // Salt/seed/nonce/blind 256 bits
	AES_256            = types.Code("b") // AES 256 bit key

	// this code is identified by two names
	Salt_128 = types.Code("0A") // Salt/seed/nonce 128 bits or number of length 128 bits (Huge)
	Huge     = types.Code("0A") // This is huge!

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

var BextCodex = []types.Code{
	StrB64_L0,
	StrB64_L1,
	StrB64_L2,
	StrB64_Big_L0,
	StrB64_Big_L1,
	StrB64_Big_L2,
}

var TextCodex = []types.Code{
	Bytes_L0,
	Bytes_L1,
	Bytes_L2,
	Bytes_Big_L0,
	Bytes_Big_L1,
	Bytes_Big_L2,
}

var DecimalCodex = []types.Code{
	Decimal_L0,
	Decimal_L1,
	Decimal_L2,
	Decimal_Big_L0,
	Decimal_Big_L1,
	Decimal_Big_L2,
}

var DigCodex = []types.Code{
	Blake3_256,
	Blake2b_256,
	Blake2s_256,
	SHA3_256,
	SHA2_256,
	Blake3_512,
	Blake2b_512,
	SHA3_512,
	SHA2_512,
}

var NonceCodex = []types.Code{
	Empty,
	Salt_128,
	Salt_256,
	Blake3_256,
	Blake2b_256,
	Blake2s_256,
	SHA3_256,
	SHA2_256,
	Blake3_512,
	Blake2b_512,
	SHA3_512,
	SHA2_512,
}

var NumCodex = []types.Code{
	Short,
	Long,
	Tall,
	Big,
	Large,
	Great,
	Huge,
	Vast,
}

var TagCodex = []types.Code{
	Tag1,
	Tag2,
	Tag3,
	Tag4,
	Tag5,
	Tag6,
	Tag7,
	Tag8,
	Tag9,
	Tag10,
	Tag11,
}

var LabelCodex = []types.Code{
	Empty,
	Tag1,
	Tag2,
	Tag3,
	Tag4,
	Tag5,
	Tag6,
	Tag7,
	Tag8,
	Tag9,
	Tag10,
	Tag11,
	StrB64_L0,
	StrB64_L1,
	StrB64_L2,
	StrB64_Big_L0,
	StrB64_Big_L1,
	StrB64_Big_L2,
	Label1,
	Label2,
	Bytes_L0,
	Bytes_L1,
	Bytes_L2,
	Bytes_Big_L0,
	Bytes_Big_L1,
	Bytes_Big_L2,
}

var PreCodex = []types.Code{
	Ed25519N,
	Ed25519,
	Blake3_256,
	Blake2b_256,
	Blake2s_256,
	SHA3_256,
	SHA2_256,
	Blake3_512,
	Blake2b_512,
	SHA3_512,
	SHA2_512,
	ECDSA_256k1N,
	ECDSA_256k1,
	Ed448N,
	Ed448,
	Ed448_Sig,
	ECDSA_256r1N,
	ECDSA_256r1,
}

var NonTransCodex = []types.Code{
	Ed25519N,
	ECDSA_256k1N,
	Ed448N,
	ECDSA_256r1N,
}

var PreNonDigCodex = []types.Code{
	Ed25519N,
	Ed25519,
	ECDSA_256k1N,
	ECDSA_256k1,
	Ed448N,
	Ed448,
	ECDSA_256r1N,
	ECDSA_256r1,
}

var SigCodex = []types.Code{
	Ed25519_Sig,
	ECDSA_256k1_Sig,
	ECDSA_256r1_Sig,
	Ed448_Sig,
}

var SeedCodex = []types.Code{
	Ed25519_Seed,
	ECDSA_256k1_Seed,
	ECDSA_256r1_Seed,
	Ed448_Seed,
}

var SMALL_VRZ_DEX = []rune{'4', '5', '6'}
var LARGE_VRZ_DEX = []rune{'7', '8', '9'}
var SMALL_VRZ_BYTES = uint32(3)
var LARGE_VRZ_BYTES = uint32(6)

type Sizage struct {
	Hs uint32
	Ss uint32
	Xs uint32
	Fs *uint32
	Ls uint32
}

var (
	_4   = uint32(4)
	_8   = uint32(8)
	_12  = uint32(12)
	_16  = uint32(16)
	_20  = uint32(20)
	_24  = uint32(24)
	_28  = uint32(28)
	_32  = uint32(32)
	_36  = uint32(36)
	_44  = uint32(44)
	_48  = uint32(48)
	_72  = uint32(72)
	_76  = uint32(76)
	_80  = uint32(80)
	_88  = uint32(88)
	_100 = uint32(100)
	_124 = uint32(124)
	_156 = uint32(156)
)

var Sizes = map[types.Code]Sizage{
	// keys & seeds

	Ed25519_Seed:     {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	Ed25519N:         {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	Ed25519:          {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	X25519:           {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	X25519_Private:   {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	ECDSA_256k1_Seed: {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	ECDSA_256r1_Seed: {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},

	ECDSA_256k1N: {Hs: 4, Ss: 0, Xs: 0, Fs: &_48, Ls: 0},
	ECDSA_256k1:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_48, Ls: 0},
	ECDSA_256r1N: {Hs: 4, Ss: 0, Xs: 0, Fs: &_48, Ls: 0},
	ECDSA_256r1:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_48, Ls: 0},

	Ed448N: {Hs: 4, Ss: 0, Xs: 0, Fs: &_80, Ls: 0},
	Ed448:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_80, Ls: 0},

	Ed448_Seed: {Hs: 1, Ss: 0, Xs: 0, Fs: &_76, Ls: 0},
	X448:       {Hs: 1, Ss: 0, Xs: 0, Fs: &_76, Ls: 0},

	X25519_Cipher_Seed: {Hs: 1, Ss: 0, Xs: 0, Fs: &_124, Ls: 0},

	// signatures

	Ed25519_Sig:     {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},
	ECDSA_256k1_Sig: {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},
	ECDSA_256r1_Sig: {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},

	Ed448_Sig: {Hs: 4, Ss: 0, Xs: 0, Fs: &_156, Ls: 0},

	// digests

	Blake3_256:  {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	Blake2b_256: {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	Blake2s_256: {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	SHA3_256:    {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	SHA2_256:    {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},

	Blake3_512:  {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},
	Blake2b_512: {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},
	SHA3_512:    {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},
	SHA2_512:    {Hs: 2, Ss: 0, Xs: 0, Fs: &_88, Ls: 0},

	// salts

	Salt_128:           {Hs: 2, Ss: 0, Xs: 0, Fs: &_24, Ls: 0},
	Salt_256:           {Hs: 1, Ss: 0, Xs: 0, Fs: &_44, Ls: 0},
	X25519_Cipher_Salt: {Hs: 4, Ss: 0, Xs: 0, Fs: &_100, Ls: 0},

	// sizes

	Short: {Hs: 1, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},
	Tall:  {Hs: 1, Ss: 0, Xs: 0, Fs: &_8, Ls: 0},
	Big:   {Hs: 1, Ss: 0, Xs: 0, Fs: &_12, Ls: 0},
	Large: {Hs: 1, Ss: 0, Xs: 0, Fs: &_16, Ls: 0},
	Great: {Hs: 1, Ss: 0, Xs: 0, Fs: &_20, Ls: 0},
	Vast:  {Hs: 1, Ss: 0, Xs: 0, Fs: &_24, Ls: 0},

	Long: {Hs: 2, Ss: 0, Xs: 0, Fs: &_8, Ls: 0},

	// labels

	Label1: {Hs: 1, Ss: 0, Xs: 0, Fs: &_4, Ls: 1},
	Label2: {Hs: 1, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},

	// tags

	Tag1:  {Hs: 2, Ss: 2, Xs: 1, Fs: &_4, Ls: 0},
	Tag2:  {Hs: 2, Ss: 2, Xs: 0, Fs: &_4, Ls: 0},
	Tag3:  {Hs: 1, Ss: 3, Xs: 0, Fs: &_4, Ls: 0},
	Tag4:  {Hs: 4, Ss: 4, Xs: 0, Fs: &_8, Ls: 0},
	Tag5:  {Hs: 2, Ss: 6, Xs: 1, Fs: &_8, Ls: 0},
	Tag6:  {Hs: 2, Ss: 6, Xs: 0, Fs: &_8, Ls: 0},
	Tag7:  {Hs: 1, Ss: 7, Xs: 0, Fs: &_8, Ls: 0},
	Tag8:  {Hs: 4, Ss: 8, Xs: 0, Fs: &_12, Ls: 0},
	Tag9:  {Hs: 2, Ss: 10, Xs: 1, Fs: &_12, Ls: 0},
	Tag10: {Hs: 2, Ss: 10, Xs: 0, Fs: &_12, Ls: 0},
	Tag11: {Hs: 1, Ss: 11, Xs: 0, Fs: &_12, Ls: 0},

	// more

	Null:   {Hs: 4, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},
	Yes:    {Hs: 4, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},
	No:     {Hs: 4, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},
	Escape: {Hs: 4, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},
	Empty:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_4, Ls: 0},

	// gram head

	GramHead:        {Hs: 2, Ss: 22, Xs: 0, Fs: &_28, Ls: 0},
	GramHeadNeck:    {Hs: 2, Ss: 22, Xs: 0, Fs: &_32, Ls: 0},
	GramHeadAID:     {Hs: 2, Ss: 22, Xs: 0, Fs: &_72, Ls: 0},
	GramHeadAIDNeck: {Hs: 2, Ss: 22, Xs: 0, Fs: &_76, Ls: 0},

	// dates

	DateTime: {Hs: 4, Ss: 0, Xs: 0, Fs: &_36, Ls: 0},

	// TBD

	TBD0S: {Hs: 4, Ss: 2, Xs: 0, Fs: &_12, Ls: 0},
	TBD0:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_8, Ls: 0},
	TBD1S: {Hs: 4, Ss: 2, Xs: 1, Fs: &_12, Ls: 1},
	TBD1:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_8, Ls: 1},
	TBD2S: {Hs: 4, Ss: 2, Xs: 0, Fs: &_12, Ls: 2},
	TBD2:  {Hs: 4, Ss: 0, Xs: 0, Fs: &_8, Ls: 2},

	// variable length data

	StrB64_L0:             {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	Bytes_L0:              {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	X25519_Cipher_L0:      {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	X25519_Cipher_QB64_L0: {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	X25519_Cipher_QB2_L0:  {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	HPKEBase_Cipher_L0:    {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	HPKEAuth_Cipher_L0:    {Hs: 2, Ss: 2, Xs: 0, Ls: 0},
	Decimal_L0:            {Hs: 2, Ss: 2, Xs: 0, Ls: 0},

	StrB64_L1:             {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	Bytes_L1:              {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	X25519_Cipher_L1:      {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	X25519_Cipher_QB64_L1: {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	X25519_Cipher_QB2_L1:  {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	HPKEBase_Cipher_L1:    {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	HPKEAuth_Cipher_L1:    {Hs: 2, Ss: 2, Xs: 0, Ls: 1},
	Decimal_L1:            {Hs: 2, Ss: 2, Xs: 0, Ls: 1},

	StrB64_L2:             {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	Bytes_L2:              {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	X25519_Cipher_L2:      {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	X25519_Cipher_QB64_L2: {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	X25519_Cipher_QB2_L2:  {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	HPKEBase_Cipher_L2:    {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	HPKEAuth_Cipher_L2:    {Hs: 2, Ss: 2, Xs: 0, Ls: 2},
	Decimal_L2:            {Hs: 2, Ss: 2, Xs: 0, Ls: 2},

	StrB64_Big_L0:             {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	Bytes_Big_L0:              {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	X25519_Cipher_Big_L0:      {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	X25519_Cipher_QB64_Big_L0: {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	X25519_Cipher_QB2_Big_L0:  {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	HPKEBase_Cipher_Big_L0:    {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	HPKEAuth_Cipher_Big_L0:    {Hs: 4, Ss: 4, Xs: 0, Ls: 0},
	Decimal_Big_L0:            {Hs: 4, Ss: 4, Xs: 0, Ls: 0},

	StrB64_Big_L1:             {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	Bytes_Big_L1:              {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	X25519_Cipher_Big_L1:      {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	X25519_Cipher_QB64_Big_L1: {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	X25519_Cipher_QB2_Big_L1:  {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	HPKEBase_Cipher_Big_L1:    {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	HPKEAuth_Cipher_Big_L1:    {Hs: 4, Ss: 4, Xs: 0, Ls: 1},
	Decimal_Big_L1:            {Hs: 4, Ss: 4, Xs: 0, Ls: 1},

	StrB64_Big_L2:             {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	Bytes_Big_L2:              {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	X25519_Cipher_Big_L2:      {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	X25519_Cipher_QB64_Big_L2: {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	X25519_Cipher_QB2_Big_L2:  {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	HPKEBase_Cipher_Big_L2:    {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	HPKEAuth_Cipher_Big_L2:    {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
	Decimal_Big_L2:            {Hs: 4, Ss: 4, Xs: 0, Ls: 2},
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

	Hards[byte('0')] = 2
	Hards[byte('4')] = 2
	Hards[byte('5')] = 2
	Hards[byte('6')] = 2

	Hards[byte('1')] = 4
	Hards[byte('2')] = 4
	Hards[byte('3')] = 4
	Hards[byte('7')] = 4
	Hards[byte('8')] = 4
	Hards[byte('9')] = 4
}

func generateBards() error {
	if len(Bards) > 0 {
		return nil
	}

	generateHards()

	for hard, i := range Hards {
		bard, err := common.B64CharToIndex(rune(hard))
		if err != nil {
			return err
		}

		Bards[bard] = i
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
