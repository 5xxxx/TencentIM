package TencentIM

import (
	oldsig "github.com/tencentyun/tls-sig-api-golang"
)

type SigVersion int

const (
	VER_NEW SigVersion = 0
	VER_OLD SigVersion = 1
)

type ServerOption interface {
	SetOption(*IMServer) error
}

type SignOption struct {
	sigVersion SigVersion
	privateKey string
}

func NewSignOption() *SignOption {
	return &SignOption{}
}

func (o *SignOption) SetSigVersion(ver SigVersion) *SignOption {
	o.sigVersion = ver
	return o
}

func (o *SignOption) SetPrivateKey(privateKey string) *SignOption {
	o.privateKey = privateKey
	return o
}

func (o *SignOption) SetOption(s *IMServer) error {
	switch o.sigVersion {
	case VER_OLD:
		if o.privateKey == "" {
			return ERR_INVALID_PRI_KEY
		}
		sig, err := oldsig.GenerateUsersigWithExpire(o.privateKey, s.AppId, s.Identifier, int64(s.Expire))
		if err != nil {
			return err
		}
		s.Sig = sig
	case VER_NEW:
		var err error
		if s.Sig, err = s.userSig(); err != nil {
			return err
		}
	default:
		return ERR_INVALID_VER
	}

	return nil
}
