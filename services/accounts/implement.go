package accounts

import (
	mAccount "github.com/ha-dev/banking/models/account"
)

type iAccount struct {
	model mAccount.IAccount
}

var _ IAccount = &iAccount{}

//go:generate mockgen -destination=../../testutils/mocks/service/iAccount_mock.go -package=mocks -source=interface.go
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
