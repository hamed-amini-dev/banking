package account

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
	"github.com/ha-dev/banking/pkg/db/localdb"
)

type iAccount struct {
	db localdb.ILocalDB
}

var _ IAccount = &iAccount{}

func New(ops ...Option) (IAccount, error) {
	s := new(iAccount)

	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ────────────────────────────────────────────────────────────────────────────────
// List all the room that exist in DB
func (s *iAccount) List() ([]*eAccount.Account, error) {
	return s.db.List(), nil
}
