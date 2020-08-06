package api

import (
	"github.com/labstack/echo"
)

func MountMemo(g *echo.Group) {
	g.POST(indexPath, createBoard)

	g.PATCH(itemPath, updateBoard)
	g.PUT(itemPath, updateBoard)
	g.DELETE(itemPath, updateBoard)
}

func createMemo(context *echo.Context) error {
	return nil
}

func updateMemo(context *echo.Context) error {
	return nil
}

func deleteMemo(context *echo.Context) error {
	return nil
}
