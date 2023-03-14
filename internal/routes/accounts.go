package routes

import (
	"net/http"

	"github.com/ha-dev/banking/internal/handler/accounts"
	lhttp "github.com/ha-dev/banking/internal/http"
)

func AccountsRoutes(th accounts.IAccount) []Route {
	return []Route{
		{
			Name:    "listOfAccounts",
			Method:  http.MethodGet,
			Path:    "/accounts",
			Handler: lhttp.DefaultHTTPHandler(th.ListAccounts),
		},
		{
			Name:    "getAccount",
			Method:  http.MethodGet,
			Path:    "/accounts/{account-id}",
			Handler: lhttp.DefaultHTTPHandler(th.GetAccount),
		},
		{
			Name:    "transferBalance",
			Method:  http.MethodPost,
			Path:    "/accounts",
			Handler: lhttp.DefaultHTTPHandler(th.Transfer),
		},
	}
}
