package counter

import (
	"github.com/jasoncolburne/cesrgo/types"
	"github.com/jasoncolburne/cesrgo/util"
)

const (
	GenericGroup                  = types.Code("-A")    // Generic Group (Universal with Override).
	BigGenericGroup               = types.Code("--A")   // Big Generic Group (Universal with Override).
	BodyWithAttachmentGroup       = types.Code("-B")    // Message Body plus Attachments Group (Universal with Override).
	BigBodyWithAttachmentGroup    = types.Code("--B")   // Big Message Body plus Attachments Group (Universal with Override).
	AttachmentGroup               = types.Code("-C")    // Message Attachments Only Group (Universal with Override).
	BigAttachmentGroup            = types.Code("--C")   // Big Attachments Only Group (Universal with Override).
	DatagramSegmentGroup          = types.Code("-D")    // Datagram Segment Group (Universal).
	BigDatagramSegmentGroup       = types.Code("--D")   // Big Datagram Segment Group (Universal).
	ESSRWrapperGroup              = types.Code("-E")    // ESSR Wrapper Group (Universal).
	BigESSRWrapperGroup           = types.Code("--E")   // Big ESSR Wrapper Group (Universal).
	FixBodyGroup                  = types.Code("-F")    // Fixed Field Message Body Group (Universal).
	BigFixBodyGroup               = types.Code("--F")   // Big Fixed Field Message Body Group (Universal).
	MapBodyGroup                  = types.Code("-G")    // Field Map Message Body Group (Universal).
	BigMapBodyGroup               = types.Code("--G")   // Big Field Map Message Body Group (Universal).
	NonNativeBodyGroup            = types.Code("-H")    // Message body Non-native enclosed with Texter
	BigNonNativeBodyGroup         = types.Code("--H")   // Big Message body Non-native enclosed with Texter
	GenericMapGroup               = types.Code("-I")    // Generic Field Map Group (Universal).
	BigGenericMapGroup            = types.Code("--I")   // Big Generic Field Map Group (Universal).
	GenericListGroup              = types.Code("-J")    // Generic List Group (Universal).
	BigGenericListGroup           = types.Code("--J")   // Big Generic List Group (Universal).
	ControllerIdxSigs             = types.Code("-K")    // Controller Indexed Signature(s) of qb64.
	BigControllerIdxSigs          = types.Code("--K")   // Big Controller Indexed Signature(s) of qb64.
	WitnessIdxSigs                = types.Code("-L")    // Witness Indexed Signature(s) of qb64.
	BigWitnessIdxSigs             = types.Code("--L")   // Big Witness Indexed Signature(s) of qb64.
	NonTransReceiptCouples        = types.Code("-M")    // NonTrans Receipt Couple(s), pre+cig.
	BigNonTransReceiptCouples     = types.Code("--M")   // Big NonTrans Receipt Couple(s), pre+cig.
	TransReceiptQuadruples        = types.Code("-N")    // Trans Receipt Quadruple(s), pre+snu+dig+sig.
	BigTransReceiptQuadruples     = types.Code("--N")   // Big Trans Receipt Quadruple(s), pre+snu+dig+sig.
	FirstSeenReplayCouples        = types.Code("-O")    // First Seen Replay Couple(s), fnu+dts.
	BigFirstSeenReplayCouples     = types.Code("--O")   // First Seen Replay Couple(s), fnu+dts.
	PathedMaterialGroup           = types.Code("-P")    // Pathed Material Group.
	BigPathedMaterialGroup        = types.Code("--P")   // Big Pathed Material Group.
	DigestSealSingles             = types.Code("-Q")    // Digest Seal Single(s), dig of sealed data.
	BigDigestSealSingles          = types.Code("--Q")   // Big Digest Seal Single(s), dig of sealed data.
	MerkleRootSealSingles         = types.Code("-R")    // Merkle Tree Root Digest Seal Single(s), dig of sealed data.
	BigMerkleRootSealSingles      = types.Code("--R")   // Merkle Tree Root Digest Seal Single(s), dig of sealed data.
	SealSourceTriples             = types.Code("-S")    // Seal Source Triple(s), pre+snu+dig of source sealing or sealed event.
	BigSealSourceTriples          = types.Code("--S")   // Seal Source Triple(s), pre+snu+dig of source sealing or sealed event.
	SealSourceCouples             = types.Code("-T")    // Seal Source Couple(s), snu+dig of source sealing or sealed event.
	BigSealSourceCouples          = types.Code("--T")   // Seal Source Couple(s), snu+dig of source sealing or sealed event.
	SealSourceLastSingles         = types.Code("-U")    // Seal Source Couple(s), pre of last source sealing or sealed event.
	BigSealSourceLastSingles      = types.Code("--U")   // Big Seal Source Couple(s), pre of last source sealing or sealed event.
	BackerRegistrarSealCouples    = types.Code("-V")    // Backer Registrar Seal Couple(s), brid+dig of sealed data.
	BigBackerRegistrarSealCouples = types.Code("--V")   // Big Backer Registrar Seal Couple(s), brid+dig of sealed data.
	TypedDigestSealCouples        = types.Code("-W")    // Typed Digest Seal Couple(s), type seal +dig of sealed data.
	BigTypedDigestSealCouples     = types.Code("--W")   // Big Typed Digest Seal Couple(s), type seal +dig of sealed data.
	TransIdxSigGroups             = types.Code("-X")    // Trans Indexed Signature Group(s), pre+snu+dig+CtrControllerIdxSigs of qb64.
	BigTransIdxSigGroups          = types.Code("--X")   // Big Trans Indexed Signature Group(s), pre+snu+dig+CtrControllerIdxSigs of qb64.
	TransLastIdxSigGroups         = types.Code("-Y")    // Trans Last Est Evt Indexed Signature Group(s), pre+CtrControllerIdxSigs of qb64.
	BigTransLastIdxSigGroups      = types.Code("--Y")   // Big Trans Last Est Evt Indexed Signature Group(s), pre+CtrControllerIdxSigs of qb64.
	ESSRPayloadGroup              = types.Code("-Z")    // ESSR Payload Group.
	BigESSRPayloadGroup           = types.Code("--Z")   // Big ESSR Payload Group.
	BlindedStateQuadruples        = types.Code("-a")    // Blinded transaction event state quadruples dig+uuid+said+state.
	BigBlindedStateQuadruples     = types.Code("--a")   // Big Blinded transaction event state quadruples dig+uuid+said+state.
	KERIACDCGenusVersion          = types.Code("-_AAA") // KERI ACDC Stack CESR Protocol Genus Version (Universal)
)

var UinversalCodex_2_0 = []types.Code{
	GenericGroup,
	BigGenericGroup,
	BodyWithAttachmentGroup,
	BigBodyWithAttachmentGroup,
	AttachmentGroup,
	BigAttachmentGroup,
	DatagramSegmentGroup,
	BigDatagramSegmentGroup,
	ESSRWrapperGroup,
	BigESSRWrapperGroup,
	FixBodyGroup,
	BigFixBodyGroup,
	MapBodyGroup,
	BigMapBodyGroup,
	NonNativeBodyGroup,
	BigNonNativeBodyGroup,
	GenericMapGroup,
	BigGenericMapGroup,
	GenericListGroup,
	BigGenericListGroup,
	KERIACDCGenusVersion,
}

var SpecialUniversalCodex_2_0 = []types.Code{
	GenericGroup,
	BigGenericGroup,
	BodyWithAttachmentGroup,
	BigBodyWithAttachmentGroup,
	AttachmentGroup,
	BigAttachmentGroup,
}

var MessageUniversalCodex_2_0 = []types.Code{
	DatagramSegmentGroup,
	BigDatagramSegmentGroup,
	ESSRWrapperGroup,
	BigESSRWrapperGroup,
	FixBodyGroup,
	BigFixBodyGroup,
	MapBodyGroup,
	BigMapBodyGroup,
	NonNativeBodyGroup,
	BigNonNativeBodyGroup,
}

var SealCodex_2_0 = []types.Code{
	DigestSealSingles,
	BigDigestSealSingles,
	MerkleRootSealSingles,
	BigMerkleRootSealSingles,
	SealSourceTriples,
	BigSealSourceTriples,
	SealSourceCouples,
	BigSealSourceCouples,
	SealSourceLastSingles,
	BigSealSourceLastSingles,
	BackerRegistrarSealCouples,
	BigBackerRegistrarSealCouples,
	TypedDigestSealCouples,
	BigTypedDigestSealCouples,
}

var Hards = map[string]int{}
var Bards = map[string]int{}

func generateHards() {
	if len(Hards) > 0 {
		return
	}

	for i := 65; i < 65+26; i++ {
		Hards["-"+string(byte(i))] = 2
	}

	for i := 97; i < 97+26; i++ {
		Hards["-"+string(byte(i))] = 2
	}

	Hards["--"] = 3
	Hards["-_"] = 5
}

func generateBards() error {
	if len(Bards) > 0 {
		return nil
	}

	generateHards()

	for hard, i := range Hards {
		bard, err := util.CodeB64ToB2(hard)
		if err != nil {
			return err
		}
		Bards[string(bard)] = i
	}

	return nil
}

func Hardage(s string) (int, bool) {
	generateHards()

	n, ok := Hards[s]
	return n, ok
}

func Bardage(b []byte) (int, bool) {
	err := generateBards()
	if err != nil {
		return -1, false
	}

	n, ok := Bards[string(b)]
	return n, ok
}
