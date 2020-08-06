package api

import (
	"github.com/labstack/echo"
)

var (
	indexPath = "/"
	itemPath  = "/:id"
)

func MountBoard(g *echo.Group) {

	g.GET(indexPath, getBoards)
	g.POST(indexPath, createBoard)

	g.GET(itemPath, getBoard)
	g.PATCH(itemPath, updateBoard)
	g.PUT(itemPath, updateboard)
	g.DELETE(itemPath, deleteBoard)
}

func createBoard(context *echo.Context) error {
	return nil
}

func updateBoard(context *echo.Context) error {
	return nil
}

func getBoards(context *echo.Context) error {
	return nil
}

func getBoard(context *echo.Context) error {
	return nil
}

func deleteBoard(context *echo.Context) error {
	return nil
}
