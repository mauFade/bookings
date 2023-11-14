package handler

import (
	"fmt"
	"log"
	"net/http"

	model "github.com/mauFade/bookings/models"
	"github.com/mauFade/bookings/pkg/config"
	"github.com/mauFade/bookings/pkg/forms"
	"github.com/mauFade/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
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

func (m *Repository) Reservation(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, request, "make-reservation.page.html", &model.TemplateData{})
}

func (m *Repository) Generals(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, request, "generals.page.html", &model.TemplateData{})
}

func (m *Repository) Majors(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, request, "majors.page.html", &model.TemplateData{})
}

func (m *Repository) Availability(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, request, "search-availability.page.html", &model.TemplateData{})
}

func (m *Repository) PostAvailability(response http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	response.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

func (m *Repository) PostReservation(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	reservation := model.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
	}

	form := forms.New(request.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, request)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(response, request, "make-reservation.page.tmpl", &model.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(request.Context(), "reservation", reservation)
	http.Redirect(response, request, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(response http.ResponseWriter, request *http.Request) {
	reservation, ok := m.App.Session.Get(request.Context(), "reservation").(model.Reservation)
	if !ok {
		log.Println("can't get item from session")
		m.App.Session.Put(request.Context(), "error", "Can't get reservation from session")
		http.Redirect(response, request, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(request.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(response, request, "reservation-summary.page.tmpl", &model.TemplateData{
		Data: data,
	})
}

func (m *Repository) Contact(request http.ResponseWriter, response *http.Request) {
	render.RenderTemplate(request, response, "contact.page.html", &model.TemplateData{})
}
