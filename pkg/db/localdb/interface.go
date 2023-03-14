package localdb

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type ILocalDB interface {
	// ─── TABALES OPERATION ──────────────────────────────────────────────────────────
	//Get One Account Base on Filed Name
	Get(FieldName string) (*eAccount.Account, error)

	//Set On Account Base on Filed Name
	Update(acc *eAccount.Account) error

	//List Accounts
	List() []*eAccount.Account
}
