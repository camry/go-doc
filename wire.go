//go:build wireinject
// +build wireinject

package main

import (
    "godoc/app"
    "godoc/routes"

    "github.com/camry/dove"
    "github.com/google/wire"
)

func wireApp() *dove.App {
    panic(wire.Build(routes.ProviderSet, app.ProviderSet))
}
