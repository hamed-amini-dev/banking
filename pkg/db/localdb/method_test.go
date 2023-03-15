package localdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	f := OptionLocalMockFile("../../../accounts-mock.json")
	db, err := New(f)
	assert.Nil(t, err)
	listAccounts := db.List()
	assert.Equal(t, len(accounts), len(listAccounts))
}

func TestGet(t *testing.T) {
	f := OptionLocalMockFile("../../../accounts-mock.json")
	db, err := New(f)
	assert.Nil(t, err)
	fetchAccount, err := db.Get("3d253e29-8785-464f-8fa0-9e4b57699db9")
	assert.Nil(t, err)
	assert.Equal(t, "3d253e29-8785-464f-8fa0-9e4b57699db9", fetchAccount.ID)
}

func TestUpdate(t *testing.T) {
	f := OptionLocalMockFile("../../../accounts-mock.json")
	db, err := New(f)
	assert.Nil(t, err)
	fetchAccount, err := db.Get("3d253e29-8785-464f-8fa0-9e4b57699db9")
	fetchAccount.Balance = "15"
	err = db.Update(fetchAccount)
	assert.Nil(t, err)
	assert.Equal(t, "15", fetchAccount.Balance)
}
