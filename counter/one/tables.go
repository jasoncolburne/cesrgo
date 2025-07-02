package one

import "github.com/jasoncolburne/cesrgo/types"

const (
	ControllerIdxSigs          = types.Code("-A")    // Qualified Base64 Indexed Signature.
	WitnessIdxSigs             = types.Code("-B")    // Qualified Base64 Indexed Signature.
	NonTransReceiptCouples     = types.Code("-C")    // Composed Base64 Couple, pre+cig.
	TransReceiptQuadruples     = types.Code("-D")    // Composed Base64 Quadruple, pre+snu+dig+sig.
	FirstSeenReplayCouples     = types.Code("-E")    // Composed Base64 Couple, fnu+dts.
	TransIdxSigGroups          = types.Code("-F")    // Composed Base64 Group, pre+snu+dig+ControllerIdxSigs group.
	SealSourceCouples          = types.Code("-G")    // Composed Base64 couple, snu+dig of given delegator/issuer/transaction event
	TransLastIdxSigGroups      = types.Code("-H")    // Composed Base64 Group, pre+ControllerIdxSigs group.
	SealSourceTriples          = types.Code("-I")    // Composed Base64 triple, pre+snu+dig of anchoring source event
	PathedMaterialGroup        = types.Code("-L")    // Composed Grouped Pathed Material Quadlet (4 char each)
	BigPathedMaterialGroup     = types.Code("--L")   // Composed Grouped Pathed Material Quadlet (4 char each)
	GenericGroup               = types.Code("-T")    // Generic Material Quadlet (Universal with override)
	BigGenericGroup            = types.Code("--T")   // Big Generic Material Quadlet (Universal with override)
	BodyWithAttachmentGroup    = types.Code("-U")    // Message Body plus Attachments Quadlet (Universal with Override).
	BigBodyWithAttachmentGroup = types.Code("--U")   // Big Message Body plus Attachments Quadlet (Universal with Override)
	AttachmentGroup            = types.Code("-V")    // Message Attachments Only Quadlet (Universal with Override)
	BigAttachmentGroup         = types.Code("--V")   // Message Attachments Only Quadlet (Universal with Override)
	NonNativeBodyGroup         = types.Code("-W")    // Message body Non-native enclosed with Texter
	BigNonNativeBodyGroup      = types.Code("--W")   // Big Message body Non-native enclosed with Texter
	ESSRPayloadGroup           = types.Code("-Z")    // ESSR Payload Group Quadlets (not implemented as quadlets)
	BigESSRPayloadGroup        = types.Code("--Z")   // Big ESSR Payload Group Quadlets (not implemented as quadlets)
	KERIACDCGenusVersion       = types.Code("-_AAA") // KERI ACDC Protocol Stack CESR Version
)

var CounterCodex = []types.Code{
	ControllerIdxSigs,
	WitnessIdxSigs,
	NonTransReceiptCouples,
	TransReceiptQuadruples,
	FirstSeenReplayCouples,
	TransIdxSigGroups,
	SealSourceCouples,
	TransLastIdxSigGroups,
	SealSourceTriples,
	PathedMaterialGroup,
	BigPathedMaterialGroup,
	GenericGroup,
	BigGenericGroup,
	BodyWithAttachmentGroup,
	BigBodyWithAttachmentGroup,
	AttachmentGroup,
	BigAttachmentGroup,
	NonNativeBodyGroup,
	BigNonNativeBodyGroup,
	ESSRPayloadGroup,
	BigESSRPayloadGroup,
	KERIACDCGenusVersion,
}

var QuadTripCodex = []types.Code{
	PathedMaterialGroup,
	BigPathedMaterialGroup,
	GenericGroup,
	BigGenericGroup,
	BodyWithAttachmentGroup,
	BigBodyWithAttachmentGroup,
	AttachmentGroup,
	BigAttachmentGroup,
	NonNativeBodyGroup,
	BigNonNativeBodyGroup,
	ESSRPayloadGroup,
	BigESSRPayloadGroup,
}

var UniversalCodex = []types.Code{
	GenericGroup,
	BigGenericGroup,
	BodyWithAttachmentGroup,
	BigBodyWithAttachmentGroup,
	AttachmentGroup,
	BigAttachmentGroup,
	NonNativeBodyGroup,
	BigNonNativeBodyGroup,
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
	NonNativeBodyGroup,
	BigNonNativeBodyGroup,
}
