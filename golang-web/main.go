package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title      string
	FirstName  string
	Usrname    string
	DeviceName string
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abot)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/apply", aply)
	http.HandleFunc("/login", loginUser)
	http.HandleFunc("/widget-form", widgetForm)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("template/style/css"))))

	http.ListenAndServe(":8000", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func abot(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w, "about.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func cntct(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w, "contact.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func aply(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Apply Page",
	}

	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func loginUser(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "User Login",
	}

	var username string
	if req.Method == http.MethodPost {
		username = req.FormValue("uname")
		pd.Usrname = username
	}

	err := tpl.ExecuteTemplate(w, "login.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func widgetForm(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Widget Form",
	}

	var widgetname string
	if req.Method == http.MethodPost {
		widgetname = req.FormValue("device")
		pd.DeviceName = widgetname
	}

	err := tpl.ExecuteTemplate(w, "widget-form.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
