package accounts

import (
	"testing"

	"github.com/golang/mock/gomock"
	eAccount "github.com/ha-dev/banking/internal/entities/account"
	"github.com/ha-dev/banking/pkg/db/localdb"
	mocks "github.com/ha-dev/banking/testutils/mocks/model"
	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountRepo := mocks.NewMockIAccount(c)
	t.Run("same data", func(t *testing.T) {
		mockAccountRepo.EXPECT().List().Return([]*eAccount.Account{
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

		//
		service := iAccount{
			model: mockAccountRepo,
		}
		list, err := service.ListAll()
		assert.Nil(t, err)
		assert.Equal(t, "3d253e29-8785-464f-8fa0-9e4b57699db9", list[0].ID)
		assert.Equal(t, "17f904c1-806f-4252-9103-74e7a5d3e340", list[1].ID)

	})

	t.Run("list empty", func(t *testing.T) {
		mockAccountRepo.EXPECT().List().Return([]*eAccount.Account{}, nil)

		//
		service := iAccount{
			model: mockAccountRepo,
		}
		list, err := service.ListAll()
		assert.Nil(t, err)
		assert.Equal(t, 0, len(list))

	})

}

func TestGet(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountRepo := mocks.NewMockIAccount(c)
	account := &eAccount.Account{
		ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Name:    "Trupe",
		Balance: "87.11",
	}
	t.Run("get user data is ok", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account, nil)
		//
		service := iAccount{
			model: mockAccountRepo,
		}
		fetchAccount, err := service.Get("3d253e29-8785-464f-8fa0-9e4b57699db9")
		assert.Nil(t, err)
		assert.Equal(t, "3d253e29-8785-464f-8fa0-9e4b57699db9", fetchAccount.ID)

	})

	t.Run("get user not find", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get(gomock.Any()).Return(nil, localdb.ErrNotFound)
		//
		service := iAccount{
			model: mockAccountRepo,
		}
		_, err := service.Get("3d253e29-8785-464f-8fa0-9e4b57699db777")
		// assert.ErrorType(t, err)
		// assert.ErrorType(t, localdb.ErrNotFound, err)
		assert.ErrorIs(t, localdb.ErrNotFound, err)
	})

}

func TestTransfer(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountRepo := mocks.NewMockIAccount(c)
	account1 := &eAccount.Account{
		ID:      "17f904c1-806f-4252-9103-74e7a5d3e340",
		Name:    "Fivespan",
		Balance: "946.15",
	}

	account2 := &eAccount.Account{
		ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Name:    "Trupe",
		Balance: "87.11",
	}

	t.Run("transfer is ok", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)
		mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account2, nil)
		mockAccountRepo.EXPECT().Update(gomock.Any()).Return(nil)
		mockAccountRepo.EXPECT().Update(gomock.Any()).Return(nil)
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "40",
		}
		_, err := service.Transfer(params)
		assert.Nil(t, err)
		assert.Equal(t, "127.11", account2.Balance)

	})

	t.Run("transfer update faild", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)
		mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account2, nil)
		mockAccountRepo.EXPECT().Update(gomock.Any()).Return(localdb.ErrNotFound)
		// mockAccountRepo.EXPECT().Update(gomock.Any()).Return(nil)
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "40",
		}
		_, err := service.Transfer(params)
		assert.Error(t, err)

	})

	t.Run("transfer update faild 2", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)
		mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account2, nil)
		mockAccountRepo.EXPECT().Update(gomock.Any()).Return(nil)
		mockAccountRepo.EXPECT().Update(gomock.Any()).Return(localdb.ErrNotFound)
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "40",
		}
		_, err := service.Transfer(params)
		assert.Error(t, err)

	})

	t.Run("transfer not valid", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)

		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "1000",
		}
		_, err := service.Transfer(params)
		// assert.Nil(t, err)
		assert.Equal(t, ErrEnoughBalance, err)

	})

	t.Run("transfer faild , to account balance err", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)
		acc2 := &eAccount.Account{
			ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Name:    "Trupe",
			Balance: "aa",
		}
		mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(acc2, nil)
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "40",
		}
		_, err := service.Transfer(params)
		//assert.Nil(t, err)
		assert.Error(t, err)
		// assert.Equal(t, ErrEnoughBalance, err)

	})

	t.Run("transfer faild , balance param err", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)
		// mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account2, nil)
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "jh",
		}
		_, err := service.Transfer(params)
		//assert.Nil(t, err)
		assert.Error(t, err)
		// assert.Equal(t, ErrEnoughBalance, err)

	})

	t.Run("transfer faild , from account balance err", func(t *testing.T) {
		acc1 := &eAccount.Account{
			ID:      "17f904c1-806f-4252-9103-74e7a5d3e340",
			Name:    "Fivespan",
			Balance: "aa",
		}
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(acc1, nil)
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "40",
		}
		_, err := service.Transfer(params)
		//assert.Nil(t, err)
		assert.Error(t, err)
		// assert.Equal(t, ErrEnoughBalance, err)

	})

	t.Run("transfer user not find", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get(gomock.Any()).Return(nil, localdb.ErrNotFound)

		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
			Balance: "1000",
		}
		_, err := service.Transfer(params)
		assert.Error(t, err)
		// assert.Equal(t, localdb.ErrNotFound, err)

	})

	t.Run("transfer user not find", func(t *testing.T) {
		mockAccountRepo.EXPECT().Get("17f904c1-806f-4252-9103-74e7a5d3e340").Return(account1, nil)
		mockAccountRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db910").Return(nil, localdb.ErrNotFound)

		//
		service := iAccount{
			model: mockAccountRepo,
		}

		//
		params := &TransferParams{
			From:    "17f904c1-806f-4252-9103-74e7a5d3e340",
			To:      "3d253e29-8785-464f-8fa0-9e4b57699db910",
			Balance: "40",
		}
		_, err := service.Transfer(params)
		assert.Error(t, err)
		// assert.Equal(t, localdb.ErrNotFound, err)

	})

}

func TestValidateTransfer(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountRepo := mocks.NewMockIAccount(c)
	fromAccount := &eAccount.Account{
		ID:      "17f904c1-806f-4252-9103-74e7a5d3e340",
		Name:    "Fivespan",
		Balance: "946.15",
	}
	t.Run("validate is true", func(t *testing.T) {
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		result, err := service.validateTransfer(fromAccount, "40")
		assert.Nil(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("validate is false", func(t *testing.T) {
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		result, err := service.validateTransfer(fromAccount, "1000")
		assert.Nil(t, err)
		assert.Equal(t, 3, result)
	})

	t.Run("balance not valid", func(t *testing.T) {
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		result, _ := service.validateTransfer(fromAccount, "aa")
		assert.Equal(t, 1, result)
	})

	t.Run("from acc balance not valid", func(t *testing.T) {
		//
		service := iAccount{
			model: mockAccountRepo,
		}

		fromAccount.Balance = "aaa"
		result, _ := service.validateTransfer(fromAccount, "40")
		assert.Equal(t, 1, result)
	})
}

func TestOptionDB(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountRepo := mocks.NewMockIAccount(c)

	service := &iAccount{
		model: mockAccountRepo,
	}

	f := InitOptionAccountModel(mockAccountRepo)
	err := f(service)
	assert.Nil(t, err)
}

func TestNew(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockAccountRepo := mocks.NewMockIAccount(c)

	f := InitOptionAccountModel(mockAccountRepo)
	_, err := New(f)
	assert.Nil(t, err)
}
