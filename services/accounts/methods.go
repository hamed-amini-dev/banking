package accounts

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

/*
  List all the account that exist in model
  - result array of account from model logic
*/

func (c *iAccount) ListAll() ([]*eAccount.Account, error) {
	return c.model.List()
}

/*
 Get account record from model logic process
 - ID string
*/

func (c *iAccount) Get(ID string) (*eAccount.Account, error) {
	return c.model.Get(ID)
}

/*
- Transfer Balance From account To Account
- Before Transfer Validate From Balance
*/
func (c *iAccount) Transfer(param *TransferParams) (*eAccount.Account, error) {
	fromAccount, err := c.model.Get(param.From)
	if err != nil {
		return nil, err
	}
	//validate transfer
	ok, err := c.validateTransfer(fromAccount, param.Balance)
	if err != nil {
		return nil, err
	}
	//if ok to transfer
	//else inform user about transfer
	if ok {
		toAccount, err := c.model.Get(param.To)
		if err != nil {
			return nil, err
		}

		fromBalance, err := fromAccount.ConvertBalanceToNumber(fromAccount.GetBalance())
		if err != nil {
			return nil, err
		}

		toBalance, err := toAccount.ConvertBalanceToNumber(toAccount.GetBalance())
		if err != nil {
			return nil, err
		}

		transferBalance, err := fromAccount.ConvertBalanceToNumber(param.Balance)
		if err != nil {
			return nil, err
		}

		finalBalanceFromAccount := fromBalance - transferBalance
		finalBalanceToAccount := toBalance + transferBalance

		//
		fromAccount.SetBalance(fromAccount.ConvertBalanceToString(finalBalanceFromAccount))
		toAccount.SetBalance(toAccount.ConvertBalanceToString(finalBalanceToAccount))

		//update db
		err = c.model.Update(fromAccount)
		if err != nil {
			return nil, err
		}

		err = c.model.Update(toAccount)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	return nil, ErrEnoughBalance
}

/*
- validateTransfer Check Transfer from account with balance is validate or not
- Check have user enough balance to transfer
*/
func (c *iAccount) validateTransfer(fromAcc *eAccount.Account, balance string) (bool, error) {

	accBalance, err := fromAcc.ConvertBalanceToNumber(fromAcc.GetBalance())
	if err != nil {
		return false, err
	}

	transferBalance, err := fromAcc.ConvertBalanceToNumber(balance)
	if err != nil {
		return false, err
	}

	if accBalance-transferBalance < 0 {
		return false, nil
	} else {
		return true, nil
	}
}
