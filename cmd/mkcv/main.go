package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strings"

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

	// Validate references
	if err := validateReferences(&cvData); err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New(path.Base(opt.Template)).Funcs(template.FuncMap{
		"markdown": markdown,
		"slugify":  slugify,
	}).ParseFiles(opt.Template)
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

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "&", "and")
	s = strings.ReplaceAll(s, "/", "-or-")
	return s
}

func validateReferences(cvData *cv.CV) error {
	// Build a map of all portfolio sub-item titles (slugified)
	portfolioTitles := make(map[string]bool)
	for _, portfolio := range cvData.Portfolio {
		for _, item := range portfolio.Items {
			slug := slugify(item.Title)
			portfolioTitles[slug] = true
		}
	}

	// Check all references in Work entries
	var missingRefs []string
	for _, work := range cvData.Work {
		for _, ref := range work.RefItems {
			slug := slugify(ref)
			if !portfolioTitles[slug] {
				missingRefs = append(missingRefs, fmt.Sprintf("reference '%s' in work entry '%s' at %s", ref, work.Position, work.Company))
			}
		}
	}

	if len(missingRefs) > 0 {
		return fmt.Errorf("references not found:\n  %s", strings.Join(missingRefs, "\n  "))
	}

	return nil
}
