package db

import (
	"context"

	"go.chromium.org/luci/gae/service/datastore"
)

const (
	boardKind = "Board"
)

type Board struct {
	Key   *datastore.Key  `gae:"$key" json:"-"`
	ID    int64           `gae:"$id" json:"id"`
	Name  string          `json:"name"`
	Index int             `json:"index"`
	Memos []Memo          `json:"memos"`
	ctx   context.Context `gae:"-" json:"-"`
}

type Memo struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

func NewBoard(ctx context.Context, id int64, name string) (b *Board) {
	b = &Board{
		Name:  name,
		ctx:   ctx,
		ID:    id,
		Memos: make([]Memo, 0),
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
	//b.ID = id

	return b, nil
}

func (b *Board) Save() error {
	if b.Key == nil {
		b.Key = datastore.NewIncompleteKeys(b.ctx, 1, boardKind, nil)[0]
	}
	return datastore.Put(b.ctx, b)
}
