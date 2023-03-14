package localdb

import (
	"encoding/json"
	"io/ioutil"

	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

type localdbConfig struct {
	path string
}

var accounts []*eAccount.Account

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
	return s, nil
}

// ────────────────────────────────────────────────────────────────────────────────

func (u *localdbConfig) Get(FieldName string) (*eAccount.Account, error) {
	return nil, nil
}

func (u *localdbConfig) List() []*eAccount.Account {
	return accounts
}
