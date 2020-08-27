package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/team-pianoforte/mimemo2/api/db"
	"go.chromium.org/luci/gae/service/datastore"
)

func MountBoard(g *echo.Group) {
	g.GET(indexPath, getBoards)
	g.POST(indexPath, createBoard)

	g.GET(itemPath, getBoard)
	g.PATCH(itemPath, updateBoard)
	g.PUT(itemPath, updateBoard)
	g.DELETE(itemPath, deleteBoard)
}

func createBoard(c echo.Context) error {
	uid := c.Param("uid")
	u, err := db.GetUser(AppEngineContext, uid)
	if err == datastore.ErrNoSuchEntity {
		return c.NoContent(http.StatusNotFound)
	}

	b := db.NewBoard(AppEngineContext, 0, "")
	if err := c.Bind(b); err != nil {
		return err
	}
	if err := b.Save(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	println(b.ID)

	u.BoardKeys = append(u.BoardKeys, b.Key)
	if err := u.Save(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, b)
}

func updateBoard(c echo.Context) error {
	return nil
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

func getBoard(c echo.Context) error {
	return nil
}

func deleteBoard(c echo.Context) error {
	return nil
}
