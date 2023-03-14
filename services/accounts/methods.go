package accounts

import (
	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

func (c *iAccount) ListAll() ([]*eAccount.Account, error) {
	return c.model.List()
}
