package app

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	hAccount "github.com/ha-dev/banking/internal/handler/accounts"
	"github.com/ha-dev/banking/internal/routes"
	mAccount "github.com/ha-dev/banking/models/account"
	"github.com/ha-dev/banking/pkg/constants"
	"github.com/ha-dev/banking/pkg/db/localdb"
	sAccount "github.com/ha-dev/banking/services/accounts"
	"github.com/spf13/viper"
)

type app struct {
	router *mux.Router
}

var (
	onceInitConfig = sync.Once{}
)

// NewApp - Loads the configuration file for setting up the server
// and returns the app server
func NewApp() (server *http.Server, err error) {

	// ─── DATABASE INITIALIZATION ────────────────────────────────────────────────────
	//init datastore

	db, err := localdb.New(localdb.OptionLocalMockFile("accounts-mock.json"))
	if err != nil {
		return nil, err
	}

	// ─────────────────────────────────────────────────────── APP INITIALIZATION ─────
	newApp := new(app)
	//Model
	modelAccount, err := mAccount.New(mAccount.OptionDB(db))
	if err != nil {
		return nil, err
	}
	//Services
	serviceAccount, err := sAccount.New(sAccount.InitOptionAccountModel(modelAccount))
	if err != nil {
		return nil, err
	}
	//Handler
	handlerAccount, err := hAccount.New(hAccount.InitOptionService(serviceAccount))
	if err != nil {
		return nil, err
	}
	//all router handle
	var allRoutes []routes.Route

	AccountRoutes := routes.AccountsRoutes(handlerAccount)
	allRoutes = append(allRoutes, AccountRoutes...)

	// create resource
	err = newApp.createResources(allRoutes...)
	if err != nil {
		return nil, err
	}

	return newApp.server(), nil
}

// server- Returns pointer to an instance of type http.Server
func (a *app) server() *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString(constants.Port)),
		Handler: a.router,
	}

	log.Println("listen on port: %v", server.Addr)

	return server
}

// createResources - Creates router instance and registers handler to the router
func (a *app) createResources(rs ...routes.Route) error {
	a.router = mux.NewRouter().StrictSlash(true)
	// a.router.Handle("/auth/warm-up", http.FileServer(http.Dir("/www/ohs/")))

	for _, r := range rs {
		err := a.router.
			Name(r.Name).
			Path(r.Path).
			Methods(r.Method, http.MethodOptions).
			HandlerFunc(r.Handler).GetError()
		if err != nil {
			return err
		}
	}

	return nil
}
