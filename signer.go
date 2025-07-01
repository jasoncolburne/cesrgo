package cesrgo

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/crypto"
	idex "github.com/jasoncolburne/cesrgo/indexer"
	iopts "github.com/jasoncolburne/cesrgo/indexer/options"
	mdex "github.com/jasoncolburne/cesrgo/matter"
	mopts "github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Signer struct {
	matter
	verfer *Verfer
}

func (s *Signer) GetVerfer() *Verfer {
	return s.verfer
}

var validSignerCodes []types.Code = []types.Code{
	mdex.Ed25519_Seed,
	mdex.ECDSA_256k1_Seed,
	mdex.ECDSA_256r1_Seed,
}

func NewSigner(transferable bool, opts ...mopts.MatterOption) (*Signer, error) {
	s := &Signer{}

	config := &mopts.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		var code types.Code
		if config.Code == nil {
			code = mdex.Ed25519_Seed
		} else {
			code = *config.Code
		}

		if !validateCode(code, validSignerCodes) {
			return nil, fmt.Errorf("unexpected code: %s", code)
		}

		opts = []mopts.MatterOption{}
		if config.Raw == nil {
			raw, err := crypto.GenerateSeed(code)
			if err != nil {
				return nil, err
			}

			opts = append(opts, mopts.WithRaw(raw))
		} else {
			opts = append(opts, mopts.WithRaw(*config.Raw))
		}

		opts = append(opts, mopts.WithCode(code))
	}

	if config.Qb2 != nil {
		if config.Code != nil || config.Raw != nil || config.Qb64 != nil || config.Qb64b != nil {
			return nil, fmt.Errorf("qb2 cannot be used with code, raw, qb64, or qb64b")
		}
	}

	if config.Qb64 != nil {
		if config.Code != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64b != nil {
			return nil, fmt.Errorf("qb64 cannot be used with code, raw, qb2, or qb64b")
		}
	}

	if config.Qb64b != nil {
		if config.Code != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64 != nil {
			return nil, fmt.Errorf("qb64b cannot be used with code, raw, qb2, or qb64")
		}
	}

	if err := NewMatter(s, opts...); err != nil {
		return nil, err
	}

	if !validateCode(s.code, validSignerCodes) {
		return nil, fmt.Errorf("unexpected code: %s", *config.Code)
	}

	if err := s.deriveVerfer(transferable); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Signer) deriveVerfer(transferable bool) error {
	verferCode, verferRaw, err := crypto.DeriveCodeAndPublicKey(s.code, s.raw, transferable)
	if err != nil {
		return err
	}

	if s.verfer, err = NewVerfer(
		mopts.WithCode(verferCode),
		mopts.WithRaw(verferRaw),
	); err != nil {
		return err
	}

	return nil
}

func (s *Signer) SignUnindexed(ser []byte) (*Cigar, error) {
	if !validateCode(s.code, validSignerCodes) {
		return nil, fmt.Errorf("unexpected code: %s", s.code)
	}

	var (
		code types.Code
		raw  types.Raw
		err  error
	)

	switch s.code {
	case mdex.Ed25519_Seed:
		code = mdex.Ed25519_Sig
	case mdex.ECDSA_256k1_Seed:
		code = mdex.ECDSA_256k1_Sig
	case mdex.ECDSA_256r1_Seed:
		code = mdex.ECDSA_256r1_Sig
	default:
		return nil, fmt.Errorf("unexpected code: %s", s.code)
	}

	raw, err = crypto.Sign(s.code, s.raw, ser)
	if err != nil {
		return nil, err
	}

	c, err := NewCigar(s.verfer, mopts.WithCode(code), mopts.WithRaw(raw))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Signer) SignIndexed(
	ser []byte,
	only bool,
	index types.Index,
	ondex *types.Ondex,
) (*Siger, error) {
	if !validateCode(s.code, validSignerCodes) {
		return nil, fmt.Errorf("unexpected code: %s", s.code)
	}

	var (
		code types.Code
		raw  types.Raw
		err  error
	)

	if only {
		if index < 64 {
			switch s.code {
			case mdex.Ed25519_Seed:
				code = idex.Ed25519_Crt
			case mdex.ECDSA_256k1_Seed:
				code = idex.ECDSA_256k1_Crt
			case mdex.ECDSA_256r1_Seed:
				code = idex.ECDSA_256r1_Crt
			default:
				return nil, fmt.Errorf("unexpected code: %s", s.code)
			}
		} else {
			switch s.code {
			case mdex.Ed25519_Seed:
				code = idex.Ed25519_Big_Crt
			case mdex.ECDSA_256k1_Seed:
				code = idex.ECDSA_256k1_Big_Crt
			case mdex.ECDSA_256r1_Seed:
				code = idex.ECDSA_256r1_Big_Crt
			default:
				return nil, fmt.Errorf("unexpected code: %s", s.code)
			}
		}
	} else {
		var odx types.Ondex
		if ondex == nil {
			odx = types.Ondex(index)
		} else {
			odx = *ondex
		}

		if uint32(odx) == uint32(index) && index < 64 {
			switch s.code {
			case mdex.Ed25519_Seed:
				code = idex.Ed25519
			case mdex.ECDSA_256k1_Seed:
				code = idex.ECDSA_256k1
			case mdex.ECDSA_256r1_Seed:
				code = idex.ECDSA_256r1
			default:
				return nil, fmt.Errorf("unexpected code: %s", s.code)
			}
		} else {
			switch s.code {
			case mdex.Ed25519_Seed:
				code = idex.Ed25519_Big
			case mdex.ECDSA_256k1_Seed:
				code = idex.ECDSA_256k1_Big
			case mdex.ECDSA_256r1_Seed:
				code = idex.ECDSA_256r1_Big
			default:
				return nil, fmt.Errorf("unexpected code: %s", s.code)
			}
		}
	}

	raw, err = crypto.Sign(s.code, s.raw, ser)
	if err != nil {
		return nil, err
	}

	siger, err := NewSiger(s.verfer, iopts.WithCode(code), iopts.WithRaw(raw))
	if err != nil {
		return nil, err
	}

	return siger, nil
}
