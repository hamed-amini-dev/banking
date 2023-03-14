package account

import (
	"fmt"
	"strconv"
)

// ─────────────────────────────────────────────────────────────────────────────
func (ac *Account) SetBalance(balance string) {
	ac.Balance = balance
}

// ─────────────────────────────────────────────────────────────────────────────
func (ac *Account) GetBalance() string {
	return ac.Balance
}

// ─────────────────────────────────────────────────────────────────────────────
func (ac *Account) ConvertBalanceToNumber(balance string) (float64, error) {
	value, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (ac *Account) ConvertBalanceToString(balance float64) string {
	return fmt.Sprintf("%.2f", balance)
}
