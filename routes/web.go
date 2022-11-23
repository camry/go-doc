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

type HTTPErrorHandler struct {
    l *glog.Helper
}

func (eh *HTTPErrorHandler) customHTTPErrorHandler(err error, c echo.Context) {
    code := http.StatusInternalServerError

    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
        eh.l.Errorw("Code", he.Code, "Message", he.Message)
    } else {
        eh.l.Error(err)
    }

    if err := c.File(fmt.Sprintf("resources/views/errors/%d.html", code)); err != nil {
        eh.l.Error(err)
    }
}

func NewEcho(l glog.Logger) *echo.Echo {
    e := echo.New()

    e.Static("/assets", "resources/assets")
    e.Static("/docs/files", "resources/docs/files")

    e.Use(middleware.EchoLogger(l))

    eh := &HTTPErrorHandler{l: glog.NewHelper(l)}
    e.HTTPErrorHandler = eh.customHTTPErrorHandler

    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("resources/views/*.html")),
    }

    home := controllers.NewHome(l)

    e.GET("/", home.Index)
    e.GET("/docs", home.RootPage)
    e.GET("/docs/", home.RootPage)
    e.GET("/docs/:version", home.RootPage)
    e.GET("/docs/:version/", home.RootPage)
    e.GET("/docs/:version/:page", home.Show)

    return e
}
