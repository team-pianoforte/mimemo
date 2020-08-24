package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.chromium.org/luci/gae/service/datastore"
)

func TestSaveAndGetUser(t *testing.T) {
	b := NewBoard(testCtx, 0, "b")
	assert.NoError(t, b.Save())

	u1 := NewUser(testCtx, "uid")
	u1.BoardKeys = []*datastore.Key{b.Key}

	// Save
	assert.NoError(t, u1.Save())
	assert.NotEmpty(t, u1.Key)
	assert.Equal(t, datastore.NewKey(testCtx, userKind, "uid", 0, nil).Encode(), u1.Key.Encode())

	NewBoard(testCtx, 0, "board").Save()
	// Get
	u2, err := GetUser(testCtx, u1.UID)
	assert.NoError(t, err)
	assert.Equal(t, u1.UID, u2.UID)
	assert.Equal(t, 1, len(u2.BoardKeys))
}

func TestGetBoards(t *testing.T) {
	b1 := NewBoard(testCtx, 0, "b1")
	b2 := NewBoard(testCtx, 0, "b2")
	assert.NoError(t, b1.Save())
	assert.NoError(t, b2.Save())

	u := NewUser(testCtx, "uid")
	u.BoardKeys = []*datastore.Key{b1.Key, b2.Key}
	assert.NoError(t, u.Save())

	boards, err := u.GetBoards()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(boards))
	assert.Equal(t, "b1", boards[0].Name)
	assert.Equal(t, "b2", boards[1].Name)
}
