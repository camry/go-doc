package app

import (
    "github.com/camry/dove"
    "github.com/camry/dove/log"
    "github.com/camry/dove/network/ghttp"
    "github.com/google/wire"
    "github.com/labstack/echo/v4"
)

var ProviderSet = wire.NewSet(NewApp)

func NewApp(e *echo.Echo, l log.Logger) *dove.App {
    log.SetLogger(l)
    hs := ghttp.NewServer(
        ghttp.Address(":3010"),
        ghttp.Handler(e),
    )
    app := dove.New(
        dove.Name("godoc"),
        dove.Version("v1.0.0"),
        dove.Server(hs),
        dove.Logger(log.GetLogger()),
    )
    return app
}
