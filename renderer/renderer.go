package renderer

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Renderer struct {
	Renderer   string
	Rootpath   string
	Secure     bool
	ServerName string
	Port       string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]float32
	CSRFToken       string
	Secure          bool
	Port            string
	SeverName       string
}

func (c *Renderer) Page(rw http.ResponseWriter, r http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(c.Renderer) {
	case "standard":
		return c.StandardPage(rw, r, view, data)
	case "jet":
	}
	return nil
}

func (c *Renderer) StandardPage(rw http.ResponseWriter, r http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", c.Rootpath, view))
	if err != nil {
		return err
	}
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	err = tmpl.Execute(rw, &td)

	if err != nil {
		return err
	}

	return nil
}
