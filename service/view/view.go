package view

import "io"

// View is the interface for adding values that can be rendered to a Writer.
type View interface {
	// Assign
	Assign(name string, value interface{})
	// Render
	Render(tplName string, to io.Writer) error
	// AddTplDirs
	AddTplDirs(paths ...string) error
	// AddTpls
	AddTpls(paths ...string) error
}

// ViewFactory is the interface for creating an initialized View.
type ViewFactory interface {
	// New returns an initialized View
	New() (View, error)
}
