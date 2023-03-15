package accounts

import "github.com/ha-dev/banking/internal/request"

//define sign interface functions
type IAccount interface {
	//Get list account
	ListAccounts(r *request.GenericRequest) (interface{}, error)
	//Get account with account id from route url
	GetAccount(r *request.GenericRequest) (interface{}, error)
	//Transfer Balance From Account To Account
	Transfer(r *request.GenericRequest) (interface{}, error)
}
