package accounts

import (
	mAccount "github.com/ha-dev/banking/models/account"
)

type iAccount struct {
	model mAccount.IAccount
}

var _ IAccount = &iAccount{}

//go:generate mockgen -destination=../../testutils/mocks/service/iAccount_mock.go -package=mocks -source=interface.go

/*
create object service for handling service account operation
need option for getting model account provider module
return service object for using  logic account functionality
*/

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
