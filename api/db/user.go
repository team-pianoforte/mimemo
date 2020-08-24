package db

import (
	"context"

	"go.chromium.org/luci/gae/service/datastore"
)

const (
	userKind = "User"
)

type User struct {
	Key       *datastore.Key   `gae:"$key" json:""`
	BoardKeys []*datastore.Key `json:""`
	UID       string           `json:"uid"`
	ctx       context.Context  `gae:"-"`
}

func NewUser(ctx context.Context, uid string) *User {
	return &User{
		Key: datastore.NewKey(ctx, userKind, uid, 0, nil),
		UID: uid,
		ctx: ctx,
	}
}

func GetUser(ctx context.Context, uid string) (*User, error) {
	u := NewUser(ctx, uid)
	err := datastore.Get(ctx, u)
	return u, err
}

func (u *User) Save() error {
	return datastore.Put(u.ctx, u)
}

func (u *User) GetBoards() (boards []*Board, err error) {
	keys := make([]*datastore.Key, len(u.BoardKeys))
	for i, v := range u.BoardKeys {
		keys[i] = v
	}
	q := datastore.NewQuery(boardKind) //.Eq("Name", "b1")
	err = datastore.GetAll(u.ctx, q, &boards)

	boards = make([]*Board, len(u.BoardKeys))
	for i, k := range u.BoardKeys {
		boards[i] = NewBoard(u.ctx, k.IntID(), "")
	}
	err = datastore.Get(u.ctx, boards)

	return
}
