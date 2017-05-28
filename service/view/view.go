package view

import "io"

//View can have values added to it and be rendered to a Writer
type View interface {
	Assign(name string, value interface{})
	Render(tplName string, to io.Writer) error
	AddTplDirs(paths ...string) error
	AddTpls(paths ...string) error
}

type ViewFactory interface {
	New() (View, error)
}
