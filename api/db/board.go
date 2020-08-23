package db

import (
	"context"

	"go.chromium.org/gae/service/datastore"
)

const (
	boardKind = "Board"
)

type Board struct {
	Key   *datastore.Key `gae:"$key"`
	Name  string
	Index int
	Memos []Memo
}

type Memo struct {
	Text  string
	Index int
}

func NewBoard(key *datastore.Key, name string) *Board {
	return &Board{
		Key:  key,
		Name: name,
	}
}

func NewMemo(text string) *Memo {
	return &Memo{Text: text}
}

func GetBoard(ctx context.Context, key *datastore.Key) (*Board, error) {
	b := NewBoard(key, "")
	if err := datastore.Get(ctx, b); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Board) Save(ctx context.Context) error {
	if b.Key == nil {
		b.Key = datastore.NewIncompleteKeys(ctx, 1, boardKind, nil)[0]
	}
	return datastore.Put(ctx, b)
}
