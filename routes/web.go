package routes

import (
    "fmt"
    "godoc/app/http/controllers"
    "godoc/routes/middleware"
    "html/template"
    "io"
    "net/http"

    "github.com/camry/g/glog"
    "github.com/google/wire"
    "github.com/labstack/echo/v4"
)

var ProviderSet = wire.NewSet(NewEcho)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

type HTTPErrorHandler struct{}

func (eh *HTTPErrorHandler) customHTTPErrorHandler(err error, c echo.Context) {
    code := http.StatusInternalServerError

    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
        glog.Errorw("Code", he.Code, "Message", he.Message)
    } else {
        glog.Error(err)
    }

    if err := c.File(fmt.Sprintf("resources/views/errors/%d.html", code)); err != nil {
        glog.Error(err)
    }
}

func NewEcho() *echo.Echo {
    e := echo.New()

    e.Static("/assets", "resources/assets")
    e.Static("/docs/files", "resources/docs/files")

    e.Use(middleware.EchoLogger())

    eh := &HTTPErrorHandler{}
    e.HTTPErrorHandler = eh.customHTTPErrorHandler

    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("resources/views/*.html")),
    }

    home := controllers.NewHome()

    e.GET("/", home.Index)
    e.GET("/docs", home.RootPage)
    e.GET("/docs/", home.RootPage)
    e.GET("/docs/:version", home.RootPage)
    e.GET("/docs/:version/", home.RootPage)
    e.GET("/docs/:version/:page", home.Show)

    return e
}
