package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account := Account{}
	account.SetBalance("47")
	balance := account.GetBalance()
	assert.Equal(t, "47", balance)
	convertBalance, err := account.ConvertBalanceToNumber(balance)
	assert.Nil(t, err)
	assert.Equal(t, 47.00, convertBalance)
	strBalance := account.ConvertBalanceToString(convertBalance)
	assert.Equal(t, "47.00", strBalance)

}
