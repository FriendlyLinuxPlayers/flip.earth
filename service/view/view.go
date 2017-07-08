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

// New is a function type that returns a new, initalized view
type New func() (*View, error)
