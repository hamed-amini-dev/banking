package accounts

import (
	sAccount "github.com/ha-dev/banking/services/accounts"
)

type iAccount struct {
	service sAccount.IAccount
}

func New(ops ...Option) (IAccount, error) {
	h := new(iAccount)
	for _, fn := range ops {
		err := fn(h)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}
