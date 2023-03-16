package account

import eAccount "github.com/ha-dev/banking/internal/entities/account"

//   List all the account that exist in DB
//   - result array of account from database

func (s *iAccount) List() ([]*eAccount.Account, error) {
	return s.db.List(), nil
}

//  Get account record from database
//  - ID string

func (s *iAccount) Get(ID string) (*eAccount.Account, error) {
	return s.db.Get(ID)
}

//   Update the account that exist in DB
//   - account struct
//     -- ID
// 	-- Name
// 	-- Balance

func (s *iAccount) Update(acc *eAccount.Account) error {
	return s.db.Update(acc)
}
