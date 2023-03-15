package accounts

import (
	mAccount "github.com/ha-dev/banking/models/account"
)

type Option func(acc *iAccount) error

/*
InitOptionAccountModel for initialize model account business logic
*/
func InitOptionAccountModel(model mAccount.IAccount) Option {
	return func(acc *iAccount) error {
		acc.model = model
		return nil
	}
}
