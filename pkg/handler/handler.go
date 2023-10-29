package handler

import (
	"net/http"

	model "github.com/mauFade/bookings/models"
	"github.com/mauFade/bookings/pkg/config"
	"github.com/mauFade/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(response http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(response, "home.page.html", &model.TemplateData{})
}

func (repo *Repository) About(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)

	stringMap["hello"] = "Hello World"
	stringMap["remote_ip"] = request.RemoteAddr

	render.RenderTemplate(response, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}
