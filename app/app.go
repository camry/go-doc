package app

import (
    "context"
    "godoc/app/logger"

    "github.com/camry/dove"
    "github.com/camry/dove/server/ghttp"
    "github.com/camry/g/glog"
    "github.com/google/wire"
    "github.com/labstack/echo/v4"
)

var ProviderSet = wire.NewSet(NewApp)

func NewApp(e *echo.Echo) *dove.App {
    hs := ghttp.NewServer(
        ghttp.Address(":3010"),
        ghttp.Handler(e),
    )
    app := dove.New(
        dove.Name("godoc"),
        dove.Version("v1.0.0"),
        dove.Server(hs),
        dove.BeforeStart(func(ctx context.Context) error {
            glog.SetLogger(logger.NewAppLogger())
            return nil
        }),
    )
    return app
}
