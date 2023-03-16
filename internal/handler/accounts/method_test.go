package accounts

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	lhttp "github.com/ha-dev/banking/internal/http"

	// "github.com/ha-dev/banking/internal/response"

	// "github.com/ha-dev/banking/internal/response"

	mocks "github.com/ha-dev/banking/testutils/mocks/service"
	"github.com/stretchr/testify/assert"

	eAccount "github.com/ha-dev/banking/internal/entities/account"
)

func TestListAccounts(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountService := mocks.NewMockIAccount(c)
	t.Run("get list account", func(t *testing.T) {
		mockAccountService.EXPECT().ListAll().Return([]*eAccount.Account{
			{
				ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
				Name:    "Trupe",
				Balance: "87.11",
			},
			{
				ID:      "17f904c1-806f-4252-9103-74e7a5d3e340",
				Name:    "Fivespan",
				Balance: "946.15",
			},
		}, nil)
		serviceHandler := iAccount{
			service: mockAccountService,
		}

		resp := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/accounts", nil)

		router := mux.NewRouter().StrictSlash(true)
		err := router.Name("listOfAccounts").Path("/accounts").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(lhttp.DefaultHTTPHandler(serviceHandler.ListAccounts)).GetError()
		assert.Nil(t, err)
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
		resultResponse := `{"success":true,"messages":null,"data":[{"id":"3d253e29-8785-464f-8fa0-9e4b57699db9","name":"Trupe","balance":"87.11"},{"id":"17f904c1-806f-4252-9103-74e7a5d3e340","name":"Fivespan","balance":"946.15"}],"meta":{}}`
		test := strings.Split(string(resp.Body.Bytes()), "\n")
		assert.Equal(t, resultResponse, test[0])

	})
}

func TestGetAccounts(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountService := mocks.NewMockIAccount(c)
	account := &eAccount.Account{
		ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Name:    "Trupe",
		Balance: "87.11",
	}
	t.Run("get account", func(t *testing.T) {
		serviceHandler := iAccount{
			service: mockAccountService,
		}
		resp := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/accounts/3d253e29-8785-464f-8fa0-9e4b57699db9", nil)
		mockAccountService.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account, nil)
		//
		router := mux.NewRouter().StrictSlash(true)
		err := router.Name("getAccount").Path("/accounts/{account-id}").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(lhttp.DefaultHTTPHandler(serviceHandler.GetAccount)).GetError()
		assert.Nil(t, err)
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
		// log.Println(rr.Body)

	})
}

func TestTransfer(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountService := mocks.NewMockIAccount(c)
	account2 := &eAccount.Account{
		ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Name:    "Trupe",
		Balance: "127.11",
	}
	t.Run("get account", func(t *testing.T) {
		serviceHandler := iAccount{
			service: mockAccountService,
		}
		resp := httptest.NewRecorder()
		var jsonData = []byte(`{
			"from": "17f904c1-806f-4252-9103-74e7a5d3e340",
			"to": "3d253e29-8785-464f-8fa0-9e4b57699db9",
			"balance":"40"
		}`)
		req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonData))
		mockAccountService.EXPECT().Transfer(gomock.Any()).Return(account2, nil)
		//
		router := mux.NewRouter().StrictSlash(true)
		err := router.Name("transferBalance").Path("/accounts").Methods(http.MethodPost, http.MethodOptions).HandlerFunc(lhttp.DefaultHTTPHandler(serviceHandler.Transfer)).GetError()
		assert.Nil(t, err)
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "127.11", account2.Balance)
	})
}

func TestOptionDB(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountService := mocks.NewMockIAccount(c)

	serviceHandler := &iAccount{
		service: mockAccountService,
	}

	f := InitOptionService(mockAccountService)
	err := f(serviceHandler)
	assert.Nil(t, err)
}

func TestNew(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountService := mocks.NewMockIAccount(c)

	f := InitOptionService(mockAccountService)
	_, err := New(f)
	assert.Nil(t, err)
}
