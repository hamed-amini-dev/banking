package account

import (
	"testing"

	"github.com/golang/mock/gomock"
	eAccount "github.com/ha-dev/banking/internal/entities/account"
	"github.com/ha-dev/banking/pkg/db/localdb"
	mocks "github.com/ha-dev/banking/testutils/mocks/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockLocalRepo := mocks.NewMockILocalDB(c)
	t.Run("gets data", func(t *testing.T) {
		mockLocalRepo.EXPECT().List().Return([]*eAccount.Account{
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
		})

		//
		service := iAccount{
			db: mockLocalRepo,
		}
		list, err := service.List()
		assert.Nil(t, err)
		assert.Equal(t, "3d253e29-8785-464f-8fa0-9e4b57699db9", list[0].ID)
		assert.Equal(t, "17f904c1-806f-4252-9103-74e7a5d3e340", list[1].ID)

	})
}

func TestGet(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockLocalRepo := mocks.NewMockILocalDB(c)
	account := &eAccount.Account{
		ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Name:    "Trupe",
		Balance: "87.11",
	}
	t.Run("get user data is ok", func(t *testing.T) {
		mockLocalRepo.EXPECT().Get("3d253e29-8785-464f-8fa0-9e4b57699db9").Return(account, nil)
		//
		service := iAccount{
			db: mockLocalRepo,
		}
		fetchAccount, err := service.Get("3d253e29-8785-464f-8fa0-9e4b57699db9")
		assert.Nil(t, err)
		assert.Equal(t, "3d253e29-8785-464f-8fa0-9e4b57699db9", fetchAccount.ID)

	})

	t.Run("get user not find", func(t *testing.T) {
		mockLocalRepo.EXPECT().Get(gomock.Any()).Return(nil, localdb.ErrNotFound)
		//
		service := iAccount{
			db: mockLocalRepo,
		}
		_, err := service.Get("3d253e29-8785-464f-8fa0-9e4b57699db777")
		assert.ErrorIs(t, localdb.ErrNotFound, err)
	})

}

func TestUpdate(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockLocalRepo := mocks.NewMockILocalDB(c)
	account := &eAccount.Account{
		ID:      "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Name:    "Trupe",
		Balance: "87.11",
	}
	t.Run("update success", func(t *testing.T) {
		mockLocalRepo.EXPECT().Update(account).Return(nil)
		service := iAccount{
			db: mockLocalRepo,
		}
		err := service.Update(account)
		assert.Nil(t, err)
	})

	t.Run("update fail", func(t *testing.T) {
		mockLocalRepo.EXPECT().Update(account).Return(localdb.ErrNotFound)
		service := iAccount{
			db: mockLocalRepo,
		}
		err := service.Update(account)
		assert.ErrorIs(t, localdb.ErrNotFound, err)
	})

}

func TestOptionDB(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockLocalRepo := mocks.NewMockILocalDB(c)

	service := &iAccount{
		db: mockLocalRepo,
	}

	f := OptionDB(mockLocalRepo)
	err := f(service)
	assert.Nil(t, err)
}

func TestNew(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	mockLocalRepo := mocks.NewMockILocalDB(c)

	f := OptionDB(mockLocalRepo)
	_, err := New(f)
	assert.Nil(t, err)
}
