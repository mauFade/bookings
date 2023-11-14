package handler

import (
	"fmt"
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

	render.RenderTemplate(response, request, "home.page.html", &model.TemplateData{})
}

func (repo *Repository) About(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)

	stringMap["hello"] = "Hello World"
	stringMap["remote_ip"] = request.RemoteAddr

	render.RenderTemplate(response, request, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(request http.ResponseWriter, response *http.Request) {
	render.RenderTemplate(request, response, "make-reservation.page.html", &model.TemplateData{})
}

func (m *Repository) Generals(request http.ResponseWriter, response *http.Request) {
	render.RenderTemplate(request, response, "generals.page.html", &model.TemplateData{})
}

func (m *Repository) Majors(request http.ResponseWriter, response *http.Request) {
	render.RenderTemplate(request, response, "majors.page.html", &model.TemplateData{})
}

func (m *Repository) Availability(request http.ResponseWriter, response *http.Request) {
	render.RenderTemplate(request, response, "search-availability.page.html", &model.TemplateData{})
}

func (m *Repository) PostAvailability(request http.ResponseWriter, response *http.Request) {
	start := response.Form.Get("start")
	end := response.Form.Get("end")

	request.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

func (m *Repository) Contact(request http.ResponseWriter, response *http.Request) {
	render.RenderTemplate(request, response, "contact.page.html", &model.TemplateData{})
}
