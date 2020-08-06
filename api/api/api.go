package api

import (
	"github.com/labstack/echo"
)

var (
	indexPath = "/"
	itemPath  = "/:id"
)

func Mount(g *echo.Group) {
	MountMemo(g.Group("/memos"))
	MountBoard(g.Group("/boards"))
}
