package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/team-pianoforte/mimemo2/api/db"
	"go.chromium.org/luci/gae/service/datastore"
)

func TestGetBoards(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	u := db.NewUser(AppEngineContext, "uid")

	b1 := db.NewBoard(AppEngineContext, 1, "b1")
	b2 := db.NewBoard(AppEngineContext, 2, "b2")
	b2.Memos = []db.Memo{{Text: "memo", Index: 0}}
	assert.NoError(t, b1.Save())
	assert.NoError(t, b2.Save())
	u.BoardKeys = []*datastore.Key{b1.Key, b2.Key}
	assert.NoError(t, u.Save())

	ctx.SetPath("/users/:uid/boards")
	ctx.SetParamNames("uid")
	ctx.SetParamValues("uid")

	expected := `[
    { "name": "b1", "memos": [], "index": 0, "id": 1 },
    { "name": "b2", "memos": [{ "text": "memo", "index": 0 }], "index": 0, "id": 2 }
  ]`

	assert.NoError(t, getBoards(ctx))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}

func TestGetBoards404(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctx.SetPath("/users/:uid/boards")
	ctx.SetParamNames("uid")
	ctx.SetParamValues("1")

	assert.NoError(t, getBoards(ctx))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
