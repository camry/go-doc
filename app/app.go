package app

import (
    "context"

    "github.com/camry/dove"
    "github.com/camry/dove/server/ghttp"
    "github.com/camry/g/glog"
    "github.com/google/wire"
    "github.com/labstack/echo/v4"
)

var ProviderSet = wire.NewSet(NewApp)

func NewApp(e *echo.Echo, l glog.Logger) *dove.App {
    hs := ghttp.NewServer(
        ghttp.Address(":3010"),
        ghttp.Handler(e),
    )
    app := dove.New(
        dove.Name("godoc"),
        dove.Version("v1.0.0"),
        dove.Server(hs),
        dove.BeforeStart(func(ctx context.Context) error {
            glog.SetLogger(l)
            return nil
        }),
    )
    return app
}
