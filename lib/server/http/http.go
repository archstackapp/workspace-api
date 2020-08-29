package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/archstack/workspace-api/lib/logger"
	"go.uber.org/zap"
)

// EchoServer wraps an echo instance.
type EchoServer struct {
	Echo   *echo.Echo
	config *Config
}

// NewEchoService creates a new EchoServer service.
func NewEchoService(logger *logger.Logger, config *Config) (*EchoServer, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.Zap.Errorw("error during handler",
			zap.String("path", c.Path()),
			"params", c.QueryParams(),
			zap.String("err", err.Error()))

		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Use(middleware.Recover())

	server := EchoServer{e, config}

	return &server, nil
}

// Start starts the echo server.
func (server *EchoServer) Start() {
	server.Echo.Start(fmt.Sprintf(":%s", server.config.Port))
}
