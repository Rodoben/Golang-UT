package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk.

	cwd, err := os.Getwd()
	if err != nil {
		// Handle the error
	}
	fmt.Println("Current Working Directory:", cwd)

	parsedTemplate, err := template.ParseFiles("cmd/templates/" + t)
	if err != nil {
		http.Error(w, "bad requccccest", http.StatusBadRequest)
		return err
	}

	// execute the template, passing it data, if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
