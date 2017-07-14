package config

import "fmt"

var (
	invalidPrefix    = "user provided %s path is invalid"
	notExistSuffix   = "does not exist"
	parseText        = "error in %s%sconfig file"
	typeAssertFailed = "string type assertion failed"
)

// Reason is used in ErrorDirInvalid and ErrorPathInvalid to indicate what
// caused the error.
type Reason int

const (
	// NotDir indicates that a directory was expected but not provided.
	NotDir Reason = iota
	// NotExist indicates that a given path does not exist.
	NotExist
	// TypeAssert indicates that a string type assertion on a path failed.
	TypeAssert
)

// ErrorDirInvalid implements error, indicating that there is a problem related
// to the value of working_dir, as provided in the config Definition.
type ErrorDirInvalid struct {
	// Path is the path that caused the error, if available.
	Path string
	// Cause is the Reason the error happened.
	Cause Reason
	// Err is the underlying error, if available.
	Err error
}

func (e *ErrorDirInvalid) Error() string {
	what := "\"working_dir\""
	why := ""
	switch e.Cause {
	case NotDir:
		why = ": is not a directory"
	case NotExist:
		why = fmt.Sprintf(": %q %s", e.Path, notExistSuffix)
	case TypeAssert:
		why = fmt.Sprintf(": %s", typeAssertFailed)
	}
	return fmt.Sprintf(invalidPrefix+"%s", what, why)
}

// ErrorPathInvalid implements error, indicating that there is a problem with
// the path to the custom user configuration file. This can indicate a problem
// related to the values of working_dir or config_file in the config Definition.
type ErrorPathInvalid struct {
	// Path is the path that caused the error, if available.
	Path string
	// Cause is the Reason the error happened. Note that NotDir should never
	// be the value, as it does not apply to the config file path.
	Cause Reason
	// Err is the underlying error, if available.
	Err error
}

func (e *ErrorPathInvalid) Error() string {
	what := "\"config_path\""
	why := ""
	switch e.Cause {
	case NotExist:
		why = fmt.Sprintf(": %q %s", e.Path, notExistSuffix)
	case TypeAssert:
		why = fmt.Sprintf(": %s", typeAssertFailed)
	}
	return fmt.Sprintf(invalidPrefix+"%s", what, why)
}

// ErrorParse
type ErrorParse struct {
	// Path is the path to the file that caused the error.
	Path string
	// IsDefault indicates if the problem is with the default config file or
	// a custom config file.
	IsDefault bool
	// IsRead indicates if the error occured while reading the file or while
	// parsing the file.
	Reading bool
	// Err is the underlying error.
	Err error
}

func (e *ErrorParse) Error() string {
	what := ""
	doing := ""
	switch e.IsDefault {
	case true:
		what = "default "
	case false:
		what = "false "
	}
	switch e.Reading {
	case true:
		doing = "reading "
	case false:
		doing = "parsing "
	}
	return fmt.Sprintf(parseText+": %s", doing, what, e.Err.Error())
}
