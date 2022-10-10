package main

import (
    "github.com/camry/dove"
    "github.com/camry/g/glog"
    "github.com/google/wire"
    "godoc/app"
    "godoc/routes"
)

func wireApp(l glog.Logger) *dove.App {
    panic(wire.Build(routes.ProviderSet, app.ProviderSet))
}
