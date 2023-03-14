package accounts

import (
	sAccount "github.com/ha-dev/banking/services/accounts"
)

type Option func(*iAccount) error

func InitOptionService(service sAccount.IAccount) Option {
	return func(is *iAccount) error {
		is.service = service
		return nil
	}
}
