package account

/*
define model functions
we need database module for using data item record

*/

import (
	"github.com/ha-dev/banking/pkg/db/localdb"
)

type iAccount struct {
	db localdb.ILocalDB
}

var _ IAccount = &iAccount{}

//go:generate mockgen -destination=../../testutils/mocks/model/iAccount_mock.go -package=mocks -source=interface.go

// create object model for handling database operation
// need option for getting database provider module
// return model object for using database record functionality

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
