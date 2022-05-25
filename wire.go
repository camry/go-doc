//go:build wireinject
// +build wireinject

package main

import (
    "github.com/camry/dove"
    "github.com/camry/dove/log"
    "github.com/google/wire"
    "godoc/app"
    "godoc/routes"
)

func wireApp(l log.Logger) *dove.App {
    panic(wire.Build(routes.ProviderSet, app.ProviderSet))
}
