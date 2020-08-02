package db

import (
	"context"
	"testing"

	"github.com/luci/gae/impl/memory"
	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	ctx := memory.Use(context.Background())

	b := NewBoard("name", "")

	assert.Empty(t, b.ID)
	assert.NoError(t, b.Save(ctx))
	assert.NotEmpty(t, b.ID)

	newB, err := GetBoard(ctx, b.ID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, b.Name, newB.Name)
}
