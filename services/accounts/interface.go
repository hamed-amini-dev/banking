package accounts

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type IAccount interface {
	//List all the account which are in DB
	ListAll() ([]*eAccount.Account, error)

	//Get the account which are in DB
	Get(ID string) (*eAccount.Account, error)

	//Transfer Balance From account To Account
	//Before Transfer Validate From Balance
	Transfer(param *TransferParams) (*eAccount.Account, error)
}
