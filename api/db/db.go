package db

import (
	"context"
	"os"

	"go.chromium.org/gae/service/datastore"
)

const (
	userKind  = "User"
	boardKind = "Borad"
	memoKind  = "Memo"
)

var (
	ProjectId = os.Getenv("GCLOUD_PROJECT_ID")
)

type client struct {
	ctx context.Context
}

func (c *client) get(e interface{}) error {
	return datastore.Get(c.ctx, e)
}

func (c *client) put(k *datastore.Key, src interface{}) error {
	return datastore.Put(c.ctx, src)
}

func NewClient(ctx context.Context) *client {
	return &client{ctx: ctx}
}

type User struct {
	Key       *datastore.Key `gae:"$key"`
	Token     string
	BoardKeys []*datastore.Key
}

type Board struct {
	ID           int64 `gae:"$id"`
	Name         string
	PasswordHash string
	MemoKeys     []*datastore.Key
	Memos        []*Memo `gae:"-"`
}

type Memo struct {
	ID       int64          `gae:"$id"`
	BoardKey *datastore.Key `gae:"$parent"`
	Text     string
}

func NewBoard(name string, passwordHash string) *Board {
	return &Board{
		Name:         name,
		PasswordHash: passwordHash,
		MemoKeys:     make([]*datastore.Key, 0),
	}
}

func NewMemo(text string) *Memo {
	return &Memo{Text: text}
}

func (b *Board) MoveMemo(index, to int) {
	memo := b.MemoKeys[index]
	memos := b.MemoKeys
	memos = append(append(memos[0:to], memo), memos[to+1:]...)

	indexToRemove := index
	if index > to {
		indexToRemove = index + 1
	}
	memos = append(memos[0:indexToRemove], memos[indexToRemove+1:]...)
}

func GetBoard(ctx context.Context, id int64) (*Board, error) {
	b := &Board{ID: id}
	if err := datastore.Get(ctx, b); err != nil {
		return nil, err
	}
	println(b.ID)

	key := datastore.KeyForObj(ctx, b)

	q := datastore.NewQuery(memoKind).Ancestor(key)
	if err := datastore.GetAll(ctx, q, &b.Memos); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Board) Save(ctx context.Context) error {
	return datastore.Put(ctx, b)
}
