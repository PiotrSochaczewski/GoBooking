package handlers

import (
	"fmt"
	"net/http"

	"github.com/PiotrSochaczewski/GoBooking/pkg/config"
	"github.com/PiotrSochaczewski/GoBooking/pkg/models"
	"github.com/PiotrSochaczewski/GoBooking/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// MakeReservation is the make reservation page handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// TraditionalHanok is a traditional hanok page handler
func (m *Repository) TraditionalHanok(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "traditional-hanok.page.tmpl", &models.TemplateData{})
}

//ModernHanok is a modern hanok page handler
func (m *Repository) ModernHanok(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "modern-hanok.page.tmpl", &models.TemplateData{})
}

// Modern is the modern house page handler
func (m *Repository) Modern(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "modern.page.tmpl", &models.TemplateData{})
}

//SearchAvailability is a search availability page handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

//PostAvailability is a search availability page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
