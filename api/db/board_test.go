package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
)

var testCtx context.Context

func TestMain(m *testing.M) {
	testCtx = memory.Use(context.Background())
	testable := datastore.GetTestable(testCtx)
	testable.AutoIndex(true)

	os.Exit(m.Run())
}

func TestSaveAndGetBoard(t *testing.T) {
	b1 := NewBoard(testCtx, 0, "name")
	b1.Memos = append(b1.Memos, *NewMemo("memo"))

	// Save
	assert.Empty(t, b1.Key)
	assert.NoError(t, b1.Save(testCtx))
	assert.NotEmpty(t, b1.Key)

	// Get
	b2, err := GetBoard(testCtx, b1.Key.IntID())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(b2.Memos))
	assert.Equal(t, "memo", b2.Memos[0].Text)
	assert.Equal(t, b1.Name, b2.Name)

	// And and Update Memo
	b2.Memos = append([]Memo{*NewMemo("memo0")}, b2.Memos...)
	b2.Memos[1].Text = "memo1"
	assert.NoError(t, b2.Save(testCtx))

	b3, err := GetBoard(testCtx, b1.Key.IntID())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(b3.Memos))
	assert.Equal(t, "memo0", b3.Memos[0].Text)
	assert.Equal(t, "memo1", b3.Memos[1].Text)
	assert.Equal(t, b1.Name, b3.Name)
}
