package accounts

import (
	"encoding/json"

	"github.com/ha-dev/banking/internal/request"
	sAccount "github.com/ha-dev/banking/services/accounts"
)

// ────────────────────────────────────────────────────────────────────────────────
func (h *iAccount) ListAccounts(r *request.GenericRequest) (interface{}, error) {
	return h.service.ListAll()
}

// ────────────────────────────────────────────────────────────────────────────────
func (h *iAccount) GetAccount(r *request.GenericRequest) (interface{}, error) {
	ID := r.GetPathParameter(accountID)
	return h.service.Get(ID)
}

// ────────────────────────────────────────────────────────────────────────────────
func (h *iAccount) Transfer(r *request.GenericRequest) (interface{}, error) {
	var params sAccount.TransferParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}

	return h.service.Transfer(&params)
}
