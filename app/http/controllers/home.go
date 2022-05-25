package controllers

import (
    "bytes"
    "fmt"
    "github.com/camry/dove/log"
    "github.com/labstack/echo/v4"
    "github.com/yuin/goldmark"
    "html/template"
    "io/ioutil"
    "net/http"
)

const DefaultVersion = "appz"

type home struct {
    l *log.Helper
}

func NewHome(l log.Logger) *home {
    return &home{l: log.NewHelper(l)}
}

func (h *home) Index(c echo.Context) error {
    return c.Render(http.StatusOK, "home.html", "")
}

func (h *home) RootPage(c echo.Context) error {
    return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/docs/%s/1", DefaultVersion))
}

func (h *home) Show(c echo.Context) error {
    version := c.Param("version")
    page := c.Param("page")

    path := fmt.Sprintf("resources/docs/%s/%s.md", version, page)

    data, err1 := ioutil.ReadFile(path)

    if err1 != nil {
        return err1
    }

    var buf bytes.Buffer

    if err2 := goldmark.Convert(data, &buf); err2 != nil {
        return err2
    }

    return c.Render(http.StatusOK, "docs.html", map[string]any{
        "content": template.HTML(buf.String()),
    })
}
