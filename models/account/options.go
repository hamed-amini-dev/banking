package account

import (
	"github.com/ha-dev/banking/pkg/db/localdb"
)

type Option func(*iAccount) error

func OptionDB(db localdb.ILocalDB) Option {
	return func(i *iAccount) error {
		i.db = db
		return nil
	}
}
