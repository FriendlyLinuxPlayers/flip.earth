package view

import "io"

// View is the interface for adding values that can be rendered to a Writer.
type View interface {
	Assign(name string, value interface{})
	Render(tplName string, to io.Writer) error
	AddTplDirs(paths ...string) error
	AddTpls(paths ...string) error
}

// ViewFactory is the interface for creating an initialized View.
type ViewFactory interface {
	New() (View, error)
}
