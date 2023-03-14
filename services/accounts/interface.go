package accounts

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type IAccount interface {
	//List all the account which are in DB
	ListAll() ([]*eAccount.Account, error)
}
