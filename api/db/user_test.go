package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.chromium.org/gae/service/datastore"
)

func TestSaveAndGetUser(t *testing.T) {
	b := NewBoard(nil, "b")
	assert.NoError(t, b.Save(testCtx))

	u1 := NewUser("uid")
	u1.BoardKeys = []*datastore.Key{b.Key}

	// Save
	assert.Empty(t, u1.Key)
	assert.NoError(t, u1.Save(testCtx))
	assert.NotEmpty(t, u1.Key)

	NewBoard(nil, "board").Save(testCtx)
	// Get
	u2, err := GetUser(testCtx, u1.Key)
	assert.NoError(t, err)
	assert.Equal(t, u1.UID, u2.UID)
	assert.Equal(t, 1, len(u2.BoardKeys))
}

func TestGetBoards(t *testing.T) {
	b1 := NewBoard(nil, "b1")
	b2 := NewBoard(nil, "b2")
	assert.NoError(t, b1.Save(testCtx))
	assert.NoError(t, b2.Save(testCtx))

	u := NewUser("uid")
	u.BoardKeys = []*datastore.Key{b1.Key, b2.Key}
	assert.NoError(t, u.Save(testCtx))

	boards, err := u.GetBoards(testCtx)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(boards))
	assert.Equal(t, "b1", boards[0].Name)
	assert.Equal(t, "b2", boards[1].Name)
}
