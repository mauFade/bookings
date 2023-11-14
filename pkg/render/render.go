package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	model "github.com/mauFade/bookings/models"
	"github.com/mauFade/bookings/pkg/config"
)

var templateCache = make(map[string]*template.Template)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(templ *model.TemplateData) *model.TemplateData {
	return templ
}

func RenderTemplate(response http.ResponseWriter, request *http.Request, templ string, templData *model.TemplateData) {
	var templatesCache map[string]*template.Template
	var err error

	if app.UseCache {
		templatesCache = app.TemplateCache
	} else {
		templatesCache, err = CreateTemplateCache()

		if err != nil {
			log.Println(err)
		}
	}

	temp, ok := templatesCache[templ]

	if !ok {
		log.Fatal("Error loading cache")
	}

	buf := new(bytes.Buffer)

	templData = AddDefaultData(templData, request)

	_ = temp.Execute(buf, templData)

	_, err = buf.WriteTo(response)

	if err != nil {
		log.Println(err)
	}
}

func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templs, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			templs, err = templs.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return cache, err
			}
		}

		cache[name] = templs
	}

	return cache, nil
}
