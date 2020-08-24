package db

import (
	"context"

	"go.chromium.org/luci/gae/service/datastore"
)

const (
	boardKind = "Board"
)

type Board struct {
	Key   *datastore.Key `gae:"$key"`
	Name  string
	Index int
	Memos []Memo
	ctx   context.Context `gae:"-"`
}

type Memo struct {
	Text  string
	Index int
}

func NewBoard(ctx context.Context, id int64, name string) (b *Board) {
	b = &Board{
		Name: name,
		ctx:  ctx,
	}
	if id == 0 {
		b.Key = nil
	} else {
		b.Key = datastore.NewKey(ctx, boardKind, "", id, nil)
	}
	return
}

func NewMemo(text string) *Memo {
	return &Memo{Text: text}
}

func GetBoard(ctx context.Context, id int64) (*Board, error) {
	b := NewBoard(ctx, id, "")
	if err := datastore.Get(ctx, b); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Board) Save() error {
	if b.Key == nil {
		b.Key = datastore.NewIncompleteKeys(b.ctx, 1, boardKind, nil)[0]
	}
	return datastore.Put(b.ctx, b)
}
