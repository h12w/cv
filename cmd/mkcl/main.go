package main

import (
	"flag"
	"log"
	"os"
	"path"
	"text/template"

	"h12.io/cv"
)

type options struct {
	Template string
	CV       string
	Job      string
	Output   string
}

func main() {
	var opt options
	flag.StringVar(&opt.Template, "template", path.Join(cv.DefaultDataPath, "letter.txt.tmpl"), "template filename")
	flag.StringVar(&opt.CV, "cv", path.Join(cv.DefaultDataPath, "resume.json"), "resume JSON filename")
	flag.StringVar(&opt.Job, "job", "job", "cover letter JSON filename")
	flag.StringVar(&opt.Output, "output", "letter.txt", "cover letter filename")
	flag.Parse()
	opt.Job = path.Join(cv.DefaultDataPath, "jobs", opt.Job+".json")

	var data cv.CoverLetter
	if err := cv.JSON(opt.Job, &data.Job); err != nil {
		log.Fatal(err)
	}
	if err := cv.JSON(opt.CV, &data.CV); err != nil {
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
	if err := tmpl.Execute(output, &data); err != nil {
		log.Fatal(err)
	}
}
