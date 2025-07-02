package two

import (
	counter "github.com/jasoncolburne/cesrgo/counter/sizage"
	"github.com/jasoncolburne/cesrgo/types"
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

var CounterCodex = []types.Code{
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
	ControllerIdxSigs,
	BigControllerIdxSigs,
	WitnessIdxSigs,
	BigWitnessIdxSigs,
	NonTransReceiptCouples,
	BigNonTransReceiptCouples,
	TransReceiptQuadruples,
	BigTransReceiptQuadruples,
	FirstSeenReplayCouples,
	BigFirstSeenReplayCouples,
	PathedMaterialGroup,
	BigPathedMaterialGroup,
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
	TransIdxSigGroups,
	BigTransIdxSigGroups,
	TransLastIdxSigGroups,
	BigTransLastIdxSigGroups,
	ESSRPayloadGroup,
	BigESSRPayloadGroup,
	BlindedStateQuadruples,
	BigBlindedStateQuadruples,
	KERIACDCGenusVersion,
}

var UniversalCodex = []types.Code{
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

var SpecialUniversalCodex = []types.Code{
	GenericGroup,
	BigGenericGroup,
	BodyWithAttachmentGroup,
	BigBodyWithAttachmentGroup,
	AttachmentGroup,
	BigAttachmentGroup,
}

var MessageUniversalCodex = []types.Code{
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

var SealCodex = []types.Code{
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

var Sizes = map[types.Code]counter.Sizage{
	GenericGroup:                  {Hs: 2, Ss: 2, Fs: 4},
	BigGenericGroup:               {Hs: 3, Ss: 5, Fs: 8},
	BodyWithAttachmentGroup:       {Hs: 2, Ss: 2, Fs: 4},
	BigBodyWithAttachmentGroup:    {Hs: 3, Ss: 5, Fs: 8},
	AttachmentGroup:               {Hs: 2, Ss: 2, Fs: 4},
	BigAttachmentGroup:            {Hs: 3, Ss: 5, Fs: 8},
	DatagramSegmentGroup:          {Hs: 2, Ss: 2, Fs: 4},
	BigDatagramSegmentGroup:       {Hs: 3, Ss: 5, Fs: 8},
	ESSRWrapperGroup:              {Hs: 2, Ss: 2, Fs: 4},
	BigESSRWrapperGroup:           {Hs: 3, Ss: 5, Fs: 8},
	FixBodyGroup:                  {Hs: 2, Ss: 2, Fs: 4},
	BigFixBodyGroup:               {Hs: 3, Ss: 5, Fs: 8},
	MapBodyGroup:                  {Hs: 2, Ss: 2, Fs: 4},
	BigMapBodyGroup:               {Hs: 3, Ss: 5, Fs: 8},
	NonNativeBodyGroup:            {Hs: 2, Ss: 2, Fs: 4},
	BigNonNativeBodyGroup:         {Hs: 3, Ss: 5, Fs: 8},
	GenericMapGroup:               {Hs: 2, Ss: 2, Fs: 4},
	BigGenericMapGroup:            {Hs: 3, Ss: 5, Fs: 8},
	GenericListGroup:              {Hs: 2, Ss: 2, Fs: 4},
	BigGenericListGroup:           {Hs: 3, Ss: 5, Fs: 8},
	ControllerIdxSigs:             {Hs: 2, Ss: 2, Fs: 4},
	BigControllerIdxSigs:          {Hs: 3, Ss: 5, Fs: 8},
	WitnessIdxSigs:                {Hs: 2, Ss: 2, Fs: 4},
	BigWitnessIdxSigs:             {Hs: 3, Ss: 5, Fs: 8},
	NonTransReceiptCouples:        {Hs: 2, Ss: 2, Fs: 4},
	BigNonTransReceiptCouples:     {Hs: 3, Ss: 5, Fs: 8},
	TransReceiptQuadruples:        {Hs: 2, Ss: 2, Fs: 4},
	BigTransReceiptQuadruples:     {Hs: 3, Ss: 5, Fs: 8},
	FirstSeenReplayCouples:        {Hs: 2, Ss: 2, Fs: 4},
	BigFirstSeenReplayCouples:     {Hs: 3, Ss: 5, Fs: 8},
	PathedMaterialGroup:           {Hs: 2, Ss: 2, Fs: 4},
	BigPathedMaterialGroup:        {Hs: 3, Ss: 5, Fs: 8},
	DigestSealSingles:             {Hs: 2, Ss: 2, Fs: 4},
	BigDigestSealSingles:          {Hs: 3, Ss: 5, Fs: 8},
	MerkleRootSealSingles:         {Hs: 2, Ss: 2, Fs: 4},
	BigMerkleRootSealSingles:      {Hs: 3, Ss: 5, Fs: 8},
	SealSourceTriples:             {Hs: 2, Ss: 2, Fs: 4},
	BigSealSourceTriples:          {Hs: 3, Ss: 5, Fs: 8},
	SealSourceCouples:             {Hs: 2, Ss: 2, Fs: 4},
	BigSealSourceCouples:          {Hs: 3, Ss: 5, Fs: 8},
	SealSourceLastSingles:         {Hs: 2, Ss: 2, Fs: 4},
	BigSealSourceLastSingles:      {Hs: 3, Ss: 5, Fs: 8},
	BackerRegistrarSealCouples:    {Hs: 2, Ss: 2, Fs: 4},
	BigBackerRegistrarSealCouples: {Hs: 3, Ss: 5, Fs: 8},
	TypedDigestSealCouples:        {Hs: 2, Ss: 2, Fs: 4},
	BigTypedDigestSealCouples:     {Hs: 3, Ss: 5, Fs: 8},
	TransIdxSigGroups:             {Hs: 2, Ss: 2, Fs: 4},
	BigTransIdxSigGroups:          {Hs: 3, Ss: 5, Fs: 8},
	TransLastIdxSigGroups:         {Hs: 2, Ss: 2, Fs: 4},
	BigTransLastIdxSigGroups:      {Hs: 3, Ss: 5, Fs: 8},
	ESSRPayloadGroup:              {Hs: 2, Ss: 2, Fs: 4},
	BigESSRPayloadGroup:           {Hs: 3, Ss: 5, Fs: 8},
	BlindedStateQuadruples:        {Hs: 2, Ss: 2, Fs: 4},
	BigBlindedStateQuadruples:     {Hs: 3, Ss: 5, Fs: 8},
	KERIACDCGenusVersion:          {Hs: 5, Ss: 3, Fs: 8},
}
