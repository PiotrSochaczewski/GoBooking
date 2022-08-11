package models

import "github.com/PiotrSochaczewski/GoBooking/internal/forms"

//TemplateData holds data sent from handler to templates
//CSRFToken = Cross Site Request Forgery Token
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
