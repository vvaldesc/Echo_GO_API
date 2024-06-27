package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Routes(e *echo.Echo) {

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("static/templates/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]string{
			"Title":   "My Go Web App",
			"Heading": "Hello, World!",
			"Message": "This is a simple Go web app.",
		})
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	e.Static("/static", "static")

}
