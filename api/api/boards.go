package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/team-pianoforte/mimemo2/api/db"
	"go.chromium.org/luci/gae/service/datastore"
)

func MountBoard(g *echo.Group) {
	g.GET(indexPath, getBoards)
}

func getBoards(c echo.Context) error {
	uid := c.Param("uid")
	u, err := db.GetUser(AppEngineContext, uid)
	if err == datastore.ErrNoSuchEntity {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return err
	}
	boards, err := u.GetBoards()
	return c.JSON(http.StatusOK, boards)
}
