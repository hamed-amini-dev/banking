package accounts

import (
	"fmt"

	"github.com/ha-dev/banking/internal/request"
)

// ────────────────────────────────────────────────────────────────────────────────
func (h *iAccount) ListAccounts(r *request.GenericRequest) (interface{}, error) {
	fmt.Println("test handler")
	return nil, nil
}
