package accounts

import (
	"encoding/json"

	"github.com/ha-dev/banking/internal/request"
	sAccount "github.com/ha-dev/banking/services/accounts"
)

// Getting list of account

func (h *iAccount) ListAccounts(r *request.GenericRequest) (interface{}, error) {
	return h.service.ListAll()
}

// Get account with
//   - account-id

func (h *iAccount) GetAccount(r *request.GenericRequest) (interface{}, error) {
	ID := r.GetPathParameter(accountID)
	return h.service.Get(ID)
}

// Transfer Blance from account to another account
//   - from
//   - to
//   - balance

func (h *iAccount) Transfer(r *request.GenericRequest) (interface{}, error) {
	var params sAccount.TransferParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}
	return h.service.Transfer(&params)
}
