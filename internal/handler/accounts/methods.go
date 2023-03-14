package accounts

import (
	"github.com/ha-dev/banking/internal/request"
)

// ────────────────────────────────────────────────────────────────────────────────
func (h *iAccount) ListAccounts(r *request.GenericRequest) (interface{}, error) {
	return h.service.ListAll()
}
