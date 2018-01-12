package main

import (
	"log"
	"os"
	"text/template"
)

// creates a pointer (container for) to a Template "Type"
var tpl *template.Template

func init() {
	// init function gets run first
	// populate the tpl pointer with a Template object
	// pointing to a specific template (tpl.gohtml in this case)
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// ExecuteTemplate takes a writer (Stdout here) a template name
	// and possibly values or nil and executes it
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
