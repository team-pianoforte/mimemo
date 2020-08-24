package db

import (
	"context"
	"os"

	"go.chromium.org/luci/gae/service/datastore"
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
