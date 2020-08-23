package db

import (
	"context"

	"go.chromium.org/gae/service/datastore"
)

const (
	userKind = "User"
)

type User struct {
	Key       *datastore.Key `gae:"$key"`
	BoardKeys []*datastore.Key
	UID       string
}

func NewUser(uid string) *User {
	return &User{UID: uid}
}

func GetUser(ctx context.Context, key *datastore.Key) (*User, error) {
	u := &User{Key: key}
	if err := datastore.Get(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Save(ctx context.Context) error {
	return datastore.Put(ctx, u)
}

func (u *User) GetBoards(ctx context.Context) (boards []*Board, err error) {
	keys := make([]*datastore.Key, len(u.BoardKeys))
	for i, v := range u.BoardKeys {
		keys[i] = v
	}
	q := datastore.NewQuery(boardKind) //.Eq("Name", "b1")
	err = datastore.GetAll(ctx, q, &boards)

	boards = make([]*Board, len(u.BoardKeys))
	for i, k := range u.BoardKeys {
		boards[i] = NewBoard(k, "")
	}
	err = datastore.Get(ctx, boards)

	return
}
