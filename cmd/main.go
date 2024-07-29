package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)


type Templates struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
  return &Templates{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  count := Count { Count: 0 }

  e.Renderer = newTemplate()
  e.GET("/", func(c echo.Context) error {
    count.Count++
    return c.Render(200, "index", count)
  })

  e.Logger.Fatal(e.Start(":42069")) http.ResponseWriter, r *http.Request
}
