package main

import (
	"html/template"
	"path/filepath"
)

type ViewDataLabs struct {
	Cisco   []string
	Windows []string
	Linux   []string
}

// кеширование шаблонов.
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

func pageData(dirs map[string]string) (ViewDataLabs, error) {
	Labs := ViewDataLabs{
		Cisco:   []string{},
		Windows: []string{},
		Linux:   []string{},
	}
	for labSection, dir := range dirs {
		switch labSection {
		case "windows":
			labs, err := filepath.Glob(filepath.Join(dir, "*.pdf"))
			if err != nil {
				return Labs, err
			}
			for _, lab := range labs {
				Labs.Windows = append(Labs.Windows, filepath.Base(lab))
			}
		case "linux":
			labs, err := filepath.Glob(filepath.Join(dir, "*.pdf"))
			if err != nil {
				return Labs, err
			}
			for _, lab := range labs {
				Labs.Linux = append(Labs.Linux, filepath.Base(lab))
			}
		case "cisco":
			labs, err := filepath.Glob(filepath.Join(dir, "*.pdf"))
			if err != nil {
				return Labs, err
			}
			for _, lab := range labs {
				Labs.Cisco = append(Labs.Cisco, filepath.Base(lab))
			}
		}
	}
	return Labs, nil
}
