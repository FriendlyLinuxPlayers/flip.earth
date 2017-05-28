package view

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type HtmlView struct {
	paths   []string
	viewBag map[string]interface{}
}

func (hv *HtmlView) Assign(name string, value interface{}) {
	hv.viewBag[name] = value
}

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

func (hv *HtmlView) AddTplDirs(paths ...string) error {
	p, err := parseDirs(paths)
	if err != nil {
		return err
	}

	hv.paths = append(hv.paths, p...)
	return nil
}

func (hv *HtmlView) AddTpls(paths ...string) error {
	hv.paths = append(hv.paths, paths...)
	return nil
}

type HtmlViewFactory struct {
	defualtDirs []string
}

func (hvs *HtmlViewFactory) New() (*HtmlView, error) {
	paths, err := parseDirs(hvs.defualtDirs)
	if err != nil {
		return nil, err
	}

	return &HtmlView{
		paths: paths,
	}, nil

}

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

type Initializer struct {
}

func (i Initializer) Init(deps map[string]interface{}, conf map[string]interface{}) (interface{}, error) {
	viewDirs, ok := conf["view_dirs"]
	if !ok {
		return nil, fmt.Errorf("html_view: could not find field 'view_dirs' in service configuration")
	}

	hvf := &HtmlViewFactory{
		defualtDirs: viewDirs.([]string),
	}

	return hvf, nil
}
