package main

import (
	"html/template"
	"path/filepath"
)

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cahe := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*page.html"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}
		cahe[name] = ts
	}
	return cahe, nil
}
