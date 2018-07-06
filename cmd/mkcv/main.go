package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"text/template"
)

type options struct {
	Template string
	JSON     string
	Output   string
}

func main() {
	var opt options
	flag.StringVar(&opt.Template, "template", "resume.html.tmpl", "template filename")
	flag.StringVar(&opt.JSON, "json", "resume.json", "resume JSON filename")
	flag.StringVar(&opt.Output, "output", "resume.html", "resume HTML filename")
	flag.Parse()

	jsonFile, err := os.Open(opt.JSON)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var cv CV
	if err := json.NewDecoder(jsonFile).Decode(&cv); err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles(opt.Template)
	if err != nil {
		log.Fatal(err)
	}
	output, err := os.Create(opt.Output)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	if err := tmpl.Execute(output, &cv); err != nil {
		log.Fatal(err)
	}
}
