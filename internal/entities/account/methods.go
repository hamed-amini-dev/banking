package account

import (
	"fmt"
	"strconv"
)

/*
set account object balance
*/
func (ac *Account) SetBalance(balance string) {
	ac.Balance = balance
}

/*
get account balance
*/
func (ac *Account) GetBalance() string {
	return ac.Balance
}

/*
tools  func for business logic process
*/
func (ac *Account) ConvertBalanceToNumber(balance string) (float64, error) {
	value, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

/*
tools  func for business logic process
*/
func (ac *Account) ConvertBalanceToString(balance float64) string {
	return fmt.Sprintf("%.2f", balance)
}
