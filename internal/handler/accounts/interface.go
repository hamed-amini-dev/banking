package accounts

import "github.com/ha-dev/banking/internal/request"

type IAccount interface {
	ListAccounts(r *request.GenericRequest) (interface{}, error)
}
