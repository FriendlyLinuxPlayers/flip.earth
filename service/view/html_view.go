package view

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	ErrNoViewDirs      = errors.New("html_view: could not find field \"view_dirs\" in service configuration")
	ErrInvalidViewDirs = errors.New("html_view: \"view_dirs\" in service configuration is invalid: []string type assertion failed")
)

// HTMLView implements View for rendering in html.
type HTMLView struct {
	paths   []string
	viewBag map[string]interface{}
}

// Assign assigns avalue to the view which can reached under the given name
func (hv *HTMLView) Assign(name string, value interface{}) {
	hv.viewBag[name] = value
}

// Render renders the view to a http.ResponseWriter and sets the Content-Type header accordingly
func (hv *HTMLView) Render(tplName string, to http.ResponseWriter) error {
	//make the paths unique
	uniqPaths := make([]string, 0, len(hv.paths))
	found := make(map[string]bool)
	for _, val := range hv.paths {
		if _, ok := found[val]; !ok {
			found[val] = true
			uniqPaths = append(uniqPaths, val)
		}
	}

	tpl, err := template.ParseFiles(uniqPaths...)
	if err != nil {
		return err
	}

	to.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.ExecuteTemplate(to, tplName, hv.viewBag)
	if err != nil {
		return err
	}

	return nil
}

// AddTplDirs adds multiple directories to the view which will then be parsed
// and the templates files within will be added to the view hierarchy.
// Please keep in mind that this parsing is expensive
func (hv *HTMLView) AddTplDirs(paths ...string) error {
	p, err := parseDirs(paths)
	if err != nil {
		return err
	}

	hv.paths = append(hv.paths, p...)
	return nil
}

// AddTpls adds template files to the view hierarchy
func (hv *HTMLView) AddTpls(paths ...string) error {
	hv.paths = append(hv.paths, paths...)
	return nil
}

// parseDirs converts defaultDirs into paths acceptable for usage in an
// HTMLView.
func parseDirs(dirs []string) ([]string, error) {
	var paths []string
	for i := 0; i < len(dirs); i++ {
		dir := dirs[i]
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				dirs = append(dirs, path)
			} else {
				paths = append(paths, path)
			}
			return nil
		})

		if err != nil {
			return nil, err
		}
	}
	return paths, nil
}

// Init initializes and returns a constructor function for new *HTMLViews
func Init(deps, conf map[string]interface{}) (func() (*HTMLView, error), error) {
	viewDirs, ok := conf["view_dirs"]
	if !ok {
		return nil, ErrNoViewDirs
	}

	viewDirsStrings, ok := viewDirs.([]string)
	if !ok {
		return nil, ErrInvalidViewDirs
	}

	hvf := func() (*HTMLView, error) {
		paths, err := parseDirs(viewDirsStrings)
		if err != nil {
			return nil, err
		}

		return &HTMLView{
			paths: paths,
		}, nil
	}

	return hvf, nil
}
