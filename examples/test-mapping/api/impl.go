package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /stuff)
func (Server) CreateStuff(ctx echo.Context) error {
	// resp := {
	// 	Ping: "pong",
	// }

	return ctx.JSON(http.StatusOK, nil)
}
