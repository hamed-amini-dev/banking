package localdb

/*
database engine module for crud standad operation on database
this module create standard function for using database provider
*/
import (
	"encoding/json"
	"io/ioutil"
	"log"

	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type localdbConfig struct {
	path string
}

var accounts []*eAccount.Account

//go:generate mockgen -destination=../../../testutils/mocks/pkg/db/localdb_mock.go -package=mocks -source=interface.go

// create object for provide database storage functionality
// need option for getting database mock file
// return  object for using database record functionality

func New(ops ...Option) (ILocalDB, error) {
	s := new(localdbConfig)
	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}
	//Reading json mock file
	file, _ := ioutil.ReadFile(s.path)
	err := json.Unmarshal([]byte(file), &accounts)
	if err != nil {
		return nil, err
	}
	//
	log.Println("Load Accounts Complete")
	//
	return s, nil
}
