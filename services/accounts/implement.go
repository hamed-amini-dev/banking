package accounts

import (
	mAccount "github.com/ha-dev/banking/models/account"
)

type iAccount struct {
	model mAccount.IAccount
}

var _ IAccount = &iAccount{}

func New(ops ...Option) (IAccount, error) {
	s := new(iAccount)
	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}
