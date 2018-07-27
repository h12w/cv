package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"

	"gopkg.in/russross/blackfriday.v2"
	"h12.io/cv"
)

type options struct {
	Template string
	CV       string
	Output   string
}

func main() {
	var opt options
	flag.StringVar(&opt.Template, "template", "resume.html.tmpl", "template filename")
	flag.StringVar(&opt.CV, "cv", path.Join(cv.DefaultDataPath, "resume.json"), "resume JSON filename")
	flag.StringVar(&opt.Output, "output", "resume.html", "resume HTML filename")
	flag.Parse()

	var cvData cv.CV
	if err := cv.JSON(opt.CV, &cvData); err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New(opt.Template).Funcs(template.FuncMap{"markdown": markdown}).ParseFiles(opt.Template)
	if err != nil {
		log.Fatal(err)
	}
	output, err := os.Create(opt.Output)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	if err := tmpl.Execute(output, &cvData); err != nil {
		log.Fatal(err)
	}
}

func markdown(args ...interface{}) template.HTML {
	s := blackfriday.Run([]byte(fmt.Sprint(args...)))
	s = bytes.TrimSpace(s)
	s = bytes.TrimPrefix(s, []byte("<p>"))
	s = bytes.TrimSuffix(s, []byte("</p>"))
	return template.HTML(s)
}
