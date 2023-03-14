package localdb

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

// ────────────────────────────────────────────────────────────────────────────────
func (u *localdbConfig) Get(FieldName string) (*eAccount.Account, error) {
	for _, v := range accounts {
		if v.ID == FieldName || v.Balance == FieldName || v.Name == FieldName {
			return v, nil
		}
	}

	return nil, ErrNotFound
}

// ────────────────────────────────────────────────────────────────────────────────
func (u *localdbConfig) List() []*eAccount.Account {
	return accounts
}

// ────────────────────────────────────────────────────────────────────────────────
func (u *localdbConfig) Update(acc *eAccount.Account) error {
	for i := range accounts {
		if accounts[i].ID == acc.ID {
			accounts[i].Name = acc.Name
			accounts[i].Balance = acc.Balance
			return nil
		}
	}
	return ErrNotFound
}
