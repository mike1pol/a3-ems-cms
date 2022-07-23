package helpers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/oxtoacart/bpool"

	. "github.com/mike1pol/rms/models"
	"encoding/json"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var templateConfig TemplateConfig

func LoadConfiguration() {
	templateConfig.TemplateLayoutPath = "views/layout/"
	templateConfig.TemplateModalsPath = "views/modals/"
	templateConfig.TemplateIncludePath = "views/"
}

type DutyUser struct {
	Duty []Duty
	User User
}

func LoadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.tpl")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.tpl")
	if err != nil {
		log.Fatal(err)
	}

	modalsFiles, err := filepath.Glob(templateConfig.TemplateModalsPath + "*.tpl")
	if err != nil {
		log.Fatal(err)
	}
	layoutFiles = append(layoutFiles, modalsFiles...)

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"buildDuty": func(d []Duty, u User) DutyUser {
			return DutyUser{
				Duty: d,
				User: u,
			}
		},
	}

	mainTemplate := template.New("main").Funcs(funcMap)

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func SendJson(w http.ResponseWriter, j interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(j)
}
