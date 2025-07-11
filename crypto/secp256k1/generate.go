package secp256k1

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/jasoncolburne/cesrgo/core/types"
)

const SEED_BYTES = 32

func GenerateSeed() (types.Raw, error) {
	priv, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}

	return types.Raw(priv.Serialize()), nil
}

func DerivePublicKey(seed types.Raw) (types.Raw, error) {
	priv := secp256k1.PrivKeyFromBytes(seed)
	pub := priv.PubKey()

	return types.Raw(pub.SerializeCompressed()), nil
}
