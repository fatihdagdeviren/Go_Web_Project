package TemplateParser

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("Views/Shared/header.html",
	"Views/Shared/footer.html",
	"Views/Home/index.html",
	"Views/UserProfile/index.html",
	"Views/User/useredit.html"))

//Display the named template
func Display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}