package account

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type IAccount interface {
	//List all the account that exist in DB
	List() ([]*eAccount.Account, error)

	//Get the account that exist in DB
	Get(ID string) (*eAccount.Account, error)

	//Update the account that exist in DB
	Update(acc *eAccount.Account) error
}
