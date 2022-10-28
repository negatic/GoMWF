package renderer

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/CloudyKit/jet"
)

type Renderer struct {
	Renderer   string
	Rootpath   string
	Secure     bool
	ServerName string
	Port       string
	JetViews   *jet.Set
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
		return c.JetPage(rw, r, view, variables, data)
	}
	return nil
}

// Rendering a template using the Go standard template engine
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

// Rendering a template using the Jet template engine
func (c *Renderer) JetPage(rw http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap
	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	t, err := c.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
