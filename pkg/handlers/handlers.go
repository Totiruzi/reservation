package handlers

import (
	"net/http"

	"github.com/Totiruzi/reservation/pkg/config"
	"github.com/Totiruzi/reservation/pkg/models"
	"github.com/Totiruzi/reservation/pkg/render"
)

// Repo the repositry used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// Contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

// Esimi page handler
func (m *Repository) Esimi(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "esimi.page.html", &models.TemplateData{})
}

// Lavinda page handler
func (m *Repository) Lavinda(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "lavinda.page.html", &models.TemplateData{})
}

// Reservattion takes you to the reservation page
func (m *Repository) Reservastion(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "reservation.page.html", &models.TemplateData{})
}

// Omocha page handler
func (m *Repository) Omoche(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "omocha.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello from About again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
// 	remoteIP := r.RemoteAddr
// 	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

// 	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
// }
