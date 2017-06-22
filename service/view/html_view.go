package view

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// HtmlView implements View for rendering in html.
type HtmlView struct {
	paths   []string
	viewBag map[string]interface{}
}

// Assign
func (hv *HtmlView) Assign(name string, value interface{}) {
	hv.viewBag[name] = value
}

// Render
func (hv *HtmlView) Render(tplName string, to io.Writer) error {
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
	err = tpl.ExecuteTemplate(to, tplName, hv.viewBag)
	if err != nil {
		return err
	}

	return nil
}

// AddTplDirs
func (hv *HtmlView) AddTplDirs(paths ...string) error {
	p, err := parseDirs(paths)
	if err != nil {
		return err
	}

	hv.paths = append(hv.paths, p...)
	return nil
}

// AddTpls
func (hv *HtmlView) AddTpls(paths ...string) error {
	hv.paths = append(hv.paths, paths...)
	return nil
}

// HtmlViewFactory implements ViewFactory for creating HtmlViews.
type HtmlViewFactory struct {
	defualtDirs []string
}

// New returns an initialized HtmlView.
func (hvs *HtmlViewFactory) New() (*HtmlView, error) {
	paths, err := parseDirs(hvs.defualtDirs)
	if err != nil {
		return nil, err
	}

	return &HtmlView{
		paths: paths,
	}, nil

}

// parseDirs converts defaultDirs into paths acceptable for usage in an HtmlView.
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

// Initializer implements the Intializer interface for HtmlView.
type Initializer struct {
}

// Init
func (i Initializer) Init(deps map[string]interface{}, conf map[string]interface{}) (interface{}, error) {
	viewDirs, ok := conf["view_dirs"]
	if !ok {
		return nil, fmt.Errorf("html_view: could not find field 'view_dirs' in service configuration")
	}

	viewDirsStrings, ok := viewDirs.([]string)
	if !ok {
		return nil, fmt.Errorf("html_view: 'view_dirs' in service configuration is invalid: []string type assertion failed")
	}
	hvf := &HtmlViewFactory{
		defualtDirs: viewDirsStrings,
	}

	return hvf, nil
}
