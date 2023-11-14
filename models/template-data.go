package model

import "github.com/mauFade/bookings/pkg/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Error     string
	Flash     string
	Warning   string
	Form      *forms.Form
}
