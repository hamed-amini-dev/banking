package routes

import (
	"net/http"

	"github.com/ha-dev/banking/internal/handler/accounts"
	lhttp "github.com/ha-dev/banking/internal/http"
)

func AccountsRoutes(th accounts.IAccount) []Route {
	return []Route{
		{
			Name:    "lisAccount",
			Method:  http.MethodGet,
			Path:    "/lisAccount",
			Handler: lhttp.DefaultHTTPHandler(th.ListAccounts),
		},
		/* {
			Name:   "getLocation",
			Method: http.MethodGet,
			Path:   "/utils/map/geocode",
			Handler: lhttp.DefaultHTTPHandler(th.GetLocationByAddress, auth.RegisterAndSessionAuthentication),
		},
		{
			Name:    "getPlace",
			Method:  http.MethodGet,
			Path:    "/utils/map/place",
			Handler: lhttp.DefaultHTTPHandler(th.AutocompleteAddress, auth.SessionAuthentication),
		},
		{
			Name : "check the health of server",
			Method : http.MethodGet,
			Path: "/health-check",
			Handler: lhttp.DefaultHTTPHandler(th.HealthCheck,nil),
		}, */
	}
}
