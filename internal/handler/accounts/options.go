package accounts

import (
	sAccount "github.com/ha-dev/banking/services/accounts"
)

type Option func(*iAccount) error

/*
Init option service for getting service business logic account layer
*/
func InitOptionService(service sAccount.IAccount) Option {
	return func(is *iAccount) error {
		is.service = service
		return nil
	}
}
