package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"os"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

//func main() {
//	e := echo.New()
//	renderer := &TemplateRenderer{
//		templates: template.Must(template.ParseGlob("*.html")),
//	}
//	e.Renderer = renderer
//
//	// Named route "foobar"
//	e.GET("/something", func(c echo.Context) error {
//		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
//			"name": "Dolly!",
//		})
//	}).Name = "foobar"
//
//	e.Logger.Fatal(e.Start(":8000"))
//}


func main(){
	name := Name{
		Fname: "ragil",
		Lname: "Maulana",
	}
	template, _ := template.ParseFiles("index.html")
	template.Execute(os.Stdout, name)

}

type Name struct {
 Fname string
 Lname string
}