package validator

import (
	"github.com/namelessup/bilibili/app/service/main/tv/internal/pkg"
	"github.com/namelessup/bilibili/library/ecode"
)

type SignerValidator struct {
	Signer *pkg.Signer
	Sign   string
	Val    interface{}
}

func (sv *SignerValidator) Validate() error {
	sign, err := sv.Signer.Sign(sv.Val)
	if err != nil {
		return err
	}
	if sign != sv.Sign {
		return ecode.TVIPSignErr
	}
	return nil
}
