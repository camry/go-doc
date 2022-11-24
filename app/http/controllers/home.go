package controllers

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"

    "github.com/camry/g/glog"
    "github.com/labstack/echo/v4"
    "github.com/russross/blackfriday/v2"
    "gopkg.in/yaml.v3"
)

const DefaultVersion = "1.0"

type home struct {
    l *glog.Helper
    v map[string]string
}

type Config struct {
    Versions map[string]string `yaml:"versions"`
}

func NewHome(l glog.Logger) *home {
    return &home{l: glog.NewHelper(l)}
}

func (h *home) loadVersions() {
    var config Config
    buf, err1 := ioutil.ReadFile("config/versions.yaml")
    if err1 != nil {
        h.l.Errorf("load config/versions.yaml error: %s", err1)
    }
    err2 := yaml.Unmarshal(buf, &config)
    if err2 != nil {
        h.l.Errorf("yaml Unmarshal config/versions.yaml error: %s", err2)
    }
    h.v = config.Versions
}

func (h *home) Index(c echo.Context) error {
    h.loadVersions()
    return c.Render(http.StatusOK, "home.html", map[string]any{
        "v": h.v,
    })
}

func (h *home) RootPage(c echo.Context) error {
    version := c.Param("version")

    h.loadVersions()

    if _, ok := h.v[version]; !ok {
        version = DefaultVersion
    }

    return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/docs/%s/1", version))
}

func (h *home) Show(c echo.Context) error {
    version := c.Param("version")
    page := c.Param("page")

    if page == "" {
        page = "1"
    }

    h.loadVersions()

    // 读取文档菜单
    path1 := fmt.Sprintf("resources/docs/%s/documentation.md", version)
    input1, err1 := ioutil.ReadFile(path1)
    if err1 != nil {
        return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/docs/%s/%s", DefaultVersion, version))
    }
    output1 := blackfriday.Run(blackfriday.Run(input1))

    // 读取文档内容
    path2 := fmt.Sprintf("resources/docs/%s/%s.md", version, page)
    input2, err2 := ioutil.ReadFile(path2)
    if err2 != nil {
        return c.Render(http.StatusOK, "docs.html", map[string]any{
            "title":          "Page not found",
            "index":          template.HTML(strings.ReplaceAll(string(output1), "{{version}}", version)),
            "content":        template.HTML(fmt.Sprintf("Markdown 文档不存在！")),
            "currentVersion": version,
            "versions":       h.v,
        })
    }
    output2 := blackfriday.Run(input2)

    reg, err3 := regexp.Compile("<h1+>([\\s\\S]*?)</h1>")
    if err3 != nil {
        return err3
    }

    var titles []string
    if _, ok := h.v[version]; ok {
        titles = append(titles, h.v[version])
    }
    titles = append(titles, strings.ReplaceAll(strings.ReplaceAll(string(reg.Find(output2)), "<h1>", ""), "</h1>", ""))

    return c.Render(http.StatusOK, "docs.html", map[string]any{
        "title":          strings.Join(titles, " - "),
        "index":          template.HTML(strings.ReplaceAll(string(output1), "{{version}}", version)),
        "content":        template.HTML(strings.ReplaceAll(string(output2), "{{version}}", version)),
        "currentVersion": version,
        "versions":       h.v,
    })
}
