package accounts

/*
define endpoint functions
we need service module for using business logic of endpoint
*/
import (
	sAccount "github.com/ha-dev/banking/services/accounts"
)

type iAccount struct {
	service sAccount.IAccount
}

// create object handler for handling route
// need option for getting service module
// return handler object for using logic functionality

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
