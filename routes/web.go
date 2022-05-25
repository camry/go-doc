package routes

import (
    "github.com/camry/dove/log"
    "github.com/google/wire"
    "github.com/labstack/echo/v4"
    "godoc/app/http/controllers"
    "html/template"
    "io"
)

var ProviderSet = wire.NewSet(NewEcho)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func NewEcho(l log.Logger) *echo.Echo {
    e := echo.New()

    e.Static("/assets", "resources/assets")
    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("resources/views/*.html")),
    }

    home := controllers.NewHome(l)

    e.GET("/", home.Index)
    e.GET("/docs/", home.RootPage)
    e.GET("/docs/:version/:page", home.Show)

    return e
}
