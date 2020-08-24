package api

import (
	"context"

	"github.com/labstack/echo"
)

var (
	AppEngineContext context.Context
	indexPath        = "/"
	itemPath         = "/:id"
)

func Mount(g *echo.Group) {
	MountBoard(g.Group("/"))
}
