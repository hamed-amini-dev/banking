package account

import (
	"github.com/ha-dev/banking/pkg/db/localdb"
)

type Option func(*iAccount) error

// Init option databse for getting database provider

func OptionDB(db localdb.ILocalDB) Option {
	return func(i *iAccount) error {
		i.db = db
		return nil
	}
}
