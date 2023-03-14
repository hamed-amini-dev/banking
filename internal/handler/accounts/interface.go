package accounts

import "github.com/ha-dev/banking/internal/request"

type IAccount interface {
	ListAccounts(r *request.GenericRequest) (interface{}, error)
	GetAccount(r *request.GenericRequest) (interface{}, error)

	//Transfer Balance From Account To Account
	Transfer(r *request.GenericRequest) (interface{}, error)
}
