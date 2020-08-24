package api

import (
	"context"
	"os"
	"testing"

	"github.com/labstack/echo"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
)

var (
	e = echo.New()
)

func TestMain(m *testing.M) {
	config = &ApiConfig{SkipAuth: true}
	AppEngineContext = memory.Use(context.Background())
	testable := datastore.GetTestable(AppEngineContext)
	testable.AutoIndex(true)
	os.Exit(m.Run())
}
