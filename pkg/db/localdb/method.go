package localdb

import eAccount "github.com/ha-dev/banking/internal/entities/account"

/*
 Get account record from in-memory database
 - ID string
*/
func (u *localdbConfig) Get(FieldName string) (*eAccount.Account, error) {
	for _, v := range accounts {
		if v.ID == FieldName || v.Balance == FieldName || v.Name == FieldName {
			return v, nil
		}
	}

	return nil, ErrNotFound
}

/*
  List all the account that exist in DB
  - result array of account from in-memory database
*/

func (u *localdbConfig) List() []*eAccount.Account {
	return accounts
}

/*
  Update the account that exist in DB
  - account struct
    -- ID
	-- Name
	-- Balance
  return error if account not exist on in-memory database
*/

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
