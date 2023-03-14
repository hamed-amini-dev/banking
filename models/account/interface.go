package account

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type IAccount interface {
	//List all the account that exist in DB
	List() ([]*eAccount.Account, error)
}
