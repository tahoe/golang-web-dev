package main

import (
	"log"
	"os"
	"text/template"
)

// creates a pointer to the template object, basically creating an empty object that
// has all the methods but no data yet.
var tpl *template.Template

// func init() seems to be a standard function that gets auto run just like main
// but probably before main
func init() {
	// 
	// populate the pointer object using ParseGlob
	// function call (Must) taking as an argument the results of a function call (ParseGlob)
	// template.Must takes in a template pointer and an error as it's input
	// which is exactly what most of the parsing methods return
	// this would be a good place (the Must method) to do auto error codes and whatnot
	// 
	// now any files in the subfolder templates will be in the tpl object
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// here we take the template object and call Execute which will execute
	// the first template it runs into
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// ExecuteTemplate: explicit template execution
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// ExecuteTemplate: explicit template execution
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// ExecuteTemplate: explicit template execution
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
